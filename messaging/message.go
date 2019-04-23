package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/heaptracetechnology/microservice-cron/result"
	"github.com/robfig/cron"
)

type Subscribe struct {
	Data           map[string]string `json:"data"`
	Endpoint       string            `json:"target"`
	Id             string            `json:"id"`
	Interval       int               `json:"interval"`
	Delay_Interval int               `json:"delay_interval"`
}
type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//Cron service
func Cron(responseWriter http.ResponseWriter, request *http.Request) {
	client := cron.New()
	decoder := json.NewDecoder(request.Body)

	var listner Subscribe
	errr := decoder.Decode(&listner)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	interval := "@every 0h0m" + string(listner.Interval) + "s"
	fmt.Println(interval)
	if listner.Delay_Interval > 0 {
		delaytime := time.Second * time.Duration(listner.Delay_Interval)
		time.Sleep(delaytime)
	}
	client.AddFunc(interval, func() {

		t, err := cloudevents.NewHTTPTransport(
			cloudevents.WithTarget(listner.Endpoint),
			cloudevents.WithStructuredEncoding(),
		)
		if err != nil {
			log.Printf("failed to create transport, %v", err)
			return
		}

		c, err := cloudevents.NewClient(t,
			cloudevents.WithTimeNow(),
		)
		if err != nil {
			log.Printf("failed to create client, %v", err)
			return
		}
		fmt.Println("Data  ::::", listner.Data)
		contentType := "application/json"
		source, err := url.Parse("cron.event.subscribe")
		event := cloudevents.Event{
			Context: cloudevents.EventContextV01{
				Source:      cloudevents.URLRef{URL: *source},
				ContentType: &contentType,
				EventID:     listner.Id,
			}.AsV01(),
			Data: listner.Data,
		}

		fmt.Println(event)

		resp, err := c.Send(context.Background(), event)
		if err != nil {
			log.Printf("failed to send: %v", err)
		}
		if resp != nil {
			fmt.Printf("Response:\n%s\n", resp)
			fmt.Printf("Got Event Response Context: %+v\n", resp.Context)
			data := event
			if err := resp.DataAs(event); err != nil {
				fmt.Printf("Got Data Error: %s\n", err.Error())
			}
			fmt.Printf("Got Response Data: %+v\n", data)
		} else {
			log.Printf("event sent at %s", time.Now())
		}
	})

	client.Start()

	message := Message{"true", "Cron subscription started", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}
