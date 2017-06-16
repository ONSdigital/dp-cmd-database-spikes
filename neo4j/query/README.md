### Benchmarks

Dataset        |Ingest    |First query  | Subsequent queries
--|--|--|--
CensusEthnicity|~8 minutes|         |
ASHE07E        |          |           |




#### ASHE07E.csv 

##### Ingest
Total time took 1m40.799147583s

##### select all data

Query: MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) RETURN r, ob
Time elapsed after query 1.673614ms
Time elapsed after streaming first result 2.816134ms
Time elapsed after streaming last result 19.70896688s
2017/06/16 11:02:14 Total number of rows in the result 1486273

Select count: 617.113391ms

##### Select a single point (filter on all dimensions)

Query: MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) WHERE r.Geography='K02000001' AND r.Year='2015' AND r.Sex='CI_0006618' AND r.`Working pattern`='CI_0006618' AND r.`Earnings`='CI_0021537' AND r.`Earnings statistics`='CI_0006603' RETURN r,ob
Time elapsed after query 3.204027752s
Time elapsed after streaming first result 6.377844306s
Time elapsed after streaming last result 6.377952815s
Total number of rows in the result 1

Select count: 6.205885321s

##### select a single dimension value

Query: MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) WHERE r.`Earnings statistics`='CI_0006603' RETURN r,ob
Time elapsed after query 1.147314ms
Time elapsed after streaming first result 2.557089ms
Time elapsed after streaming last result 2.155416875s
Total number of rows in the result 123856

Select count: 1.328864271s

##### select multiple dimension values

Query: MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) WHERE r.Geography='K02000001' AND r.Year='2015' AND r.Sex='CI_0006618' AND r.`Working pattern`='CI_0006618' AND r.`Earnings`='CI_0021537' AND (r.`Earnings statistics`='CI_0006603' OR r.`Earnings statistics`='CI_0006604') RETURN r,ob
Time elapsed after query 2.942279368s
Time elapsed after streaming first result 6.893523156s
Time elapsed after streaming last result 6.893658169s
Total number of rows in the result 2

Select count: 6.999783961s

##### select cross-dimension and get multiple values

Query: MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation) WHERE r.Sex='CI_0005444' AND r.`Earnings`='CI_0021537' AND r.`Earnings statistics`='CI_0021539' RETURN r,ob
Time elapsed after query 1.338209ms
Time elapsed after streaming first result 16.933874ms
Time elapsed after streaming last result 1.373889806s
Total number of rows in the result 5161

Select count: 1.385792354s

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
