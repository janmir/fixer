package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	sources = []string{
		"http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml",
	}
	client *http.Client
)

//Response ...
type Response events.APIGatewayProxyResponse

//CubeParent data
type CubeParent struct {
	Cube CubeTime `xml:"Cube"`
}

//CubeTime data
type CubeTime struct {
	Time string `xml:"time,attr"`
	Cube []Cube `xml:"Cube"`
}

//Cube data
type Cube struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

//EuroCenterBankXML Europen Bank XMl data structure
type EuroCenterBankXML struct {
	Subject string `xml:"subject"`
	Sender  struct {
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube CubeParent `xml:"Cube"`
}

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

		fetchSources()
	default:
		lambda.Start(Handler)
	}
}

func fetchSources() error {
	for _, v := range sources {
		log.Println(v)

		//fetch the data
		data := getURL(v)

		exml := EuroCenterBankXML{}
		err := xml.Unmarshal(data, &exml)
		catch(err)
		log.Printf("%+v", exml)
	}
	return nil
}

func getURL(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	catch(err)

	res, err := client.Do(req)
	catch(err)

	data, err := ioutil.ReadAll(res.Body)
	catch(err)

	return data
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
