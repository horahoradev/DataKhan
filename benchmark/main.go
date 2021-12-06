package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ViewObj struct {
	IP        string `json:"ip"`
	Useragent string `json:"useragent"`
	URI       string `json:"uri"`
}

func main() {
	url := "http://localhost:8080/view/"

	payload := ViewObj{
		IP:        "127.0.0.1",
		Useragent: "myuseragent",
		URI:       "https://myaweseomwebsite.com/horseelevators",
	}
	json, err := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
