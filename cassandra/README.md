## Apache cassandra Database spike

### Setup
* ```brew install cassandra``` (version 3.10_1)
* ```go get github.com/gocql/gocql```

### Create keyspace (Database)
```
CREATE KEYSPACE V3Data WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};
```
Note replication is 1!

### Creating a table
```
DROP TABLE ukbaa01a;
CREATE TABLE ukbaa01a (
  Observation text,
  Data_Marking text,
  Observation_Type_Value text,
  Dimension_Hierarchy_1 text,
  Dimension_Name_1 text,
  Dimension_Value_1 text,
  Dimension_Hierarchy_2 text,
  Dimension_Name_2 text,
  Dimension_Value_2 text,
  Dimension_Hierarchy_3 text,
  Dimension_Name_3 text,
  Dimension_Value_3 text,
  Dimension_Hierarchy_4 text,
  Dimension_Name_4 text,
  Dimension_Value_4 text,
  PRIMARY KEY (Dimension_Value_1, Dimension_Value_2)
  );
```
```
DROP TABLE ASHE07E;
CREATE table ASHE07E (
Observation text,
Data_Marking text,
Observation_Type_Value text,
Dimension_Hierarchy_1 text,
Dimension_Name_1 text,
Dimension_Value_1 text,
Dimension_Hierarchy_2 text,
Dimension_Name_2 text,
Dimension_Value_2 text,
Dimension_Hierarchy_3 text,
Dimension_Name_3 text,
Dimension_Value_3 text,
Dimension_Hierarchy_4 text,
Dimension_Name_4 text,
Dimension_Value_4 text,
Dimension_Hierarchy_5 text,
Dimension_Name_5 text,
Dimension_Value_5 text,
Dimension_Hierarchy_6 text,
Dimension_Name_6 text,
Dimension_Value_6 text,
PRIMARY KEY (Dimension_Value_1, Dimension_Value_2,));
```

### Query Commands


### Performance tests for Cassandra
Ingesting times    
Using a batch size of 100  
* ukbaa01a = 0.67 seconds (5 MB, rows 39424)
* ASHE07E = 48.731878 second (285 MB, rows 1486272)

Querying times  
For both Data sets queries where around 3-4 milliseconds 

Limits  
See http://www.datastax.com/dev/blog/basic-rules-of-cassandra-data-modeling and https://www.datastax.com/dev/blog/the-most-important-thing-to-know-in-cassandra-data-modeling-the-primary-key

For the v3data we will need to create multiple tables to be able to filter on different parts of the 
dataset. This will lead to creating a lot of tables for datasets that have many dimensions. You can force cassandra
to perform a filter scan by this throws out Performance.

### Add CSV data manually 
```
COPY ukbaa01a(Observation,Data_Marking,Observation_Type_Value,Dimension_Hierarchy_1,Dimension_Name_1,Dimension_Value_1,Dimension_Hierarchy_2,Dimension_Name_2,Dimension_Value_2,Dimension_Hierarchy_3,Dimension_Name_3,Dimension_Value_3,Dimension_Hierarchy_4,Dimension_Name_4,Dimension_Value_4) FROM 'UKBAA01a.csv' WITH DELIMITER=',' AND HEADER=true;
```
#### Useful Links
http://docs.datastax.com/en/cql/3.3/index.html