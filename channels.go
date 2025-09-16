package main

type IncomingMessage struct {
	From    string //id
	Message string
}

var incoming chan IncomingMessage = make(chan IncomingMessage)
var chatter chan string = make(chan string)
