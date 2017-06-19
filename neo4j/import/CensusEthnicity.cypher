USING PERIODIC COMMIT
LOAD CSV WITH HEADERS FROM 'file:///CensusEthnicity.csv' AS line
  MERGE (d1:CensusEthnicity_Dimension1 { value: line.Dimension_Value_1 })
  MERGE (d2:CensusEthnicity_Dimension2 { value: line.Dimension_Value_2 })
  MERGE (d3:CensusEthnicity_Dimension3 { value: line.Dimension_Value_3 })
  CREATE (o:CensusEthnicity_Observation { Observation: line.Observation })
  CREATE (o)-[:has]->(d1)
  CREATE (o)-[:has]->(d2)
  CREATE (o)-[:has]->(d3)
;

//Added 10663380 labels, created 10663380 nodes, set 10663380 properties, created 31862442 relationships, completed after 1002668 ms.
//Added 10620814 labels, created 10620814 nodes, set 10620814 properties, created 31862442 relationships, completed after 1026336 ms.
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'K04000001'}) WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2) return o

MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E12000009'}), (d3:CensusEthnicity_Dimension3 {value:'Other ethnic group: Thai'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o

MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E12000009'}), (d3:CensusEthnicity_Dimension3 {value:'White: Colombian'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o
MATCH (d3:CensusEthnicity_Dimension3 {value:'White: Colombian'}) WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) return o

MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E12000009'}) WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (o)-[:has]->(d3:CensusEthnicity_Dimension3) return o, d1, d2, d3