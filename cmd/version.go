package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of birthday-reminder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("birthday-reminder version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}