package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"fixer/eu"
	. "fixer/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Converter interface
type Converter interface {
	Convert(from, to string) (float32, error)
	Trend(from, to string) (string, error)
	Rate(from, to string) (float32, error)
}

var (
	sources = []Converter{
		eu.Fx{},
	}
)

//Response ...
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Okay so your other function also executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}

func init() {
	client = http.DefaultClient
}

func main() {
	switch {
	case _dev:
		log.Println("This is dev!")

		convert()
	default:
		lambda.Start(Handler)
	}
}

func convert() error {
	return nil
}
