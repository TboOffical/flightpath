package main

import (
	"log"

	"gorm.io/gorm"
)

const (
	EventTypeIncoming = 0
	EventTypeLogging  = 1
)

type Event struct {
	gorm.Model
	NiceTitle string
	Type      int
	Data      string
}

// e logs a string to the event log under the logging category
func e(message string) {
	chatter <- message
	if config.EnableDBEventLogging {
		db.Create(&Event{
			NiceTitle: "System Message",
			Type:      EventTypeLogging,
			Data:      message,
		})
	}
}

func pushEvent(title string, Type int, data string) {
	if config.EnableDBEventLogging {
		db.Create(&Event{
			NiceTitle: title,
			Type:      Type,
			Data:      data,
		})
	}
}

func printChatter() {
	for {
		select {
		case msg := <-chatter:
			if config.EnableChatterEventPrinting {
				log.Println("EC|", msg)
			}
		}
	}
}
