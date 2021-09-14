# syntax=docker/dockerfile:1
ARG GOLANG_VERSION=1.17-alpine

# Builder image
FROM golang:${GOLANG_VERSION} as builder

WORKDIR /go/src/github.com/myugen/go-mss-twirp/

COPY go.mod ./
COPY go.sum ./
RUN echo 'Downloading go.mod dependencies' \
    && go mod download

COPY *.go ./
RUN echo 'Installing tools from tools.go' \
    && cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %
RUN echo 'Generating gRPC services' \
    && go generate ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


# Generate clean, final image for end users
FROM alpine:latest as app

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/myugen/go-mss-twirp/app ./

ENTRYPOINT [ "./app" ]