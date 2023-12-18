package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendPostRequest(pubkey string) {
	fmt.Println("Sending POST request to /gapi/verified_users")
	payload := []map[string]string{
		{
			"pubkey": pubkey,
			"kind":   "blue",
			"desc":   "string",
		},
	}
	payloadBytes, _ := json.Marshal(payload)
	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("POST", "https://devapi.freefrom.space/gapi/verified_users", body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer dev")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
