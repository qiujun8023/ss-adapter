package main

import (
	"flag"
)

type CMDParams struct {
	*Config
	ConfigFilePath string
	IsPrintVersion bool
}

func parseCMDParams() *CMDParams {
	var params CMDParams
	var config Config

	params.Config = &config
	flag.BoolVar(&params.IsPrintVersion, "v", false, "print version")
	flag.StringVar(&params.ConfigFilePath, "c", "config.json", "specify config file")
	flag.StringVar(&config.ManagerServer, "s", "", "ip address of your manager server")
	flag.IntVar(&config.ManagerPort, "p", 0, "port number of your manager server")
	flag.StringVar(&config.APIURL, "a", "", "ss-panel api url, eg: http://ss.example.com")
	flag.IntVar(&config.NodeID, "i", 0, "ss-panel node id")
	flag.StringVar(&config.NodeToken, "k", "", "ss-panel node token")

	flag.Parse()

	return &params
}
