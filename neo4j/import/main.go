package main

import (
	"fmt"
	"flag"
	"time"
	"strconv"
	"os"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/csv"
)

func main() {

	var filePath *string = flag.String("file", "", "the filepath of the csv to import.")
	flag.Parse()
	fmt.Println("Filepath: " + *filePath)

	datasetMap := make(map[string]interface{})
	datasetMap["title"] = *filePath

	fmt.Printf("%+v\n", datasetMap)

	sendBatchUpdateToNeo("UNWIND {batch} as row MERGE (ds:Dataset {title: \"Test dataset\"})", []map[string]interface{}{ datasetMap })

	importCsvWithBatch(*filePath, sendBatchWithSingleRelation)
}

type DimensionOption struct {
	Dimension string
	Option    string
	Hierarchy string
}

type Observation struct {
	value   string
	options []DimensionOption
	index   int
}

type Parameters struct {
	Batch []map[string]interface{} `json:"batch"`
}

type Statement struct {
	Statement  string `json:"statement"`
	Parameters Parameters `json:"parameters"`
}

type Request struct {
	Statements []Statement `json:"statements"`
}

// Parse the CSV and create batches of rows
func importCsvWithBatch(filePath string, sendBatchFunc func(observations []*Observation)) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)

	var index = 0
	var batchSize = 2000
	var batchNumber = 1
	var isFinalBatch = false
	var totalRows int

	observations := make([]*Observation, batchSize)

	// Scan and discard header row (for now) - the data rows contain sufficient information about the structure
	if _, err := csvReader.Read(); err != nil {
		fmt.Printf("Encountered error immediately when processing header row")
		return
	}

	start := time.Now()

	for !isFinalBatch {

		fmt.Println("Processing batch number " + strconv.Itoa(batchNumber) + " index: " + strconv.Itoa(index))
		batchOffset := 0
		for batchIndex := 0; batchIndex < batchSize && !isFinalBatch; batchIndex++ {

			// each row in the batch
			line, err := csvReader.Read()
			if err != nil {
				fmt.Println("EOF reached, no more records to process", nil)
				isFinalBatch = true
				observations = observations[0:batchIndex] // the last batch is smaller than batch size, so resize the slice.
				fmt.Println(strconv.Itoa(batchIndex) + " messages in the final batch.")
				totalRows = ((batchNumber - 1) * batchSize) + batchIndex
				fmt.Println(strconv.Itoa(totalRows) + " messages in total.")
			} else {
				observation := createObservation(line, index)
				observations[batchOffset] = observation
				index++
				batchOffset++
			}
		}

		// send batch
		batchStart := time.Now()
		sendBatchFunc(observations)
		elapsed := time.Since(batchStart)
		fmt.Printf("Batch took %s\n", elapsed)

		//break
		batchNumber++
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time took %s\n", elapsed)
}

// take a CSV line and turn into observation instance.
func createObservation(tokens []string, index int) *Observation {

	dimensionOptions := make([]DimensionOption, 0)

	// parse each dimension that consists of 3 columns
	for i := 3; i < len(tokens); i += 3 {
		dimensionOption := &DimensionOption{
			Hierarchy: tokens[i],
			Dimension: tokens[i+1],
			Option:    tokens[i+2],
		}

		dimensionOptions = append(dimensionOptions, *dimensionOption)
	}

	return &Observation{
		value:   tokens[0],
		options: dimensionOptions,
		index:   index,
	}
}

// use a single relation for all options between dataset and observation.
func sendBatchWithSingleRelation(observations []*Observation) {

	batch := make([]map[string]interface{}, len(observations))
	dimensions := ""

	// build template query
	query := "UNWIND {batch} as row MATCH (ds:Dataset) WHERE ds.title = row.datasetName CREATE (ob:Observation {value: row.value})"
	for i, observation := range observations[0].options {
		optionId := "option" + strconv.Itoa(i)

		prefix := ", "
		if i == 0 {
			prefix = " "
		}
		dimensions = dimensions + prefix + "`" + observation.Dimension + "`:row." + optionId
	}
	query = query + ", (ob)<-[:hasObservation {" + dimensions + "}]-(ds)"

	// populate the batch array of parameters to feed into the query.
	for i, observation := range observations {

		params := map[string]interface{}{
			"value":       observation.value,
			"datasetName": "Test dataset",
		}

		for i, option := range observation.options {
			dimensionId := "dimension" + strconv.Itoa(i)
			optionId := "option" + strconv.Itoa(i)
			hierarchyId := "hierarchy" + strconv.Itoa(i)
			params[dimensionId] = option.Dimension
			params[optionId] = option.Option
			params[hierarchyId] = option.Hierarchy
		}

		batch[i] = params
	}

	//fmt.Println("query:", query)
	//fmt.Println("batch:", batch)

	sendBatchUpdateToNeo(query, batch)
}

// Send batch update to Neo. Uses HTTP API as I could not get the golang driver to work with batching.
func sendBatchUpdateToNeo(query string, batch []map[string]interface{}) {

	// build request object graph as required by Neo4j
	parameters := Parameters{
		Batch: batch,
	}
	statement := Statement{
		Statement:  query,
		Parameters: parameters,
	}
	statements := make([]Statement, 1)
	statements[0] = statement
	request := Request{
		Statements: statements,
	}
	url := "http://localhost:7474/db/data/transaction/commit"
	fmt.Println("URL:>", url)
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jsonBytes))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Accept", "application/json; charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
