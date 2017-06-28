### Benchmarks

#### CensusEthnicity.csv 

##### Ingest 
Total time took ~16 minutes per GB

##### Query multiple ethnicity values

MATCH (d3:CensusEthnicity_Dimension3) \
WHERE d3.value = 'Other ethnic group: Thai' OR d3.value = 'White: Colombian' OR d3.value = 'Other ethnic group: Punjabi' OR d3.value = 'White: Albanian' \
WITH d3 MATCH (d1:CensusEthnicity_Dimension1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), \
(d3:CensusEthnicity_Dimension3)<-[:has]-(o) \
return o

1gb: 1.3s
15gb: ~13.5m
20gb: ~20m

##### Query filtering by ethnicity only

```
MATCH (d3:CensusEthnicity_Dimension3 {value:'White: Afghan'}) \
WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) \
return o
```

1gb:  0.3 seconds
15gb: ~3-4m (634710 results)
20gb: ~5.5m (846280 results)

##### Query filtering by geography only
```
MATCH (d2:CensusEthnicity_Dimension2 {value:'E12000009'}) 
WITH d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), 
    (o)-[:has]->(d3:CensusEthnicity_Dimension3) 
return o, d1, d2, d3
```

15gb: ~100ms (3765 results)
20gb: ~100ms (5020 results)

##### Match on year and geography

```
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), \
    (d2:CensusEthnicity_Dimension2 {value:'K04000001'}) \
WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2) \
return o
```

```
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), \
    (d2:CensusEthnicity_Dimension2 {value:'E92000001'}) \
WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2) \
return o
```

1gb ~6ms
15gb: ~100ms (3765 results)
20gb: ~150ms (5020 results)

##### Query all dimensions
```
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), \
    (d2:CensusEthnicity_Dimension2 {value:'E92000001'}), \
    (d3:CensusEthnicity_Dimension3 {value:'White: Albanian'}) \
WITH d1,d2,d3 \
MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o
```

1gb:  ~50ms  (1 results)
15gb: ~50ms  (15 results)
20gb: ~50ms  (20 results)

##### Query multiple values of the same dimension

```
    MATCH (d2:CensusEthnicity_Dimension2) \
    WHERE d2.value = 'K04000001' OR d2.value = 'W02000336' \
    WITH d2 MATCH (d1:CensusEthnicity_Dimension1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), \
    (d3:CensusEthnicity_Dimension3)<-[:has]-(o) \
    return o

```
15gb: ~200ms  (7530 results)
20gbgb: ~200ms  (10040 results)

```
MATCH (d2:CensusEthnicity_Dimension2), (d3:CensusEthnicity_Dimension3 {value:'White: Albanian'})
WHERE d2.value = 'W02000336' OR d2.value = 'K04000001'
WITH d2, d3 MATCH (d1:CensusEthnicity_Dimension1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2),
(d3:CensusEthnicity_Dimension3)<-[:has]-(o) return o
```

15gb: ~70ms  (30 results)
20gb: ~70ms  (40 results)