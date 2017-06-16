package main

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"time"
	"log"
	"flag"
)

func main() {

	var query string
	flag.StringVar(&query, "q", "", "query")
	flag.Parse()

	log.Printf("Query: %s", query)

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://localhost:7687")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	startTime := time.Now()

	stmt, err := conn.PrepareNeo(query)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.QueryNeo(nil)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(startTime)
	log.Printf("Time elapsed after query %s\n", elapsed)

	data, _, err := rows.NextNeo()
	if err != nil {
		panic(err)
	}

	rowCount := 1

	elapsed = time.Since(startTime)
	log.Printf("Time elapsed after streaming first result %s\n", elapsed)
	log.Printf("First row data: %#v\n", data)

	for err == nil {
		data, _, err = rows.NextNeo()
		if err == nil {
			rowCount = rowCount + 1
			//log.Printf("Row %d data: %#v\n", rowCount, data)
		}
	}

	elapsed = time.Since(startTime)
	log.Printf("Time elapsed after streaming last result %s\n", elapsed)
	log.Printf("Total number of rows in the result %d\n", rowCount)

	stmt.Close()
}
