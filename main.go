package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var version = "0.0.1"
var log = logrus.New()

func main() {
	params := parseCMDParams()
	if params.IsPrintVersion {
		log.Infof("shadowsocks adapter version %s", version)
		os.Exit(0)
	}

	config, err := loadConfig(params.ConfigFilePath)
	if err != nil {
		log.Fatalf("error during load config, %v", err)
	}

	conn, err := getConnect(config)
	if err != nil {
		log.Fatalf("error during listen: %s", err)
	}
	defer conn.Close()

	go syncUser(conn, config)
	go handleMessage(conn)

	for {
		sendPing(conn)
		time.Sleep(time.Duration(config.SyncInterval) * time.Second)
	}
}
