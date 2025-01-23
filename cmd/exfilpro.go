package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exfilpro",
	Short: "A tool for data leak analysis",
	Long: `ExfilPro is a CLI tool designed for analyzing web content or files 
for potential sensitive information leaks. 
It supports filtering results by specific types of leaks and provides detailed output, 
including line numbers and file paths.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'exfilpro --help' to see the available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(fileCmd)
}
