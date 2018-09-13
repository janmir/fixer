package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/aws/aws-lambda-go/events"
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

func getURL(url string) ([]byte, error) {
	data := []byte{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	return data, nil
}

func catch(err error) {
	if err != nil {
		_, file, no, _ := runtime.Caller(1)
		ss := strings.Split(file, "/")
		file = ss[len(ss)-1]
		report := fmt.Sprintf("%s:%d, %+v", file, no, err)

		//sent error by mail
		// reportMessage(report)
		log.Fatal(report)
	}
}
