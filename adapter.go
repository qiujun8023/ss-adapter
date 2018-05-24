package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Account struct {
	Port     int    `json:"server_port"`
	Password string `json:"password"`
}

func getConnect(config *Config) (*net.UDPConn, error) {
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(config.ManagerServer), Port: config.ManagerPort}
	return net.DialUDP("udp", srcAddr, dstAddr)
}

func sendMessage(conn *net.UDPConn, message string) {
	// conn.Write([]byte("hello"))
	// data := make([]byte, 1506)
	conn.Write([]byte(message))
	fmt.Println("send manage message: ", message)
}

func sendPing(conn *net.UDPConn) {
	sendMessage(conn, "ping")
}

func sendAddAccount(conn *net.UDPConn, account *Account) {
	res, err := json.Marshal(account)
	if err != nil {
		fmt.Println("json err:", err)
	}
	sendMessage(conn, "add: "+string(res))
}

func sendRemoveAccount(conn *net.UDPConn, account *Account) {
	res, err := json.Marshal(struct {
		*Account
		Password string `json:"password,omitempty"`
	}{
		Account: account,
	})

	if err != nil {
		fmt.Println("json err:", err)
	}
	sendMessage(conn, "remove: "+string(res))
}

func receiveMessage(conn *net.UDPConn) {
	data := make([]byte, 1506)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}

		fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}
