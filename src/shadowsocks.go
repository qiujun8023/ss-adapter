package main

import (
	"encoding/json"
	"net"
	"strings"
)

func getConnect(config *jsonConfig) (*net.UDPConn, error) {
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(config.ManagerServer), Port: config.ManagerPort}
	return net.DialUDP("udp", srcAddr, dstAddr)
}

func sendMessage(conn *net.UDPConn, message string) {
	conn.Write([]byte(message))
	log.Debugf("send manage message: %s", message)
}

func sendPing(conn *net.UDPConn) {
	sendMessage(conn, "ping")
}

func sendAddAccount(conn *net.UDPConn, user *ssUser) {
	res, err := json.Marshal(map[string]interface{}{
		"server_port": user.Port,
		"password":    user.Password,
	})
	if err != nil {
		log.Errorf("error during json marshal: %s", err)
	}
	sendMessage(conn, "add: "+string(res))
}

func sendRemoveAccount(conn *net.UDPConn, user *ssUser) {
	res, err := json.Marshal(map[string]int{
		"server_port": user.Port,
	})
	if err != nil {
		log.Errorf("error during json marshal: %s", err)
	}
	sendMessage(conn, "remove: "+string(res))
}

func handleTrafficMessage(conn *net.UDPConn, config *jsonConfig, data []byte) {
	var trafficMap map[string]int
	err := json.Unmarshal(data, &trafficMap)
	if err != nil {
		log.Errorf("error during json unmarshal: %s", err)
		return
	}

	traffic := map[string]ssTraffic{}
	for port, flow := range trafficMap {
		traffic[port] = ssTraffic{0, 0, flow}
	}

	syncTraffic(config, traffic)
	syncUser(conn, config)
}

func handleMessage(conn *net.UDPConn, config *jsonConfig) {
	data := make([]byte, 1506)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Errorf("error during read: %s", err)
		}

		log.Debugf("receive %s from <%s>", data[:n], remoteAddr)

		if strings.HasPrefix(string(data[:n]), "stat:") {
			handleTrafficMessage(conn, config, data[6:n])
		}
	}
}
