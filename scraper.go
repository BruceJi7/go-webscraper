package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	var requestURL string = "https://www.devdungeon.com/archive"

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
	fmt.Println("Beginning request")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Request completed successfully")
	}
	defer response.Body.Close()

	outFile, err := os.Create("output.html")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File created")
	}
	defer outFile.Close()

	// Copy data from the response to the file
	n, err := io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to file:", n)

}
