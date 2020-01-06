package cmd

import (
	"bytes"

	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat the change log file",
	Long:  "Reformats changelog input following keepachangelog.com spec",
	Run: func(cmd *cobra.Command, args []string) {
		changelog := parser.Parse(inputFile)

		var buf bytes.Buffer
		changelog.Render(&buf)

		outputFile.ReadFrom(&buf)
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
