package cmd

import (
	"exfilpro/util"
	"log"

	"github.com/spf13/cobra"
)

var filePath string

var fileCmd = &cobra.Command{
	Use:   "file [path]",
	Short: "Scans local files for sensitive data",
	Long:  "File Mode: Performs analysis of local files to search for metadata and sensitive data.",
	Run: func(cmd *cobra.Command, args []string) {
		if filePath == "" {
			log.Fatal("Error: You need to specify the file path using the flag -p --path")
		}
		util.AnalyzeFile(filePath)
	},
}

func init() {
	fileCmd.Flags().StringVarP(&filePath, "path", "p", "", "Path of the file to be parsed (required)")
}
