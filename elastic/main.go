package main

import (
	"context"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/olivere/elastic.v5"
)

var (
	esDestURL   = "http://127.0.0.1:9200"
	esDestIndex = "v3data"
	esDestType  = "test1"
)

func main() {
	var docType, csvFile string
	flag.StringVar(&esDestURL, "es-dest-url", esDestURL, "elasticsearch destination URL")
	flag.StringVar(&esDestIndex, "es-dest-index", esDestIndex, "elasticsearch destination index")
	flag.StringVar(&esDestType, "es-dest-type", esDestType, "elasticsearch destination type")
	flag.StringVar(&csvFile, "file-location", "", "A V3 csv file")
	flag.Parse()

	file, err := os.Open(csvFile)
	if err != nil {
		log.Panic(err)
	}
	searchClient, _ := elastic.NewClient(
		elastic.SetURL(esDestURL),
		elastic.SetMaxRetries(5),
		elastic.SetSniff(false))
	bulk, _ := searchClient.BulkProcessor().
		BulkSize(30000).
		Workers(8).
		FlushInterval(time.Millisecond * 1000).
		Do(context.Background())
	bulk.Start(context.Background())
	log.Printf("Using file : %s into type : %s", csvFile, docType)
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
			docs := make(map[string]string)
			for i := 0; i < len(record); i++ {
				docs[csvHeaders[i]] = record[i]
			}
			request := elastic.NewBulkIndexRequest().
				Index(esDestIndex).
				Type(esDestType).
				Doc(docs)

			startTimeOfRequest := time.Now()
			bulk.Add(request)
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

	bulk.Flush()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("Completed, Rows %d, Time %d ns - %f s \nLength of all requests is : %vs", rowCount, duration.Nanoseconds(), duration.Seconds(), lengthOfTime.Seconds())
}
