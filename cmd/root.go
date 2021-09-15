package cmd

import (
	"fmt"
	"os"

	"github.com/myugen/go-mss-twirp/constants"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   constants.AppLabel,
		Short: fmt.Sprintf("Root command for %s", constants.AppName),
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
