package main

import (
	"flag"
)

type CMDParams struct {
	ConfigFilePath string
	IsPrintVersion bool
}

func parseCMDParams() (params *CMDParams) {
	params = &CMDParams{}

	flag.BoolVar(&params.IsPrintVersion, "v", false, "print version")
	flag.StringVar(&params.ConfigFilePath, "c", "config.json", "specify config file")

	flag.Parse()

	return params
}
