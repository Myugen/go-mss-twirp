# syntax=docker/dockerfile:1
ARG GOLANG_VERSION=1.17-alpine

# Builder image
FROM golang:${GOLANG_VERSION} as builder

WORKDIR /go/src/github.com/myugen/go-mss-twirp/

RUN echo 'Installing system dependencies' \
    && apk update \
    && apk add protoc

COPY . ./

ENV CGO_ENABLED=0 \
    GO111MODULE="on" \
    GOOS=linux

RUN echo 'Downloading go.mod dependencies' \
    && go mod download
RUN echo 'Installing tools from tools.go' \
    && cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %
RUN echo 'Generating gRPC services' \
    && go generate ./...

RUN go build -a -o app .
RUN chmod +x app


# Generate clean, final image for end users
FROM alpine:latest as app

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/myugen/go-mss-twirp/app ./

ENTRYPOINT [ "./app" ]

CMD [ "--help" ]