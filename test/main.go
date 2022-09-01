package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.SQSEvent) error {

	for k, v := range request.Records {
		log.Printf("Record: %v, Body: %v", k, v.Body)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
