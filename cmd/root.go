package cmd

import (
	"fmt"
	"os"
	//homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "bazzar",
	Short: "Interact with abuse.ch intel feed",
}
 
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

  
func init() {
	//cobra.OnIti
	rootCmd.AddCommand(sampleCmd)
	rootCmd.AddCommand(urlCmd)
	//fmt.Println(rootCmd)
}


  
