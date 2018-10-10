package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func doRequest(config *jsonConfig, req *http.Request, res interface{}) error {
	req.Header.Add("Node-Token", config.NodeToken)

	client := &http.Client{
		Timeout: time.Duration(15 * time.Second),
	}

	log.Infof("http %s request to %s", req.Method, req.URL)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("error during request: %s", err)
		return err
	}
	defer resp.Body.Close()

	log.Infof("response status %s", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error during read body: %s", err)
		return err
	}

	if resp.StatusCode >= 400 {
		log.Errorf("error response %s", string(body))
		return errors.New("response status is " + resp.Status)
	}

	log.Debugf("response body %s", string(body))
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Errorf("error during parse json: %s", err)
		return err
	}

	return nil
}

func fetchUsers(config *jsonConfig) (users []ssUser, err error) {
	url := config.APIURL + "/api/nodes/" + config.NodeID + "/users"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return users, doRequest(config, req, &users)
}

func uploadTraffic(config *jsonConfig, data []ssTraffic) error {
	url := config.APIURL + "/api/nodes/" + config.NodeID + "/traffic"
	jsonBody, _ := json.Marshal(data)
	bodyReader := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	req.Header.Add("Content-Type", "application/json")

	var result map[string]interface{}
	return doRequest(config, req, &result)
}
