package user

// Go generate directives are a convenient way to describe compilation of proto
// files. They let users run 'go generate ./...', no Makefile necessary.
//
// This particular particular one invokes protoc using '$GOPATH/src'.
//
// This is used to tell protoc where to look up .proto files, through
// --proto_path.
//
// It is also used to tell protoc where to put output generated files, through
// --twirp_out and --go_out.

//go:generate protoc --proto_path=$GOPATH/src --twirp_out=$GOPATH/src --go_out=$GOPATH/src github.com/myugen/go-mss-twirp/rpc/user/service.proto
