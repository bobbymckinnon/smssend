package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	host := os.Getenv("SMSAPI_API")
	apiKey := os.Getenv("SMSAPI_API_KEY")

	client := &http.Client{}

	for _, message := range sqsEvent.Records {

		uri := fmt.Sprintf("/message?%s", message.Body)
		req, _ := http.NewRequest(http.MethodPost, host+uri, nil)
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("token", apiKey)

		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("\n Response Err : %s \n", err)
		}

		defer res.Body.Close()
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
