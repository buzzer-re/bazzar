package cmd

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"
	"github.com/spf13/cobra"
)

type UrlArgs struct {
	Url    string
	ToJson bool
	Check  bool
	List   bool
	Num    int
}

var urlArgs UrlArgs

var urlCmd = &cobra.Command{
	Use:   "url url|host",
	Short: "Query urlhaus information",
	Args: func(cmd *cobra.Command, args []string) error {
		if urlArgs.List {
			return nil
		}

		if len(args) != 0 {
			urlArgs.Url = args[0]
			return nil
		}
		return errors.New("You need to pass a url or host, but you can normally list")
	},
	Run: func(cmd *cobra.Command, args []string) {

		if urlArgs.List {
			if !urlArgs.ToJson {
				bluePrint("Listing the last %d new urls...\n", urlArgs.Num)
			}

			rawJson, lastSamples := abuse.QueryLast(urlArgs.Num)
			if urlArgs.ToJson {
				fmt.Println(rawJson)
				return
			}
			dumpLastUrls(lastSamples)
			return
		}

		if urlArgs.Url != "" {
			var (
				hostResponse abuse.HostResponse
				urlResponse  abuse.URLResponse
				rawJson      string
				isUrl        bool
			)

			isUrl = utils.IsFullUrl(urlArgs.Url)

			if isUrl {
				rawJson, urlResponse = abuse.QueryUrl(urlArgs.Url)
			} else if utils.IsHost(urlArgs.Url) {
				urlArgs.Url = utils.CleanHost(urlArgs.Url)
				rawJson, hostResponse = abuse.QueryHost(urlArgs.Url)
			} else {
				log.Fatalln("Wrong url format")
			}

			if urlArgs.ToJson {
				fmt.Println(rawJson)
				return
			}

			hostExists := hostResponse.QueryStaus == "ok"
			urlExists := urlResponse.QueryStatus == "ok"

			if hostExists && urlExists {
				bluePrint("%s is  not listed in urlhaus database\n", urlArgs.Url)
			} else {
				if hostExists {
					dumpHostSuspicious(hostResponse)
				} else {
					dumpUrlSuspicious(urlResponse)
				}
			}

			return
		}

		cmd.Help()
	},
}

func dumpLastUrls(lastUrls abuse.LastUrls) {
	bluePrint("URLs:\n")

	for _, url := range lastUrls.Urls {
		redPrint("\t%s - %s (%s)\n", url.URL, url.Host, url.Threat)

	}
}

func dumpHostSuspicious(host abuse.HostResponse) {
	structFields := reflect.TypeOf(host)
	structValues := reflect.ValueOf(host)
	bluePrint("HOST: ")
	redPrint("%s\n", host.Host)

	dumpReflectedStruct(structFields, structValues, 1)

}

func dumpUrlSuspicious(url abuse.URLResponse) {
	bluePrint("URL: ")
	redPrint("%s\n", url.URL)

	structFields := reflect.TypeOf(url)
	structValues := reflect.ValueOf(url)
	dumpReflectedStruct(structFields, structValues, 1)

}

func init() {
	urlCmd.Flags().StringVarP(&urlArgs.Url, "url", "u", "", "Get URL information")
	urlCmd.Flags().BoolVarP(&urlArgs.ToJson, "json", "j", false, "Output in JSON format")
	urlCmd.Flags().BoolVarP(&urlArgs.List, "list", "l", false, "List new urls")
	urlCmd.Flags().IntVarP(&urlArgs.Num, "num", "n", 20, "Number of urls to list")
}
