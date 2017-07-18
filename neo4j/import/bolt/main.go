package main

import (
	"time"
	"fmt"
	"strconv"
	"sync"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// http://neo4j.com/docs/rest-docs/current/#rest-api-batch-ops

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go addBatch(&wg)
	//go addBatch(&wg)
	//go addBatch(&wg)
	//go addBatch(&wg)

	wg.Wait()
}

func addBatch(wg *sync.WaitGroup) {
	start := time.Now()

	driver := bolt.NewDriver()
	conn, _ := driver.OpenNeo("bolt://localhost:7687")
	defer conn.Close()

	var batchSize = 10000

	query := `MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE d1.id = "1"
  AND d2.id = "2"
  AND d3.id = "3"
  AND d4.id = "4"
CREATE (o:observation { value:{value}}),
       (o)-[:isValueOf]->(d1),
       (o)-[:isValueOf]->(d2),
       (o)-[:isValueOf]->(d3),
       (o)-[:isValueOf]->(d4)`

	//queries := make([]string, 0)
	//data := make([]map[string]interface{}, 0)

	for batchIndex := 0; batchIndex < batchSize; batchIndex++ {

		conn.ExecNeo(query, map[string]interface{}{"value": 666})
		//queries = append(queries, query)
		//data = append(data, map[string]interface{}{ "value":"345" })

	}

	//conn.ExecPipeline(queries, data...)

	//for _, result := range results {
	//	numResult, _ := result.RowsAffected()
	//	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 2 (per each iteration)
	//}

	elapsed := time.Since(start)
	fmt.Printf("json took %s\n", elapsed)

	//fmt.Println(string(b))

	wg.Done()
}

func createObservationCommand(updateIndex int) *update {
	return &update{
		Method: "POST",
		To:     "/node",
		ID:     updateIndex,
		Body: map[string]string{
			"value": "12345",
		},
	}
}

// Add a label for the given observation index
func addLabelCommand(updateIndex int, observationIndex int) *update {
	return &update{
		Method: "POST",
		To:     "{" + strconv.Itoa(observationIndex) + "}/labels",
		ID:     updateIndex,
		Body:   "observation",
	}
}

func addRelationCommand(updateIndex int, observationIndex int, dimensionID int) *update {
	return &update{
		Method: "POST",
		To:     "{" + strconv.Itoa(observationIndex) + "}/relationships",
		ID:     updateIndex,
		Body: map[string]string{
			"to":   "/node/" + strconv.Itoa(dimensionID),
			"type": "isValueOf",
		},
	}
}

type update struct {
	Method string `json:"method"`
	To     string `json:"to"`
	ID     int `json:"id"`
	Body   interface{} `json:"body"`
}
