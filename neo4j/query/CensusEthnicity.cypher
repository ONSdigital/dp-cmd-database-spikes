
// Match on all dimensions
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E12000009'}), (d3:CensusEthnicity_Dimension3 {value:'Other ethnic group: Thai'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'W02000336'}), (d3:CensusEthnicity_Dimension3 {value:'White: Colombian'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'K04000001'}), (d3:CensusEthnicity_Dimension3 {value:'Other ethnic group: Punjabi'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E92000001'}), (d3:CensusEthnicity_Dimension3 {value:'White: Albanian'}) WITH d1,d2,d3 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3)<-[:has]-(o) return o

// Match on single dimension
MATCH (d3:CensusEthnicity_Dimension3 {value:'Other ethnic group: Somali'}) WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) return o
MATCH (d3:CensusEthnicity_Dimension3 {value:'White: Bosnian'}) WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) return o
MATCH (d3:CensusEthnicity_Dimension3 {value:'White: Brazilian'}) WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) return o
MATCH (d3:CensusEthnicity_Dimension3 {value:'White: Afghan'}) WITH d3 MATCH (d3)<-[:has]-(o:CensusEthnicity_Observation) return o

MATCH (d2:CensusEthnicity_Dimension2 {value:'E12000009'}) WITH d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (o)-[:has]->(d3:CensusEthnicity_Dimension3) return o, d1, d2, d3


// Match on year and geography
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'K04000001'}) WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2) return o
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'W02000336'}) WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2) return o
MATCH (d1:CensusEthnicity_Dimension1 {value:'2011'}), (d2:CensusEthnicity_Dimension2 {value:'E12000009'}) WITH d1,d2 MATCH (d1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (o)-[:has]->(d3:CensusEthnicity_Dimension3) return o, d1, d2, d3

// match on multiple geographies
MATCH (d2:CensusEthnicity_Dimension2) WHERE d2.value = 'K04000001' OR d2.value = 'W02000336' WITH d2 MATCH (d1:CensusEthnicity_Dimension1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3:CensusEthnicity_Dimension3)<-[:has]-(o) return o

MATCH (d2:CensusEthnicity_Dimension2), (d3:CensusEthnicity_Dimension3 {value:'White: Albanian'}) WHERE d2.value = 'W02000336' OR d2.value = 'K04000001' WITH d2 MATCH (d1:CensusEthnicity_Dimension1)<-[:has]-(o:CensusEthnicity_Observation)-[:has]->(d2), (d3:CensusEthnicity_Dimension3)<-[:has]-(o) return o
