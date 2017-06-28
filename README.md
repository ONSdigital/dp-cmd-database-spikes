Database spikes for customise my data
===================================

### Summary

The purpose of this project was to find the most appropriate database for storing and querying ONS datasets with dynamic queries.

##### Cassandra / google big table (distributed wide column store)

- Filter large tables quickly given a known key
- Tables built around queries
- Specific queries not known up front
- Not practical to model a table around every combination / permutation of filter

##### MongoDb (document database)

- Stores individual documents
- Created a single collection to hold the dataset
- document per row in the dataset
- Filtering was reasonably fast, but not efficient
- using a lot of disk swap space as it could not all be held in memory

##### Elastic (document database)

- Same considerations as MongoDb

##### Hadoop (Distributed map reduce)

- Can solve the problem as is scalable
- Not optimal for the problem
- More of a sledgehammer solution

##### Apache Ignite (in memory data grid)



##### Neo4j (graph database)

- Optimised for traversing relationship between data
- More performant with more filters applied


The input files used in the tests are zipped in the input-files directory.

### Test queries

Queries are in 'pseudo' sql due to variances in the databases. They are provided only to show the variations on filters

File size |Rows    |Dimensions  | File name
--|--|--|--
1030612770|10620815|3           |CensusEthnicity.csv
285529619 |1486273 |6           |ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv
82417638  |652159  |4           |RGVA01.csv
4554415   |39425   |4           |UKBAA01a.csv

#### ASHE07E dataset

##### select all data

```
SELECT * from observation
```

##### select a single point (filter on all dimensions)

```
SELECT * from observation
WHERE Geography="K02000001"
AND Year="2015"
AND Sex="CI_0006618"
AND `Working pattern`="CI_0006618"
AND Earnings="CI_0021537"
AND `Earnings statistics`="CI_0006603"
```

##### select a single dimension value

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
```

(123856 results)

##### select multiple dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
OR `Earnings statistics`="CI_0006604"
```

(247712 results)

##### select cross-dimension and get multiple values

```
SELECT * from observation
WHERE Earnings="CI_0021537"
AND   Sex="CI_0005444"
AND   `Earnings statistics`="CI_0021539"
```

(5161 results)
