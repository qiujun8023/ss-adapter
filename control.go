package main

import "net"

func doSomething(conn *net.UDPConn, config *Config) {
	users, err := fetchUsers(config)
	log.Infof("users %s %s", err, users)
	account := Account{54321, "12345"}
	for i := 0; i < 1; i++ {
		if i%2 == 0 {
			sendAddAccount(conn, &account)
		} else {
			sendRemoveAccount(conn, &account)
		}
	}
}
