package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TranslateRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Text   string `json:"text"`
}

type Configuration struct {
	Username string
	Password string
}

func Translate(source, target, text string) []byte {
	translateRequest := TranslateRequest{
		Source: source,
		Target: target,
		Text:   text,
	}
	requestData, _ := json.Marshal(translateRequest)

	translateUrl := "https://gateway.watsonplatform.net/language-translator/api/v2/translate"
	req, err := http.NewRequest("POST", translateUrl, bytes.NewBuffer(requestData))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return []byte("")
	}

	config := getCredentials()

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return []byte("")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Copy: ", err)
		return []byte("")
	}
	return body
}

func getCredentials() Configuration {
	file, err := os.Open("./conf.json")
	if err != nil {
		log.Fatal("No conf.json found!")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Trouble reading config data")
	}

	return config
}
