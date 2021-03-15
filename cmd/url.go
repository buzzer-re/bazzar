package cmd

import (
	"github.com/spf13/cobra"
)

type UrlArgs struct {

}

var urlCmd = &cobra.Command{
	Use: "url",
	Short: "Query urlhaus information",
	Run: func (cmd *cobra.Command, args []string) {

	},
}