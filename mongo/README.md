### Performance tests for mongodb (v 3.4.4)
Ingest times  
Using the bulk api  
ukbaa01a = 0.89 seconds (5 MB, rows 39424)  
ASHE07E = 43.81 second (285 MB, rows 1486272)  

Query times  
ukbaa01a = 21 milliseconds   
ASHE07E = 850 milliseconds / 6 milliseconds with index  


### 10GB Of Data
db.census10gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011","Dimension_Name_3":"Ethnicity"})    
Time : 250459 milliseconds
(Scanned all documents 106208140 )

db.census10gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011"})    
Time : 250459 milliseconds
(Scanned all documents 106208140)

### 20GB Of Data
Query : db.census20gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011","Dimension_Name_3":"Ethnicity"})   
Time : 592432 milliseconds (Scanned all documents 221614280 )  

Query : db.census20gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011"})    
Time : 443143 milliseconds (Scanned all documents 221614280)

### 30GB Of Data
Query : db.census30gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011","Dimension_Name_3":"Ethnicity"})   
Time : 818036 milliseconds (Scanned all documents 318624420 )

Query : db.census30gb.find({"Dimension_Value_2":"K04000001","Dimension_Value_1":"2011"})   
Time : 699677 milliseconds
(Scanned all documents 318624420