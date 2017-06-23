package main

import (
	"encoding/csv"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	igniteURL = "http://127.0.0.1:8080"
	cacheName = ""
)

func main() {
	var csvFile string
	flag.StringVar(&igniteURL, "ignite-url", igniteURL, "apache ignite URL")
	flag.StringVar(&cacheName, "cache-name", cacheName, "cache name for file")
	flag.StringVar(&csvFile, "file-location", "", "A V3 csv file")
	flag.Parse()

	if csvFile == "" || cacheName == "" {
		log.Fatalf("Missing file-location or cache name or both")
	}

	file, err := os.Open(csvFile)
	if err != nil {
		log.Panic(err)
	}

	var Url *url.URL

	Url, err = url.Parse(igniteURL + "/ignite")
	if err != nil {
		log.Panic(err)
	}

	// Create cache
	cacheUrl := Url
	cacheParameters := url.Values{}
	cacheParameters.Add("cmd", "getorcreate")
	cacheParameters.Add("cacheName", cacheName)

	cacheUrl.RawQuery = cacheParameters.Encode()

	req, err := http.NewRequest("GET", cacheUrl.String(), nil)
	if err != nil {
		log.Printf("error forming request to create cache, error is : %s", err)
		log.Panic(err)
	}

	var httpClient http.Client
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("error creating cache, error is : %s", err)
		log.Panic(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("error creating cache, error is : %s", err)
		log.Panic(err)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	res.Body.Close()

	csvReader := csv.NewReader(file)
	skipFirstRow := false
	rowCount := 0
	var lengthOfTime time.Duration
	var csvHeaders []string
	startTime := time.Now()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		if skipFirstRow {
			doc := "{"
			for i := 0; i < len(record); i++ {
				if i+1 == len(record) {
					doc = doc + csvHeaders[i] + ":" + record[i] + "}"
				} else {
					doc = doc + csvHeaders[i] + ":" + record[i] + ","
				}
			}

			parameters := url.Values{}
			parameters.Add("cmd", "add")
			parameters.Add("key", strconv.Itoa(rowCount+1))
			parameters.Add("val", doc)
			parameters.Add("cacheName", cacheName)

			Url.RawQuery = parameters.Encode()

			startTimeOfRequest := time.Now()
			sendRequest(&httpClient, Url)
			endTimeOfRequest := time.Now()

			durationOfRequest := endTimeOfRequest.Sub(startTimeOfRequest)
			rowCount = rowCount + 1
			lengthOfTime = lengthOfTime + durationOfRequest
		} else {
			skipFirstRow = true
			csvHeaders = make([]string, len(record))
			copy(csvHeaders, record)
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("Completed, Rows %d, Time %d ns - %f s \nLength of all requests is : %vs", rowCount, duration.Nanoseconds(), duration.Seconds(), lengthOfTime.Seconds())
}

func sendRequest(httpClient *http.Client, Url *url.URL) {
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		log.Printf("error forming request to load doc, error is : %v", err)
		log.Panic(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("1, error writing document to cache, error is : %v", err)
		log.Panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("2, error writing document to cache, error is : %v", err)
		log.Panic(err)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
}
