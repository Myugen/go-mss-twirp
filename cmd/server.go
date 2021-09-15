package cmd

import (
	"net/http"

	"github.com/myugen/go-mss-twirp/internal/userserver"
	"github.com/myugen/go-mss-twirp/rpc/user"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server application command",
	RunE:  runServer,
}

func runServer(_ *cobra.Command, _ []string) error {
	server := userserver.New()
	twirpHandler := user.NewUserServer(server)

	return http.ListenAndServe(":8080", twirpHandler)
}
