package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type AppConfig struct {
	DefaultApiKey              string `json:"defaultApiKey"`
	EnableDBEventLogging       bool   `json:"enableDbEventLogging"`
	EnableChatterEventPrinting bool   `json:"enableChatterEventPrinting"`
}

func (conf *AppConfig) loadConfig() error {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		return errors.New("no config file present, please create config.json")
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		return err
	}

	return nil
}

func (conf *AppConfig) runConfigActions() {
	//see if the default key is in the db
	log.Println("Running config actions...")

	var apik ApiKey
	tx := db.First(&apik, "nickname = ?", "default")
	if tx.RowsAffected == 0 {
		//add the key
		log.Println("Default api key added")
		err := apik.addKeyToDB(conf.DefaultApiKey, "default")
		if err != nil {
			return
		}
	}

}
