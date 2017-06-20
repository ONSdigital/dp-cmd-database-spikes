
MATCH (d1:ASHE07E_Dimension1 {value:'2015'}), (d2:ASHE07E_Dimension2 {value:'K02000001'}) WITH d1 MATCH (d1)<-[:has]-(o:ASHE07E_Observation)-[:has]->(d2)

MATCH (d3:ASHE07E_Dimension3 {value:'CI_0021537'}), (d5:ASHE07E_Dimension5 {value:'CI_0005444'}), (d6:ASHE07E_Dimension6 {value:'CI_0021539'}) WITH d3,d5,d6 MATCH (d3)<-[:has]-(o:ASHE07E_Observation)-[:has]->(d5), (o)-[:has]->(d6) return count(o)

MATCH (d1:ASHE07E_Dimension1 {value:'2015'}), (d2:ASHE07E_Dimension2 {value:'K02000001'}), (d3:ASHE07E_Dimension3 {value:'CI_0021537'}), (d4:ASHE07E_Dimension4 {value:'CI_0006618'}),  (d5:ASHE07E_Dimension5 {value:'CI_0006618'}), (d6:ASHE07E_Dimension6 {value:'CI_0006603'}) WITH d1,d2,d3,d4,d5,d6 MATCH (d1)<-[:has]-(o:ASHE07E_Observation)-[:has]->(d2), (d3)<-[:has]-(o)-[:has]->(d4), (d5)<-[:has]-(o)-[:has]->(d6) return count(o)