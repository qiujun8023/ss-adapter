package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type Config struct {
	ManagerServer string `json:"manager_server"`
	ManagerPort   int    `json:"manager_port"`
	APIURL        string `json:"api_url"`
	NodeID        int    `json:"node_id"`
	NodeToken     string `json:"node_token"`
}

func loadConfigFromFile(path string) (config *Config, err error) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	config = &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return
}

func updateConfig(old, new *Config) {
	// Using reflection here is not necessary, but it's a good exercise.
	// For more information on reflections in Go, read "The Laws of Reflection"
	// http://golang.org/doc/articles/laws_of_reflection.html
	newVal := reflect.ValueOf(new).Elem()
	oldVal := reflect.ValueOf(old).Elem()

	// typeOfT := newVal.Type()
	for i := 0; i < newVal.NumField(); i++ {
		newField := newVal.Field(i)
		oldField := oldVal.Field(i)
		// log.Printf("%d: %s %s = %v\n", i,
		// typeOfT.Field(i).Name, newField.Type(), newField.Interface())
		switch newField.Kind() {
		case reflect.String:
			s := newField.String()
			if s != "" {
				oldField.SetString(s)
			}
		case reflect.Int:
			i := newField.Int()
			if i != 0 {
				oldField.SetInt(i)
			}
		}
	}
}

func loadConfig(configFilePath string, cmdParamsConfig *Config) *Config {
	config, err := loadConfigFromFile(configFilePath)
	if err != nil {
		config = cmdParamsConfig
		if !os.IsNotExist(err) {
			log.Fatalf("error reading %s: %v", configFilePath, err)
		}
	} else {
		updateConfig(config, cmdParamsConfig)
	}

	if config.ManagerPort == 0 {
		config.ManagerPort = 8839
	}

	if config.ManagerServer == "" {
		config.ManagerServer = "127.0.0.1"
	}

	return config
}
