package userserver

import (
	"context"

	"github.com/twitchtv/twirp"

	pb "github.com/myugen/go-mss-twirp/rpc/user"
)

type server struct{}

func (s *server) Create(_ context.Context, in *pb.UserIn) (*pb.UserOut, error) {
	if in.Username == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "The username field can not be empty")
	}
	if in.Name == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "The name field can not be empty")
	}
	if in.Mail == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "The mail field can not be empty")
	}

	return &pb.UserOut{
		Id:       "new",
		Username: in.Username,
		Mail:     in.Mail,
		Name:     in.Name,
	}, nil
}

func New() pb.User {
	return &server{}
}
