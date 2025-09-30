package main

import (
	"errors"
	"fmt"
	"log"
)

type OutletModule struct {
	Name          string
	ID            string
	ListenFrom    []string
	ConfigOptions map[string]interface{}
	Publish       func(string, map[string]interface{}) error //content from node/config options
	configured    bool
}

func (m *OutletModule) Configure(options map[string]interface{}) {
	log.Println("OutletModule", m.Name, "configured.")
	m.ConfigOptions = options
	m.configured = true
}

func (m *OutletModule) Start() error {
	if !m.configured {
		return errors.New("outlet module not configured")
	}

	go func() {
		for {
			select {
			case msg := <-incoming:
				//Check if the message applies to this node
				if in_str(msg.From, m.ListenFrom) {
					//If it does, publish the message with the publish function

					err = m.Publish(msg.Message, m.ConfigOptions)
					if err != nil {
						log.Println(err)
						e(fmt.Sprint("Error publishing message ", err))
						return
					}
				}
			}
		}
	}()

	log.Println("OutletModule ", m.Name, "started.")
	return nil
}
