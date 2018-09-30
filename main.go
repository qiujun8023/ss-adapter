package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var version = "0.0.1"
var log = logrus.New()

func main() {
	// 获取配置文件
	params := parseCMDParams()
	config := loadConfig(params.ConfigFilePath, params.Config)
	if params.IsPrintVersion {
		log.Infof("shadowsocks adapter version %s", version)
		os.Exit(0)
	}

	conn, err := getConnect(config)
	if err != nil {
		log.Fatalf("error during listen: %s", err)
	}
	defer conn.Close()

	go doSomething(conn, config)
	go receiveMessage(conn)

	for {
		sendPing(conn)
		time.Sleep(5 * time.Second)
	}
}
