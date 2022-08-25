package main

import (
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)
	//if err != nil {
	//	panic(err)
	//}
	for i := 0; i < 10; i++ {
		err := w.Publish("sample_topic", []byte("test"+strconv.Itoa(i)))
		if err != nil {
			log.Panic("Could not connect")
		}
	}
	w.Stop()
}
