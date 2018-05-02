package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd.PersistentFlags().StringP("config", "c", "", "the config file to use")

	rootCmd.PersistentFlags().StringP("workers", "w", "", "number of workers to start")
	rootCmd.PersistentFlags().StringP("debug", "d", "", "start in debug mode")
	rootCmd.PersistentFlags().IntP("port", "p", 0, "port to listen on")

	return &rootCmd
}
