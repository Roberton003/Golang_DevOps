package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func SendSMS(msg string, phone string) {
	endpoint := "http://rest.nexmo.com/sms/json"
	data := url.Values{}
	apiKey := os.Getenv("NEXMO_API_KEY")
	if api_key == "" {
		panic("NEXMO_API_KEY_SECRET")
	}
	apiSecret := os.Getenv("NEXMO_API_SECRET")
	if apiSecret == "" {
		panic("NEXMO_API_SECRET is not set")
	}
	data.Set("api_key", api_key)
	data.Set("api_secret", apiSecret)
	data.Set("to", phone)
	data.Set("from", "Sistema de alertas XGH")
	data.Set("text",message)
	cliete := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s", body)
}