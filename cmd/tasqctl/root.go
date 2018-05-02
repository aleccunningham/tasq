package tasqctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "example",
		Run: run,
	}
	rootCmd.Flags().IntP("port", "p", 0, "Port for workers")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("--- here ---")
}
