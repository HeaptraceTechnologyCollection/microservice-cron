package cron

import (
	b "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/heaptracetechnology/microservice-cron/result"
	"github.com/robfig/cron"
)

type Subscribe struct {
	Data     Data   `json:"data"`
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
}

type Data struct {
	Interval     int64 `json:"interval"`
	InitialDelay int64 `json:"initial_delay"`
}

type RequestPayload struct {
	Data     map[string]string `json:"data"`
	Endpoint string            `json:"endpoint"`
	Id       string            `json:"id"`
}

//Cron service
func TriggerCron(responseWriter http.ResponseWriter, request *http.Request) {

	hc := http.Client{}
	client := cron.New()
	decoder := json.NewDecoder(request.Body)

	var listener Subscribe
	errr := decoder.Decode(&listener)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	stringInterval := strconv.FormatInt(listener.Data.Interval, 16)
	interval := "@every 0h0m" + stringInterval + "s"
	fmt.Println(interval)
	if listener.Data.InitialDelay > 0 {

		delaytime := time.Second * time.Duration(listener.Data.InitialDelay)
		time.Sleep(delaytime)
	}
	client.AddFunc(interval, func() {

		var request RequestPayload

		requestBody := new(b.Buffer)
		err := json.NewEncoder(requestBody).Encode(request)
		if err != nil {
			fmt.Println(" request err :", err)
		}

		req, errr := http.NewRequest("POST", listener.Endpoint, requestBody)
		if errr != nil {
			fmt.Println(" request err :", errr)
		}
		req.Header.Set("Content-Type", "application/json")
		_, errrs := hc.Do(req)
		if errrs != nil {
			fmt.Println("Client error", errrs)
		}
	})
	client.Start()

}
