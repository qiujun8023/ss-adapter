package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 获取配置文件
	params := parseCMDParams()
	config := loadConfig(params.ConfigFilePath, params.Config)
	if params.IsPrintVersion {
		fmt.Println("shadowsocks adapter version", VERSION)
		os.Exit(0)
	}

	conn, err := getConnect(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during listen: %s\n", err)
	}
	defer conn.Close()

	go doSomething(conn, config)
	go receiveMessage(conn)

	for {
		sendPing(conn)
		time.Sleep(5 * time.Second)
	}
}
