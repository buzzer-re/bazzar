package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"
)

const ZIP_PASSWORD = "infected"

type SampleArgs struct {
	listLast   bool
	hashGet    string
	sampleInfo bool
	outputFile string
	numList    int
	toJson     bool
}

var sampleArgs SampleArgs = SampleArgs{}

var regex, _ = regexp.Compile("\n")

var sampleCmd = &cobra.Command{
	Use:   "sample [flags] sha256",
	Short: "Interact with samples in Malware Bazzar",
	Args: func(cmd *cobra.Command, args []string) error {
		if sampleArgs.listLast {
			return nil
		}

		if len(args) != 0 {
			sampleArgs.hashGet = args[0]
			return nil
		}

		return errors.New("You need to pass at least the sample hash, but you can normally list")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if sampleArgs.hashGet != "" {
			if sampleArgs.sampleInfo {
				rawJson, sampleQuery := abuse.QuerySampleInfo(sampleArgs.hashGet)
				if len(sampleQuery.Data) > 0 {
					if !sampleArgs.toJson {
						dumpSample(&sampleQuery.Data[0])
						return
					}
					fmt.Println(rawJson)
				}

			} else {
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
			}

			return
		}

		if sampleArgs.listLast {
			bluePrint("Loading last %d entries...\n", sampleArgs.numList)
			latestSamples := abuse.GetLatestSamples(sampleArgs.numList)
			bluePrint("Last %d entries:\n", len(latestSamples.Data))
			for _, sampleInfo := range latestSamples.Data {
				redPrint("%s - %s\n", sampleInfo.Sha256Hash, sampleInfo.FileName)
			}

			return
		}

		cmd.Help()
	},
}

func init() {
	sampleCmd.Flags().BoolVarP(&sampleArgs.listLast, "list-last", "l", false, "List last 100 entries in Malware Bazzar")

	sampleCmd.Flags().StringVarP(&sampleArgs.outputFile, "output", "o", "", "Output sample path")

	sampleCmd.Flags().BoolVarP(&sampleArgs.sampleInfo, "info", "i", false, "Get sample info")
	sampleCmd.Flags().BoolVarP(&sampleArgs.toJson, "json", "j", false, "Output info in json format")

	sampleArgs.numList = 100
}

func dumpSample(sample *abuse.SampleInfo) {
	structFields := reflect.TypeOf(*sample)
	structValues := reflect.ValueOf(*sample)
	bluePrint("%s:\n", sample.Sha256Hash)
	dumpReflectedStruct(structFields, structValues, 1)
}
