package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	messageHandler struct{}
	Message        struct {
		Name      string `json:"name"`
		Content   string `json:"content"`
		Timestamp string `json:"time_stamp"`
	}
)

func main() {
	//instantiate config
	config := nsq.NewConfig()
	//tweak several common setup the config
	//max number of times this consumer will attempt to process a message before giving up
	config.MaxAttempts = 10
	//max number of messages to allow in flight
	config.MaxInFlight = 5
	//maximum duration when Requeueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0
	//init topic name and channel
	topic := "sample_topic"
	channel := "channel_example"
	//creating consumer
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	//Set the handler for message received by this consumer.
	consumer.AddHandler(&messageHandler{})
	//use nsqlookupd to find nsqd instances
	consumer.ConnectToNSQD("127.0.0.1:4150")
	//wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	//gracefully stop the consumer
	defer consumer.Stop()
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	//process the message
	//var request Message
	//if err := json.Unmarshal(m.Body, &request); err != nil {
	//	log.Println(err)
	//	return err
	//}
	////print the message
	//log.Println("Name :", request.Name)
	//log.Println(request)
	////log body
	//fmt.Println("bod")
	log.Println(string(m.Body))
	return nil
}
