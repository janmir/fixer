package main

import (
	"log"
	"net/http"

	"fixer/eu"

	"github.com/aws/aws-lambda-go/lambda"
)

//Converter interface
type Converter interface {
	Convert(from, to string) (float32, error)
}

var (
	client  *http.Client
	sources = []Converter{
		eu.Convert{},
	}
)

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
