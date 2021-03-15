package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/aandersonl/bazzar/pkg/abuse"
)

type SampleArgs struct {
	listLast bool
	hashGet string
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
			fmt.Println(abuse.GetSample(sampleArgs.hashGet))
		}
	},	
}

func init() {
	sampleCmd.Flags().BoolVarP(&sampleArgs.listLast, "list-last", "l", false, "List last entries in Malware Bazzar")
	sampleCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&sampleArgs.hashGet, "hash", "H", "", "Get sample by hash: sha1, sha256, imphash, tlsh, telfhash")
}