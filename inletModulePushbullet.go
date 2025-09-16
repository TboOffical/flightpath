package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Push struct {
	Pushes []struct {
		Active                  bool     `json:"active"`
		Iden                    string   `json:"iden"`
		Created                 float64  `json:"created"`
		Modified                float64  `json:"modified"`
		Type                    string   `json:"type"`
		Dismissed               bool     `json:"dismissed"`
		GUID                    string   `json:"guid"`
		Direction               string   `json:"direction"`
		SenderIden              string   `json:"sender_iden"`
		SenderEmail             string   `json:"sender_email"`
		SenderEmailNormalized   string   `json:"sender_email_normalized"`
		SenderName              string   `json:"sender_name"`
		ReceiverIden            string   `json:"receiver_iden"`
		ReceiverEmail           string   `json:"receiver_email"`
		ReceiverEmailNormalized string   `json:"receiver_email_normalized"`
		SourceDeviceIden        string   `json:"source_device_iden"`
		AwakeAppGuids           []string `json:"awake_app_guids"`
		Body                    string   `json:"body"`
	} `json:"pushes"`
}

func pushbulletPublisher(opts map[string]interface{}, id string) {
	apiKey := fmt.Sprint(opts["api_key"])

	u := url.URL{Scheme: "wss", Host: "stream.pushbullet.com", Path: "/websocket/" + apiKey}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		e("Failed to connect to pushbullet websocket")
		return
	}

	e("Subscribed to pushbullet event stream")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		if fmt.Sprintf("%s", message) == "{\"type\":\"tickle\",\"subtype\":\"push\"}" {
			e("Push detected, downloading...")

			t := time.Now()
			ue := "https://api.pushbullet.com/v2/pushes?modified_after=" + fmt.Sprint(t.Unix())

			log.Println("Requesting ", ue)

			client := http.Client{Timeout: time.Second * 10}
			req, err := http.NewRequest("GET", ue, nil)
			if err != nil {
				log.Println(err)
				return
			}

			req.Header.Add("Access-Token", apiKey)

			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}

			var push Push
			err = json.Unmarshal(body, &push)
			if err != nil {
				log.Println(err)
			}

			pushDataJson, _ := json.Marshal(push.Pushes[0])

			incoming <- IncomingMessage{
				From:    id,
				Message: string(pushDataJson),
			}
		}

	}

}

func newPushbulletInlet(id string) *InletModule {
	return &InletModule{
		ID:        id,
		Name:      "pushbullet",
		Publisher: pushbulletPublisher,
	}
}
