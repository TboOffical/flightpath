package main

import "errors"

func textModify(content string, config map[string]interface{}, task string) string {
	switch task {
	case "prefix":
		prefix := config["prefix"].(string)
		return prefix + content
	default:
		return ""
	}
}

func textVerify(config map[string]interface{}, task string) error {
	switch task {
	case "prefix":
		if config["prefix"] == nil {
			return errors.New("no prefix")
		}
	default:
		return errors.New("unsupported text modifier task")
	}

	return nil
}

func newTextModifier(id string, task string, listenFrom []string) *ModifierModule {
	return &ModifierModule{
		Name:       "text",
		ID:         id,
		Task:       task,
		ListenFrom: listenFrom,
		Modify:     textModify,
		Verify:     textVerify,
		configured: false,
	}
}
