package main

import (
	"errors"
	"log"
)

//In order for data to come in, inlet inlets need to be registered. Their job
//is to subscribe to their corresponding services and send data to the incoming
//channel which will then figure out where to route the data next

type InletModule struct {
	ID            string
	Name          string
	ConfigOptions map[string]interface{}
	configured    bool
	Publisher     func(map[string]interface{}, string)
}

func (m *InletModule) Configure(options map[string]interface{}) {
	log.Println("InletModule", m.Name, "configured.")
	m.ConfigOptions = options
	m.configured = true
}

func (m *InletModule) Start() error {
	if !m.configured {
		return errors.New("inlet module not configured")
	}

	go m.Publisher(m.ConfigOptions, m.ID)
	log.Println("InletModule", m.Name, "started.")
	return nil
}
