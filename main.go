package main

import (
	// "fmt"
	//"os"

	//"github.com/aandersonl/bazzar/pkg/abuse"

	"github.com/aandersonl/bazzar/cmd"
)

func main() {
	//fmt.Println(abuse.GetLatestSamples())
	//fmt.Println(abuse.QuerySampleInfo(os.Args[1]))
	//fmt.Println(abuse.QuerySignature("TrickBot", 50))

	cmd.Execute()
}
