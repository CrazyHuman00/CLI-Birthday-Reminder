package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute git merge recursively",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func RunCommand(name string, args []string) error {
	command := name
	for _, a := range args {
		command += " " + a
	}
	fmt.Printf("Running... [%s]\n", command)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// func AddCommand(name string, args []string) error {
// 	// AddCommand
// }

// func UpdateCommand(name string, args []string) error {
// 	// UpdateCommand
// }

// func DeleteCommand(name string, args []string) error {
// 	// DeleteCommand
// }
