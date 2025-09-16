package main

import "fmt"

func timeTriggerPublisher(opts map[string]interface{}, id string) {
	fmt.Println(opts)
	for {

	}
}

func newTimeTriggerInlet(id string) *InletModule {
	return &InletModule{
		ID:        id,
		Name:      "time_trigger",
		Publisher: timeTriggerPublisher,
	}
}
