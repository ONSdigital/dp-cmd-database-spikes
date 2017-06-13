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

func main() {
	var docType, csvFile string

	flag.StringVar(&csvFile, "f", "", "A V3 csv file")
	flag.StringVar(&docType, "t", "", "type name")
	flag.Parse()

	file, err := os.Open(csvFile)
	if err != nil {
		log.Panic(err)
	}
	searchClient, _ := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetMaxRetries(5),
		elastic.SetSniff(false))
	bulk, _ := searchClient.BulkProcessor().
		BulkSize(15000).
		Workers(32).
		FlushInterval(time.Millisecond * 1000).
		Do(context.Background())
	bulk.Start(context.Background())
	log.Printf("Using file : %s into type : %s", csvFile, docType)
	csvReader := csv.NewReader(file)
	skipFirstRow := false
	rowCount := 0
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
				Index("v3data").
				Type(docType).
				Doc(docs)
			bulk.Add(request)
			rowCount = rowCount + 1
		} else {
			skipFirstRow = true
			csvHeaders = make([]string, len(record))
			copy(csvHeaders, record)
		}
	}
	bulk.Flush()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("Completed, Rows %d, Time %d ns - %f s", rowCount, duration.Nanoseconds(), duration.Seconds())
}
