package main

import (
	"time"
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"bytes"
	"sync"
)

// http://neo4j.com/docs/rest-docs/current/#rest-api-batch-ops


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
	start := time.Now()

	var batchSize = 1000

	updates := make([]*update, 0)
	commandIndex := 0

	for batchIndex := 0; batchIndex < batchSize; batchIndex++ {

		observation := createObservationCommand(batchIndex)
		updates = append(updates, observation)
		commandIndex++
	}


	for batchIndex := 0; batchIndex < batchSize; batchIndex++ {

		addLabel := addLabelCommand(commandIndex, batchIndex)
		updates = append(updates, addLabel)
		commandIndex++

		dim1 := addRelationCommand(commandIndex, batchIndex, 0)
		updates = append(updates, dim1)
		commandIndex++

		dim2 := addRelationCommand(commandIndex, batchIndex, 1)
		updates = append(updates, dim2)
		commandIndex++

		dim3 := addRelationCommand(commandIndex, batchIndex, 2)
		updates = append(updates, dim3)
		commandIndex++

		dim4 := addRelationCommand(commandIndex, batchIndex, 3)
		updates = append(updates, dim4)
		commandIndex++
	}

	b, err := json.Marshal(updates)
	if err != nil {
		fmt.Println(err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("json took %s\n", elapsed)

	//fmt.Println(string(b))

	url := "http://localhost:7474/db/data/batch"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
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
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	elapsed = time.Since(start)
	fmt.Printf("took %s\n", elapsed)

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
