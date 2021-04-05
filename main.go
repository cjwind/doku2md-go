package main

import (
	"flag"
	"fmt"
	"github.com/cjwind/doku2md-go/doku2md"
	"io/ioutil"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("doku2md-go [filepath]")
		return
	}

	filepath := flag.Arg(0)

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	converter := doku2md.Converter{}
	output := converter.DokuToMd(string(content))
	fmt.Print(output)
}
