package main

import (
	"net"
	"strconv"
)

type ssUser struct {
	UserID      int    `json:"userId"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	IsLocked    bool   `json:"isLocked"`
	IsDeleted   bool   `json:"isDeleted"`
}

type ssTraffic struct {
	UserID   int `json:"userId"`
	FlowUp   int `json:"flowUp"`
	FlowDown int `json:"flowDown"`
}

var lastUser map[string]ssUser
var lastTraffic map[string]ssTraffic

func stopOrStartServer(conn *net.UDPConn, user *ssUser) {
	_, isRun := lastTraffic[strconv.Itoa(user.Port)]

	if isRun && user.IsDeleted {
		log.Infof("stop server at port [%d] reason: deleted", user.Port)
		sendRemoveAccount(conn, user)
	} else if isRun && user.IsLocked {
		log.Infof("stop server at port [%d] reason: disable", user.Port)
		sendRemoveAccount(conn, user)
	} else if isRun && user.Password != user.OldPassword {
		log.Infof("stop server at port [%d] reason: password changed", user.Port)
		sendRemoveAccount(conn, user)
	} else if !isRun && !user.IsDeleted && !user.IsLocked {
		log.Infof("start server at port [%d] pass [%s]", user.Port, user.Password)
		sendAddAccount(conn, user)
	}
}

func syncUser(conn *net.UDPConn, config *jsonConfig) {
	if lastTraffic == nil {
		return
	} else if lastUser == nil {
		lastUser = make(map[string]ssUser)
	}

	data, err := fetchUsers(config)
	if err != nil {
		log.Errorf("error during fetch user: %s", err)
		return
	}

	// 初始化用户
	users := map[string]ssUser{}
	for _, user := range data {
		user.IsDeleted = false
		users[strconv.Itoa(user.Port)] = user
	}

	// 判断用户是否删除
	for port, user := range lastUser {
		_, isExist := users[port]
		if isExist == false {
			user.IsDeleted = true
			users[strconv.Itoa(user.Port)] = user
		}
	}

	for port, user := range users {
		user.OldPassword = lastUser[port].Password
		stopOrStartServer(conn, &user)
		if user.IsDeleted {
			delete(lastUser, strconv.Itoa(user.Port))
		} else {
			lastUser[strconv.Itoa(user.Port)] = user
		}
	}
}

func syncTraffic(config *jsonConfig, traffic map[string]ssTraffic) {
	data := []ssTraffic{}
	for port, item := range traffic {
		user, isExist := lastUser[port]
		if isExist == false {
			continue
		}
		item.UserID = user.UserID
		data = append(data, item)
	}

	lastTraffic = traffic

	if len(data) != 0 {
		uploadTraffic(config, data)
	}
}
