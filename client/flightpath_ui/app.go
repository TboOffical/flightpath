package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ServerInfo struct {
	Name     string
	Addr     string
	Comments string
	ApiKey   string
}

type AppConfig struct {
	Servers        []ServerInfo
	OnboardingStep int //what step of the setup process is done
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

type updatePathParams struct {
	ID       int    `json:"id"`
	PathData string `json:"path_data"`
}

func (a *App) SavePath(addr string, apikey string, id int, data string) error {
	client := &http.Client{}

	dataBase64 := base64.StdEncoding.EncodeToString([]byte(data))

	params := updatePathParams{
		ID:       id,
		PathData: dataBase64,
	}

	paramsJson, _ := json.Marshal(params)

	fmt.Println(string(paramsJson))

	req, err := http.NewRequest("POST", addr+"/v1/paths/update", bytes.NewBuffer(paramsJson))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Token", apikey)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to update path")
	}

	return nil
}

type PathsResponse []struct {
	ID        int         `json:"ID"`
	CreatedAt string      `json:"CreatedAt"`
	UpdatedAt string      `json:"UpdatedAt"`
	DeletedAt interface{} `json:"DeletedAt"`
	Data      string      `json:"Data"`
}

func (a *App) GetPathsFromServer(addr string, apikey string) PathsResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", addr+"/v1/paths", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Token", apikey)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data PathsResponse
	err = json.Unmarshal(respData, &data)
	if err != nil {
		panic(err)
	}

	return data
}

func (a *App) GetInfoFromServer(addr string, apikey string) map[string]interface{} {
	client := &http.Client{}
	req, err := http.NewRequest("GET", addr+"/v1/server/server_info", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Token", apikey)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respData))

	var data map[string]interface{}
	err = json.Unmarshal(respData, &data)
	if err != nil {
		panic(err)
	}

	return data
}

type Docs []struct {
	Module string  `json:"Module"`
	Type   int64   `json:"Type"`
	Params []Param `json:"Params"`
	Tasks  []Task  `json:"Tasks"`
}

type Param struct {
	ParamName        string `json:"ParamName"`
	ParamType        string `json:"ParamType"`
	ParamDescription string `json:"ParamDescription"`
	ParamJSONField   string `json:"ParamJsonField"`
}

type Task struct {
	Name   string  `json:"Name"`
	Params []Param `json:"Params"`
}

func (a *App) GetDocsFromServer(addr string, apikey string) Docs {
	client := &http.Client{}
	req, err := http.NewRequest("GET", addr+"/v1/server/node_info", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Token", apikey)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respData))

	var data Docs
	err = json.Unmarshal(respData, &data)
	if err != nil {
		panic(err)
	}

	return data
}

func (a *App) AddServer(name string, addr string, apikey string) {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config AppConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	config.Servers = append(config.Servers, ServerInfo{
		Name:   name,
		ApiKey: apikey,
		Addr:   addr,
	})

	fileText, _ := json.MarshalIndent(config, "", "  ")

	err = os.WriteFile("./config.json", fileText, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *App) LoadConfig() *AppConfig {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		//create the file
		conf := &AppConfig{
			Servers:        nil,
			OnboardingStep: 0,
		}

		dataString, _ := json.Marshal(conf)
		err = os.WriteFile("./config.json", dataString, 0644)
		if err != nil {
			panic("unable to write file " + err.Error())
		}

		return conf
	}

	var config AppConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("Unable to parse config.json: " + err.Error())
		return nil
	}

	return &config
}

func (a *App) LoadAppConfig() {

}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
