package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// build application infrastructure
func buildPaths() error {
	log.Println("Building Paths")
	e("Path Loading Started")

	//load the paths from the db
	var paths []Path
	tx := db.Find(&paths).Limit(200)
	if tx.Error != nil {
		return tx.Error
	}

	numPaths := len(paths)

	for i, path := range paths {
		e(fmt.Sprint("Building Path ", i+1, "/", numPaths))
		err = loadPathFromBytesAndCreate(path.Data)
		if err != nil {
			log.Println("Error building Path:", err)
			e(fmt.Sprint("Failed to build path with index", i))
			return err
		}
	}

	return nil
}

// create the correct pathway in the system from the .json path config data
func loadPathFromBytesAndCreate(data string) error {
	var dataObject map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataObject)

	if err != nil {
		return err
	}

	pathName := fmt.Sprint(dataObject["title"])
	e(fmt.Sprint("Building ", pathName, "..."))

	if dataObject["nodes"] == nil {
		return errors.New("invalid path syntax: no nodes detected in map")
	}

	nodes := dataObject["nodes"].([]interface{})

	for _, node := range nodes {
		n, err := newNodeFromInterface(node)
		if err != nil {
			return err
		}

		//make sure the node has an ID
		if len(n.ID) == 0 {
			return errors.New("invalid path syntax: no id detected in node config")
		}

		//if the node is an inlet, create an inlet

		if n.Type == NodeTypeInlet {
			switch n.Module {
			case "pushbullet":
				i := newPushbulletInlet(n.ID)
				i.Configure(n.Config)
				err = i.Start()
				if err != nil {
					return err
				}
			case "time_trigger":
				i := newTimeTriggerInlet(n.ID)
				i.Configure(n.Config)
				err = i.Start()
				if err != nil {
					return err
				}
			default:
				return errors.New("module not found, if you just added this module make sure to add it to the this list")
			}
		}

		//if it's a modifier, create the modifier

		if n.Type == NodeTypeModifier {
			switch n.Module {
			case "text":
				i := newTextModifier(n.ID, n.Task, n.ListenFrom)
				i.Configure(n.Config)
				err = i.Start()
				if err != nil {
					return err
				}
			default:
				return errors.New("module not found")
			}
		}

		//finally, if it's an outlet, create that
		/*todo: late on make sure all of this is not manual, since the starting procedure is the same for all of the types
		  todo: it can just loop through an array of modules and start the ones that are needed
		*/

		if n.Type == NodeTypeOutlet {
			switch n.Module {
			case "email":
				i := newEmailOutlet(n.ID, n.ListenFrom)
				i.Configure(n.Config)
				err = i.Start()
				if err != nil {
					return err
				}
			default:
				return errors.New("module not found")
			}
		}

	}

	return nil
}
