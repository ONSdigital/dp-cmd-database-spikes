### Benchmarks

Dataset        |Ingest    |First query  | Subsequent queries
--|--|--|--
CensusEthnicity|~8 minutes|         |
ASHE07E        |          |           |




#### ASHE07E.csv 

##### Ingest
Total time took 1m40.799147583s


#### CensusEthnicity.csv 

##### Ingest 
Total time took 7m37.645472317s


##### Query all results
```
MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) RETURN ob,r
```

First run
- Time elapsed after query 27.488651ms
- Time elapsed after streaming first result 36.217935ms
- Time elapsed after streaming last result 2m35.384395931s
- Total number of rows in the result 10620815

Second run
- Time elapsed after query 3.83537ms
- Time elapsed after streaming first result 11.402184ms
- Time elapsed after streaming last result 2m16.134099315s
- Total number of rows in the result 10620815


##### Query filtering by geography only
```
MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) 
    WHERE r.Geography='K04000001' 
    RETURN ob,r
```
First run
- Time elapsed after query 82.056461ms
- Time elapsed after streaming first result 18.770262822s
- Time elapsed after streaming last result 18.774979055s
- Total number of rows in the result 252

Second run
- Time elapsed after query 3.105634ms
- Time elapsed after streaming first result 7.528075834s
- Time elapsed after streaming last result 7.531624573s
- Total number of rows in the result 252

##### Query all dimensions
```
MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) 
    WHERE r.Geography='K04000001' 
    AND r.Year='2011' 
    AND r.Ethnicity='White: Irish' 
    RETURN r,ds,ob
```

First run
- Time elapsed after query 41.90339ms
- Time elapsed after streaming first result 28.83618943s
- Time elapsed after streaming last result 28.837618711s
- Total number of rows in the result 1

Second run
- Time elapsed after query 49.91393ms
- Time elapsed after streaming first result ~8s
- Time elapsed after streaming last result ~8s
