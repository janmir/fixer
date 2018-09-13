package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

var (
	client *http.Client
)

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
