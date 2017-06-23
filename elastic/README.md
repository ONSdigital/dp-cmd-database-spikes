### Install elasticsearch (v 5.1.1)

##### Preconditions:
Install `wget` with homebrew, `brew install wget`

##### Installation:
* `wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.1.1.zip`
* `unzip elasticsearch-5.1.1.zip`

##### Run elasticsearch and test connectivity
* Locate elasticsearh-5.1.1 directory, then run `./elasticsearh-5.1.1/bin/elasticsearch`
* `curl http://localhost:9200` or `http://127.0.0.1:9200`

### Import mapping and data sets
* `curl -XPUT http://127.0.0.1:9200/v3data -d@./config/ASHE07E.json`
* `go get;go build`
* `cd ../input-files;unzip ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv.zip; cd ../elastic`
* `./elastic -file-location=../input-files/ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv -es-dest-type=test3`

### Performance tests for Elastic search (v 5.1)
#### Ingest times using the bulk api

Without mappings (not including indexing fields)
* UKBAA01a =  1.6 seconds (5 MB, rows 39,424)
* RGVA01   = 26.8 seconds ( MB, rows 652,158)
* ASHE07E  = 89.7 seconds (285 MB, rows 1486272)
Above times include logic in go app.

Length of time spent making requests:
* UKBAA01a =  1.03 seconds
* RGVA01   = 18.22 seconds
* ASHE07E  = 63.59 seconds

Note: The optimum configuration for ingesting data has a bulk size of 30,000 and 8 workers.

With mappings (including indexing fields)
* UKBAA01a =   seconds (5 MB, rows 39,424)
* RGVA01   =  seconds ( MB, rows 652,158)
* ASHE07E  = 73.90 seconds (285 MB, rows 1486272)
Above times include logic in go app.

Length of time spent making requests:
* UKBAA01a =   seconds
* RGVA01   =  seconds
* ASHE07E  = 47.91 seconds

Query times
* UKBAA01a =  0.005 seconds
* RGVA01   =  ? seconds
* ASHE07E  =  ? seconds
