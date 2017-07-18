package main

import (
	"time"
	"fmt"
	"sync"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(10)

	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)
	go addBatch(&wg)

	wg.Wait()
}

func addBatch(wg *sync.WaitGroup) {

	defer wg.Done()

	driver := bolt.NewDriver()
	conn, _ := driver.OpenNeo("bolt://localhost:7687")
	defer conn.Close()

	var batchSize = 1000

	rows := make([]interface{}, 0)

	for batchIndex := 0; batchIndex < batchSize; batchIndex++ {

		row := map[string]interface{}{
			"v":"6660",
			"d1":0,
			"d2":1,
			"d3":2,
			"d4":3,
		}

		rows = append(rows, row)
	}

	query := `UNWIND $rows AS row MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE id(d1) = row.d1
  AND id(d2) = row.d2
  AND id(d3) = row.d3
  AND id(d4) = row.d4
CREATE (o:observation { value:row.v}),
       (o)-[:isValueOf]->(d1),
       (o)-[:isValueOf]->(d2),
       (o)-[:isValueOf]->(d3),
       (o)-[:isValueOf]->(d4)`

	//fmt.Printf("%+v", rows)
	//fmt.Println("query: " + query)

	start := time.Now()

	result, err := conn.ExecNeo(query, map[string]interface{}{"rows": rows})
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Error(err)
		return
	}

	numResult, _ := result.RowsAffected()
	fmt.Printf("CREATED ROWS: %d\n", numResult)

	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)

}
