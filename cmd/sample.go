package cmd

import (
	"fmt"
	"os"
	"regexp"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/fatih/color"

	"github.com/aandersonl/bazzar/pkg/abuse"
	"github.com/aandersonl/bazzar/pkg/utils"

)

const ZIP_PASSWORD = "infected"

type SampleArgs struct {
	listLast bool
	hashGet string
	sampleInfo bool
	outputFile string

	toJson bool
}

var sampleArgs SampleArgs = SampleArgs{}


// colors print
var  (
	bluePrint func(format string, a ...interface{}) = color.New(color.FgBlue).PrintfFunc()
	redPrint  func(format string, a ...interface{}) = color.New(color.FgRed).PrintfFunc()
)

var regex, _ = regexp.Compile("\n")

var sampleCmd = &cobra.Command{
	Use: "sample",
	Short: "Interact with samples in Malware Bazzar",
	Run: func (cmd *cobra.Command, args []string) {
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

		cmd.Help()
	},	
}

func init() {
	sampleCmd.Flags().BoolVarP(&sampleArgs.listLast, "list-last", "l", false, "List last entries in Malware Bazzar")

	sampleCmd.Flags().StringVarP(&sampleArgs.hashGet, "hash", "H", "", "Get sample by sha256 hash")
	sampleCmd.Flags().StringVarP(&sampleArgs.outputFile, "output", "o", "", "Output sample path")

	sampleCmd.Flags().BoolVarP(&sampleArgs.sampleInfo, "info", "i", false, "Get sample info")
	sampleCmd.Flags().BoolVarP(&sampleArgs.toJson, "json", "j", false, "Output info in json format")
}


func dumpSample(sample *abuse.SampleInfo) {
	structFields := reflect.TypeOf(*sample)
	structValues := reflect.ValueOf(*sample)
	bluePrint("%s:\n", sample.Sha256Hash)
	dumpReflectedStruct(structFields, structValues, 1)
}



func dumpReflectedStruct(structFields reflect.Type, structValues reflect.Value, level int) {
	var tab string
	for i := 0; i < level; i++ {
		tab += "\t"
	}

	num := structFields.NumField()
	for i := 0; i < num ; i++ {
		field := structFields.Field(i)
		value := structValues.Field(i)
		switch value.Kind() {
		case reflect.String:
			valueStr := value.Interface().(string)
			if valueStr == "" {
				continue
			}
			cleanValue := regex.ReplaceAllString(valueStr, " ")
			bluePrint("%s%s: ",tab, field.Name)
			redPrint("%s\n", cleanValue)
		
		case reflect.Int:
			bluePrint("%s%s: ", tab, field.Name)	
			redPrint("%d\n", value)
		case reflect.Slice:
			bluePrint("%s%s: ", tab, field.Name)
			numElements := value.Len()
			for j := 0; j < numElements; j++ {
				el := value.Index(j)
				switch el.Kind() {
				case reflect.String:
					redPrint("%s", el)
					if j != numElements - 1 {
						redPrint(",")
					}
				case reflect.Struct:
					stFields := reflect.TypeOf(el.Interface())
					stValues := reflect.ValueOf(el.Interface())
					fmt.Println()
					dumpReflectedStruct(stFields, stValues, level + 1)
				}

			}
			fmt.Println()

		case reflect.Struct:
			fmt.Println()
			stFields := reflect.TypeOf(value.Interface())
			stValues := reflect.ValueOf(value.Interface())
			
			bluePrint("%s%s:\n",tab, field.Name)
			dumpReflectedStruct(stFields, stValues, level + 1)
		}
	}
}