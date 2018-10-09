package main

import "net"

// func doSomething(conn *net.UDPConn, config *Config) {
// 	users, err := fetchUsers(config)
// 	log.Infof("users %s %s", err, users)
// 	account := Account{54321, "12345"}
// 	for i := 0; i < 1; i++ {
// 		if i%2 == 0 {
// 			sendAddAccount(conn, &account)
// 		} else {
// 			sendRemoveAccount(conn, &account)
// 		}
// 	}
// }

func syncUser(conn *net.UDPConn, config *Config) bool {
	users, err := fetchUsers(config)
	if err != nil {
		log.Errorf("error during fetch user: %s", err)
		return false
	}

	sendRemoveAccount(conn, &Account{54321, ""})
	for _, user := range users {
		sendAddAccount(conn, &Account{user.Port, user.Password})
	}

	return true
}
