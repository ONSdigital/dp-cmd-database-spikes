CREATE CONSTRAINT ON (d:dimension) ASSERT d.id IS UNIQUE;
create (d:dimension { id:"1" });
create (d:dimension { id:"2" });
create (d:dimension { id:"3" });
create (d:dimension { id:"4" });


MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE d1.id = "1" AND d2.id = "2" AND d3.id = "3" AND d4.id = "4"
RETURN d1,d2,d3,d4


MATCH (d1:dimension), (d2:dimension), (d3:dimension), (d4:dimension)
  WHERE d1.id = "1"
  AND d2.id = "2"
  AND d3.id = "3"
  AND d4.id = "4"
CREATE (o:observation { value:"123"}),
       (o)-[:isValueOf]->(d1),
       (o)-[:isValueOf]->(d2),
       (o)-[:isValueOf]->(d3),
       (o)-[:isValueOf]->(d4)
RETURN d1,d2,d3,d4


MATCH (o:observation) return o

// delete all observations.
MATCH (o)-[r:isValueOf]->() DELETE r,o

MATCH (n:observation)-[r:isValueOf]-() DELETE r,n




profile match (n:test) where id(n) = 12150114 return n;
profile match (n:test) where n.value = 'K04000001' return n
CREATE (do:testdimensionoption {value:'male'});
CREATE (o:testobservation)-[:has]->(do:testdimensionoption {value:'male'})
CREATE (o)-[:has]->(d3);