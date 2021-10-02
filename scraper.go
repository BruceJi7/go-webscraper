package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	var requestURL string = "https://www.devdungeon.com/"

	//Setup client, use timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Initialise HTTP request before sending it
	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "TUTORIAL SCRAPER WHEEEE")

	// Make HTTP GET request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)

}
