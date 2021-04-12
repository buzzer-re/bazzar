package cmd

import (
	"fmt"
	"log"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"
	"github.com/spf13/cobra"
)

type UrlArgs struct {
	Url    string
	ToJson bool
}

var urlArgs UrlArgs

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Query urlhaus information",
	Run: func(cmd *cobra.Command, args []string) {

		if urlArgs.Url != "" {
			var (
				hostResponse abuse.HostResponse
				urlResponse  abuse.URLResponse
				rawJson      string
			)

			if utils.IsHost(urlArgs.Url) {
				rawJson, hostResponse = abuse.QueryHost(urlArgs.Url)
			} else if utils.IsFullUrl(urlArgs.Url) {
				rawJson, urlResponse = abuse.QueryUrl(urlArgs.Url)
			} else {
				log.Fatalln("Wrong url format")
			}

			if urlArgs.ToJson {
				fmt.Println(rawJson)
			} else {
				fmt.Println(hostResponse)
				fmt.Println(urlResponse)
			}
		}

		cmd.Help()
	},
}

func init() {
	urlCmd.Flags().StringVarP(&urlArgs.Url, "url", "u", "", "Get URL information")
	urlCmd.Flags().BoolVarP(&urlArgs.ToJson, "json", "j", false, "Output in JSON format")
}
