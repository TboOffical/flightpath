package main

import (
	"errors"
	"fmt"
)

const (
	NodeTypeInlet    = 0
	NodeTypeOutlet   = 1
	NodeTypeModifier = 2
)

type Node struct {
	ID         string
	Type       int
	Module     string
	Task       string
	ListenFrom string
	Config     map[string]interface{}
}

func newNodeFromInterface(ni interface{}) (*Node, error) {
	node := ni.(map[string]interface{})

	if node["type"] == nil ||
		node["id"] == nil ||
		node["module"] == nil {
		return nil, errors.New("node type is missing")
	}

	n := new(Node)

	n.ID = fmt.Sprint(node["id"])
	n.Module = fmt.Sprint(node["module"])

	switch node["type"] {
	case "inlet":
		n.Type = NodeTypeInlet
	case "modifier":
		n.ListenFrom = fmt.Sprint(node["listen_from"])
		n.Type = NodeTypeModifier
		n.Task = fmt.Sprint(node["task"])
	case "outlet":
		n.ListenFrom = fmt.Sprint(node["listen_from"])
		n.Type = NodeTypeOutlet
	default:
		return nil, errors.New("node type is invalid: should be inlet, modifier or outlet")
	}

	if node["config"] == nil {
		return nil, errors.New("node config is missing")
	}

	n.Config = node["config"].(map[string]interface{})

	return n, nil
}
