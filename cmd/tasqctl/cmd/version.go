package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tasq",
	Long:  "All software has versions. This is tasq's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tasq v0.1 -- HEAD")
	},
}
