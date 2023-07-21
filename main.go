package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	file := flag.String("f", "mock.json", "filepath to the file to be published into the topic")
	topic := flag.String("t", "mock", "topic name to be used to publish the message")
	flag.Parse()

	content, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	w := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    *topic,
		Balancer: &kafka.LeastBytes{},
	}

	if err := w.WriteMessages(context.Background(), kafka.Message{Value: content}); err != nil {
		panic(err)
	}

	fmt.Println("Published successfully")
}
