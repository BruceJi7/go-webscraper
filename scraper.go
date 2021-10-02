package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func processHREF(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {

		if strings.HasPrefix(href, "http") {
			fmt.Println("External link: ", href)
		}
	}
}

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

	// Create goquery document
	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal("Error loading HTTP response body ", err)
	} else {
		fmt.Println("Loaded response body into goquery")
	}

	//Find all links and print them if they exist
	document.Find("a").Each(processHREF)
}
