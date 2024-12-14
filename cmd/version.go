package cmd

import (
	"github.com/naturelr/gitlab2ea/pkg/versions"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		versions.Print()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
