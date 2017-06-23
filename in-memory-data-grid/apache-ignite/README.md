### Install apache ignite (v 2.0.0)

##### Preconditions:
Install `wget` with homebrew, `brew install wget`

##### Installation:
* `wget http://mirror.catn.com/pub/apache//ignite/2.0.0/apache-ignite-2.0.0-src.zip`
* `unzip apache-ignite-2.0.0-src.zip`

##### Setup, build and run apache ignite
* Set the following environment variables:
    `export IGNITE_HOME="$HOME/apache-ignite-2.0.0-src"`
    `export JAVA_OPTS="-mx512m -ms256m -Djava.net.preferIPv4Stack=true"`
* Run `mvn clean package -DskipTests`
* Run `mvn clean package -DskipTests -Prelease,lgpl`
* Run `mvn clean package -DskipTests -Dignite.edition=hadoop`
* Locate apache-ignite-2.0.0 directory, then run `./apache-ignite-2.0.0-src/bin/ignite.sh`
and you will see similar output to this:
```
[02:49:12] Ignite node started OK (id=ab5d18a6)
[02:49:12] Topology snapshot [ver=1, nodes=1, CPUs=8, heap=1.0GB]
```
* To see Visor control run the following, `./apache-ignite-2.0.0-src/bin/ignitevisorcmd.sh`

### Performance tests for Apache Ignite (v 2.0.0)

Prior to running go script, do the following:
* Stop all nodes/instances of apache-ignite
* Locate apache-ignite-2.0.0 directory, then run `./apache-ignite-2.0.0-src/bin/ignite.sh examples/config/example-cache.xml`
* Back in current directory (dp-cmd-database-spikes/in-memory-data-grid/apache-ignite), run `go get; go build`
* Run `./apache-ignite -file-location=../../input-files/ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv -cache-name=ASH`
* Similar for other data sets but remember to change the cache name flag

#### Ingest times using the REST API
Without indexing
* UKBAA01a =   6.4 seconds (   5 MB, rows    39,424)
* RGVA01   = 102.8 seconds (  79 MB, rows   652,158)
* ASHE07E  = 249.8 seconds ( 285 MB, rows 1,486,272)
Above times include logic in go app.

Length of time spent making requests:
* UKBAA01a =   5.66 seconds
* RGVA01   =  91.35 seconds
* ASHE07E  = 214.90 seconds

Note: The above results are based on running 1 node.

To create indexes against data set:
* Go to apache-ignite-2.0.0-src
* Run `cp ~/<location of dp-cmd-database-spikes>/in-memory-data-grid/apache-ignite/example-index-cache.xml examples/config/example-index-cache.xml`
* Run `./apache-ignite-2.0.0-src/bin/ignite.sh examples/config/example-index-cache.xml`
* Reload data

With indexing
* UKBAA01a = ? seconds (   5 MB, rows    39,424)
* RGVA01   = ? seconds (  79 MB, rows   652,158)
* ASHE07E  = ? seconds ( 285 MB, rows 1,486,272)
Above times include logic in go app.

Length of time spent making requests:
* UKBAA01a = ? seconds
* RGVA01   = ? seconds
* ASHE07E  = ? seconds

#### Querying
* See documentation for sql querying [here](https://apacheignite.readme.io/v2.0/docs/rest-api#section-sql-query-execute)

Query times
* UKBAA01a =  ? seconds
* RGVA01   =  ? seconds
* ASHE07E  =  ? seconds
