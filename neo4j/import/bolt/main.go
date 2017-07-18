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
			"value":"6660",
			"dim1":"1",
			"dim2":"2",
			"dim3":"3",
			"dim4":"4",
		}

		rows = append(rows, row)
	}

	query := `UNWIND $rows AS row MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE d1.id = row.dim1
  AND d2.id = row.dim2
  AND d3.id = row.dim3
  AND d4.id = row.dim4
CREATE (o:observation { value:row.value}),
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
