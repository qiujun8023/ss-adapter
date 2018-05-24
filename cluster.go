package main

import "net"

func doSomething(conn *net.UDPConn, config *Config) {
	account := Account{54321, "12345"}
	for i := 0; i < 1; i++ {
		if i%2 == 0 {
			sendAddAccount(conn, &account)
		} else {
			sendRemoveAccount(conn, &account)
		}
	}
}
