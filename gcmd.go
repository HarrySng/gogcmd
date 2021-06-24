package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var staticURL = "https://gcmd.earthdata.nasa.gov/static/kms/"

func main() {
	makeRequest(staticURL)
}

func makeRequest(url string) {

	req, err := http.NewRequest("GET", url, nil)
	handleError(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	handleError(err)

	if resp.StatusCode != 200 {
		err := errors.New("Something went wrong")
		// Add some troubleshooting info here
		fmt.Println(resp.StatusCode)
		handleError(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		handleError(err)
		fmt.Println(body)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
