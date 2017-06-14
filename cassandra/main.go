package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	//stmt := "INSERT INTO ukbaa01a (Observation,Data_Marking,Observation_Type_Value,Dimension_Hierarchy_1,Dimension_Name_1,Dimension_Value_1,Dimension_Hierarchy_2,Dimension_Name_2,Dimension_Value_2,Dimension_Hierarchy_3,Dimension_Name_3,Dimension_Value_3,Dimension_Hierarchy_4,Dimension_Name_4,Dimension_Value_4) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	log.Printf("Ingesting V3 file")
	var csvFile, stmt string
	var batchSize int

	flag.StringVar(&csvFile, "f", "", "A V3 csv file")
	flag.StringVar(&stmt, "s", "", "CQL insert statement")
	flag.IntVar(&batchSize, "b", -1, "Batch size to upload")
	flag.Parse()

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "v3data"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Panic(err)
	}
	batcher := session.NewBatch(gocql.LoggedBatch)

	file, err := os.Open(csvFile)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Using file : %s and CQL : %s", csvFile, stmt)
	csvReader := csv.NewReader(file)
	var skipFirstRow = false
	rowCount := 0
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

			batcher.Query(stmt, InterfaceSlice(record)...)
			rowCount = rowCount + 1
		}
		skipFirstRow = true
		if len(batcher.Entries) > batchSize {
			err1 := session.ExecuteBatch(batcher)
			if err1 != nil {
				log.Panic(err1)
			}
			batcher = session.NewBatch(gocql.LoggedBatch)
		}
	}
	err1 := session.ExecuteBatch(batcher)
	if err1 != nil {
		log.Panic(err1)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("Completed, Rows %d, Time %d ns - %f s", rowCount, duration.Nanoseconds(), duration.Seconds())
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
