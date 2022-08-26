package nsqApp

import (
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitConsumer(topic string, channel string) *nsq.Consumer {
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

	//creating consumer
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	return consumer
}

func NSQConsumerConnection(ns *nsq.Consumer) {
	ns.ConnectToNSQLookupd(viper.GetString("NSQ_LOOKUPD"))

	//wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	//gracefully stop the consumer
	ns.Stop()
}

func InitProducer() *nsq.Producer {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(viper.GetString("NSQ_D"), config)
	return w
}
