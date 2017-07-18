LOAD CSV WITH HEADERS FROM "file:///obs.csv" AS line WITH line
MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE d1.id = line.dim1
  AND d2.id = line.dim2
  AND d3.id = line.dim3
  AND d4.id = line.dim4
CREATE (o:observation { value:line.value}),
       (o)-[:isValueOf]->(d1),
       (o)-[:isValueOf]->(d2),
       (o)-[:isValueOf]->(d3),
       (o)-[:isValueOf]->(d4)