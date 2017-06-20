USING PERIODIC COMMIT
LOAD CSV WITH HEADERS FROM 'file:///CensusEthnicity5g.csv' AS line
  MERGE (d1:CensusEthnicity_Dimension1 { value: line.Dimension_Value_1 })
  MERGE (d2:CensusEthnicity_Dimension2 { value: line.Dimension_Value_2 })
  MERGE (d3:CensusEthnicity_Dimension3 { value: line.Dimension_Value_3 })
  CREATE (o:CensusEthnicity_Observation { Observation: line.Observation })
  CREATE (o)-[:has]->(d1)
  CREATE (o)-[:has]->(d2)
  CREATE (o)-[:has]->(d3)
;

//Added 10663380 labels, created 10663380 nodes, set 10663380 properties, created 31862442 relationships, completed after 1002668 ms.