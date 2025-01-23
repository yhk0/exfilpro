package cmd

import (
	"exfilpro/util"

	"github.com/spf13/cobra"
)

var url string
var filter string

var webCmd = &cobra.Command{
	Use:   "web [url]",
	Short: "Analyzes leaks in a URL",
	Long:  "This command scans a URL for potential leaks of sensitive data.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		filter, _ := cmd.Flags().GetString("filter")

		util.AnalyzeWeb(url, filter) // Chamar a função de análise para web
	},
}

func init() {
	webCmd.Flags().StringP("filter", "f", "", "Filter results by data type (email, password, API Key, reCAPTCHA)")
}
