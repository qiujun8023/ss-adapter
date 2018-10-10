package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonConfig struct {
	ManagerServer string `json:"manager_server"`
	ManagerPort   int    `json:"manager_port"`
	APIURL        string `json:"api_url"`
	NodeID        string `json:"node_id"`
	NodeToken     string `json:"node_token"`
	SyncInterval  int    `json:"sync_interval"`
}

func loadConfig(path string) (config *jsonConfig, err error) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config = &jsonConfig{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return
}
