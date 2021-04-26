package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"
)

const ZIP_PASSWORD = "infected"

type SampleArgs struct {
	listLast       bool
	hashGet        string
	outputFile     string
	numList        int
	toJson         bool
	rawPrint       bool
	downloadSample bool
	queryTag       string
}

var sampleArgs SampleArgs = SampleArgs{}

var regex, _ = regexp.Compile("\n")

var sampleCmd = &cobra.Command{
	Use:   "sample [flags] sha256",
	Short: "Interact with samples in Malware Bazzar",
	Args: func(cmd *cobra.Command, args []string) error {
		if sampleArgs.listLast || sampleArgs.queryTag != "" {
			return nil
		}

		if len(args) != 0 {
			sampleArgs.hashGet = args[0]
			return nil
		}

		return errors.New("You need to pass at least the sample hash, but you can normally list")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if sampleArgs.rawPrint {
			bluePrint = color.New(color.FgWhite).PrintfFunc()
			redPrint = color.New(color.FgWhite).PrintfFunc()
		}

		if sampleArgs.hashGet != "" {
			if sampleArgs.downloadSample {
				fmt.Printf("Downloading %s...\n", sampleArgs.hashGet)
				sampleData, err := abuse.GetSample(sampleArgs.hashGet)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error on get sample: %v\n", err)
					return
				}
				unpacked, fileName := utils.Unzip(sampleData, ZIP_PASSWORD)

				var outputFile string

				if sampleArgs.outputFile != "" {
					outputFile = sampleArgs.outputFile
				} else {
					outputFile = fileName
				}

				utils.SaveFile(unpacked, outputFile)

			} else {
				rawJson, sampleQuery := abuse.QuerySampleInfo(sampleArgs.hashGet)
				if len(sampleQuery.Data) > 0 {
					if !sampleArgs.toJson {
						dumpSample(&sampleQuery.Data[0])
						return
					}
					fmt.Println(rawJson)
				}

			}

			return
		}

		if sampleArgs.listLast || sampleArgs.queryTag != "" {
			filenameSize := 16
			var latestSamples abuse.Response

			if sampleArgs.listLast {
				bluePrint("Loading last %d entries...\n", sampleArgs.numList)
				latestSamples = abuse.GetLatestSamples(sampleArgs.numList)
				bluePrint("Last %d entries:\n", len(latestSamples.Data))
			} else {
				bluePrint("Querying last entries of tag:%s\n", sampleArgs.queryTag)
				latestSamples = abuse.GetSampleByTag(sampleArgs.queryTag)
			}

			if len(latestSamples.Data) == 0 {
				redPrint("No samples found!\n")
				return
			}

			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 8, 8, 0, '\t', 0)
			defer w.Flush()

			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "Sha256", "Filename", "Filesize", "Filetype")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "--------", "--------", "--------", "--------")
			for _, sampleInfo := range latestSamples.Data {
				if len(sampleInfo.FileName) > filenameSize {
					sampleInfo.FileName = sampleInfo.FileName[:filenameSize] + "..."
				}

				fmt.Fprintf(w, "\n %s\t%s\t%d Kb\t%s", sampleInfo.Sha256Hash, sampleInfo.FileName, sampleInfo.FileSize, sampleInfo.FileType)
			}
			fmt.Println()
			return
		}

		cmd.Help()
	},
}

func init() {
	sampleCmd.Flags().BoolVarP(&sampleArgs.listLast, "list-last", "l", false, "List last 100 entries in Malware Bazzar")
	sampleCmd.Flags().StringVarP(&sampleArgs.outputFile, "output", "o", "", "Output sample path")
	sampleCmd.Flags().StringVarP(&sampleArgs.queryTag, "tag", "t", "", "Query a sample list by tag, eg: revil")
	sampleCmd.Flags().BoolVarP(&sampleArgs.toJson, "json", "j", false, "Output info in json format")
	sampleCmd.Flags().BoolVarP(&sampleArgs.rawPrint, "raw", "r", false, "Output without colors")
	sampleCmd.Flags().BoolVarP(&sampleArgs.downloadSample, "get", "g", false, "Download sample")

	sampleArgs.numList = 100
}

func dumpSample(sample *abuse.SampleInfo) {
	structFields := reflect.TypeOf(*sample)
	structValues := reflect.ValueOf(*sample)
	bluePrint("%s:\n", sample.Sha256Hash)
	dumpReflectedStruct(structFields, structValues, 1)
}
