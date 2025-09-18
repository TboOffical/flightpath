package main

import (
	"errors"
	"log"
)

//Modifiers listen to the incoming channel for messages from nodes that they are
//listening to. Once they get these messages they modify them and write them back
//to the incoming channel for other nodes to be triggered off of

type ModifierModule struct {
	Name          string
	ID            string
	Task          string
	ListenFrom    []string
	ConfigOptions map[string]interface{}
	Modify        func(string, map[string]interface{}, string) string
	Verify        func(map[string]interface{}, string) error
	configured    bool
}

func (m *ModifierModule) Configure(options map[string]interface{}) {
	log.Println("ModifierModule", m.Name, "configured.")
	m.ConfigOptions = options
	m.configured = true
}

func (m *ModifierModule) Start() error {
	if !m.configured {
		return errors.New("modifier module not configured")
	}

	//check if the configuration is correct
	if err := m.Verify(m.ConfigOptions, m.Task); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case msg := <-incoming:
				//Check if the message applies to this node
				if in_str(msg.From, m.ListenFrom) {
					//If it does modify the message and send it back to the incoming ch

					newContent := m.Modify(msg.Message, m.ConfigOptions, m.Task)
					d("Modifier ", m.ID, " ran with final value of ", newContent)
					incoming <- IncomingMessage{
						From:    m.ID,
						Message: newContent,
					}
				}
			}
		}
	}()

	log.Println("ModifierModule", m.Name, "started.")
	return nil
}
