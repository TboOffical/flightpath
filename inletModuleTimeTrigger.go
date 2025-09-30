package main

import (
	"fmt"
	"log"
	"time"
)

func timeTriggerPublisher(opts map[string]interface{}, id string) {
	fmt.Println(opts)

	if opts["delay"] == nil {
		log.Println("delay option is required")
		return
	}

	delay := opts["delay"].(int)

	for {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		incoming <- IncomingMessage{
			From:    id,
			Message: "{\"message\": \"tick\"}",
		}
	}
}

func newTimeTriggerInlet(id string) *InletModule {
	return &InletModule{
		ID:        id,
		Name:      "time_trigger",
		Publisher: timeTriggerPublisher,
	}
}

func registerTimeTrigger() {
	e := newDocsEntry("time_trigger", NodeTypeInlet)
	e.AddParam("Delay", "int", "Time between ticks", "delay")
	AppDocs = append(AppDocs, e)
}
