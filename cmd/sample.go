package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
	// "github.com/fatih/color"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"

)

const ZIP_PASSWORD = "infected"

type SampleArgs struct {
	listLast bool
	hashGet string
	sampleInfo bool
}

var sampleArgs SampleArgs = SampleArgs{}

var sampleCmd = &cobra.Command{
	Use: "sample",
	Short: "Interact with samples in Malware Bazzar",
	Run: func (cmd *cobra.Command, args []string) {

		cmd.Help()
	},	
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Download Malware Bazzar samples using your criteria",
	Run: func (cmd *cobra.Command, args []string) {
		if sampleArgs.hashGet != "" {
			if sampleArgs.sampleInfo {
				sampleQuery := abuse.QuerySampleInfo(sampleArgs.hashGet)
				if len(sampleQuery.Data) > 0 {
					fmt.Printf("%s:\n", sampleArgs.hashGet)
					dumpSample(&sampleQuery.Data[0])
				}

			} else {
				fmt.Printf("Downloading %s\n", sampleArgs.hashGet)
				sampleData, err := abuse.GetSample(sampleArgs.hashGet)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error on get sample: %v\n", err)
					return
				}
				//TODO unpack the zip
				utils.SaveFile(sampleData, sampleArgs.hashGet)
			}


			return
		}

		cmd.Help()
	},	
}

func init() {
	sampleCmd.Flags().BoolVarP(&sampleArgs.listLast, "list-last", "l", false, "List last entries in Malware Bazzar")
	sampleCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&sampleArgs.hashGet, "hash", "H", "", "Get sample by sha256 hash")
	getCmd.Flags().BoolVarP(&sampleArgs.sampleInfo, "info", "i", false, "Get sample info")
}

func dumpSample(sample *abuse.SampleInfo) {
	structFields := reflect.TypeOf(*sample)
	structValues := reflect.ValueOf(*sample)

	num := structFields.NumField()

	for i := 0; i < num ; i++ {
		field := structFields.Field(i)
		value := structValues.Field(i)

		switch value.Kind() {
		case reflect.String:
			//v := value.String()
			fmt.Printf("\t%s: %s\n", field.Name, value)
		}
	}

}