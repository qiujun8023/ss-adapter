package main

import (
	"flag"
)

type cmdParams struct {
	ConfigFilePath string
	IsPrintVersion bool
}

func parseCMDParams() (params *cmdParams) {
	params = &cmdParams{}

	flag.BoolVar(&params.IsPrintVersion, "v", false, "print version")
	flag.StringVar(&params.ConfigFilePath, "c", "config.json", "specify config file")

	flag.Parse()

	return params
}
