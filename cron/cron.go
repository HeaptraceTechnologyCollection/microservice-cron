package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go"
	"github.com/oms-services/cron/result"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Subscribe struct {
	Data      Data   `json:"data"`
	Endpoint  string `json:"endpoint"`
	Id        string `json:"id"`
	IsTesting bool   `json:"istesting"`
}

type Data struct {
	Interval     int64 `json:"interval"`
	InitialDelay int64 `json:"initialDelay"`
}

type RequestPayload struct {
	Data     map[string]string `json:"data"`
	Endpoint string            `json:"endpoint"`
	Id       string            `json:"id"`
}
type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//Cron service
func TriggerCron(responseWriter http.ResponseWriter, request *http.Request) {

	client := cron.New()
	decoder := json.NewDecoder(request.Body)

	var listener Subscribe
	errr := decoder.Decode(&listener)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	if listener.Data.InitialDelay < 0 || listener.Data.Interval < 0 {
		message := Message{"false", "Invalid time", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	stringInterval := strconv.Itoa(int(listener.Data.Interval))
	interval := "@every 0h0m" + stringInterval + "s"
	client.AddFunc(interval, func() {
		if listener.Data.InitialDelay > 0 {
			delaytime := time.Second * time.Duration(listener.Data.InitialDelay)
			time.Sleep(delaytime)
		}

		contentType := "application/json"
		t, err := cloudevents.NewHTTPTransport(
			cloudevents.WithTarget(listener.Endpoint),
			cloudevents.WithStructuredEncoding(),
		)

		if err != nil {
			log.Printf("failed to create transport, %v", err)
			return
		}

		cloudClient, err := cloudevents.NewClient(t,
			cloudevents.WithTimeNow(),
		)

		source, err := url.Parse(listener.Endpoint)
		event := cloudevents.Event{
			Context: cloudevents.EventContextV01{
				EventID:     listener.Id,
				EventType:   "triggers",
				Source:      cloudevents.URLRef{URL: *source},
				ContentType: &contentType,
			}.AsV01(),
			Data: "",
		}
		resp, evt, err := cloudClient.Send(context.Background(), event)
		if err != nil {
			log.Printf("failed to send: %v (%v)", err, evt)
			fmt.Println(resp)
		}

	})

	client.Start()
	message := Message{"true", "Cron event triggered", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

	if listener.IsTesting == true {
		time.Sleep(5 * time.Second)
		client.Stop()
	}
}
