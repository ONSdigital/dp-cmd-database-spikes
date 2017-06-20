USING PERIODIC COMMIT
LOAD CSV WITH HEADERS FROM 'file:///ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv' AS line
  MERGE (d1:ASHE07E_Dimension1 { value: line.Dimension_Value_1 })
  MERGE (d2:ASHE07E_Dimension2 { value: line.Dimension_Value_2 })
  MERGE (d3:ASHE07E_Dimension3 { value: line.Dimension_Value_3 })
  MERGE (d4:ASHE07E_Dimension4 { value: line.Dimension_Value_4 })
  MERGE (d5:ASHE07E_Dimension5 { value: line.Dimension_Value_5 })
  MERGE (d6:ASHE07E_Dimension6 { value: line.Dimension_Value_6 })
  CREATE (o:ASHE07E_Observation { Observation: line.Observation })
  CREATE (o)-[:has]->(d1)
  CREATE (o)-[:has]->(d2)
  CREATE (o)-[:has]->(d3)
  CREATE (o)-[:has]->(d4)
  CREATE (o)-[:has]->(d5)
  CREATE (o)-[:has]->(d6)
;

// Added 1486733 labels, created 1486733 nodes, set 851037 properties, created 8917632 relationships, completed after 172306 ms.
