package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	dbSession, err := mgo.Dial("localhost")
	var collectionName, csvFile string
	var ingestMultipleTimes int

	flag.StringVar(&csvFile, "f", "", "A V3 csv file")
	flag.StringVar(&collectionName, "c", "", "Collection name")
	flag.IntVar(&ingestMultipleTimes, "m", -1, "Ingest multiple times")
	flag.Parse()

	collection := dbSession.DB("v3data").C(collectionName)
	bulkLoader := collection.Bulk()
	file, err := os.Open(csvFile)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Using file : %s into mongodb collection : %s", csvFile, collectionName)
	for m := 0; m < ingestMultipleTimes; m++ {
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
				mRow := bson.M{}
				for i := 0; i < len(record); i++ {
					mRow[csvHeaders[i]] = record[i]
				}
				bulkLoader.Insert(mRow)
				rowCount = rowCount + 1
			} else {
				skipFirstRow = true
				csvHeaders = make([]string, len(record))
				copy(csvHeaders, record)
			}
		}
		bulkLoader.Run()
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Completed, Rows %d, Time %d ns - %f s", rowCount, duration.Nanoseconds(), duration.Seconds())
		//bulkLoader = collection.Bulk()
	}
}
