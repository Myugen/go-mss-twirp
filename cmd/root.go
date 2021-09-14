package cmd

import (
	"fmt"
	"github.com/myugen/go-mss-twirp/utils/constants"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFile  string
	userLicense string
	rootCmd     = &cobra.Command{
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
