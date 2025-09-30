package main

import "github.com/gofiber/fiber/v2"

var AppDocs []*DocsEntry

type ParamDoc struct {
	ParamName        string
	ParamType        string
	ParamDescription string
	ParamJsonField   string
}

type Task struct {
	Name   string
	Params []ParamDoc
}

type DocsEntry struct {
	Module string
	Type   int
	Params []ParamDoc
	Tasks  []Task
}

func newDocsEntry(module string, moduleType int) *DocsEntry {
	return &DocsEntry{
		Module: module,
		Type:   moduleType,
	}
}

func (entry *DocsEntry) AddParam(name string, t string, description string, json string) {
	entry.Params = append(entry.Params, ParamDoc{name, t, description, json})
}

func (entry *DocsEntry) AddTask(t Task) {
	entry.Tasks = append(entry.Tasks, t)
}

func getDocsHandler(c *fiber.Ctx) error {
	//Just send back the node docs
	return c.JSON(AppDocs)
}
