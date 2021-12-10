package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ViewObj struct {
	IP        string `json:"ip"`
	Useragent string `json:"useragent"`
	URI       string `json:"uri"`
}

func main() {
	url := "http://localhost:8080/view/"

	for i := 0; i < 1000; i++ {
		log.Printf("Run %d", i)

		payloads := []ViewObj{
			{
				IP:        "127.0.0.1",
				Useragent: "myuseragent",
				URI:       "https://myaweseomwebsite.com/horseelevators",
			},
			{
				IP:        "127.255.0.1",
				Useragent: "myuseragent5",
				URI:       "https://myaweseomwebsite.com/horseelevators3",
			},
			{
				IP:        "127.252.0.1",
				Useragent: "myuseragent8",
				URI:       "https://myaweseomwebsite.com/horseelevators2358923958",
			},
		}

		for _, payload := range payloads {

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
	}

	for i := 0; i < 100; i++ {
		resp, err := http.Get("http://localhost:8080/requestcounts/")
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			panic(resp.StatusCode)
		}
	}

}
