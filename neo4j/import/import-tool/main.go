package main

import (
	"time"
	"fmt"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"os"
)

// http://neo4j.com/docs/rest-docs/current/#rest-api-batch-ops

func main() {
	start := time.Now()

	driver := bolt.NewDriver()
	conn, _ := driver.OpenNeo("bolt://localhost:7687")
	defer conn.Close()

	var batchSize = 10000

	csvContent := "value,dim1,dim2,dim3,dim4\n"

	for batchIndex := 0; batchIndex < batchSize; batchIndex++ {
		csvContent += "677,1,2,3,4\n"
	}

	f, err := os.Create("/usr/local/Cellar/neo4j/3.2.0/libexec/import/obs.csv")
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	f.WriteString(csvContent)
	f.Sync()

	//	query := `LOAD CSV WITH HEADERS FROM "file:///obs.csv" AS line WITH line
	//MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
	//  WHERE d1.id = "1"
	//  AND d2.id = "2"
	//  AND d3.id = "3"
	//  AND d4.id = "4"
	//CREATE (o:observation { value:line.value}),
	//       (o)-[:isValueOf]->(d1),
	//       (o)-[:isValueOf]->(d2),
	//       (o)-[:isValueOf]->(d3),
	//       (o)-[:isValueOf]->(d4)`
	//	conn.ExecNeo(query, map[string]interface{}{"value": 666})

	elapsed := time.Since(start)
	fmt.Printf("json took %s\n", elapsed)

	//fmt.Println(string(b))
}
