### Install apache ignite (v 2.0.0)

##### Preconditions:
Install `wget` with homebrew, `brew install wget`

##### Installation:
* `wget http://mirror.catn.com/pub/apache//ignite/2.0.0/apache-ignite-2.0.0-src.zip`
* `unzip apache-ignite-2.0.0-src.zip`

##### Setup, build and run apache ignite
* Set the following environment variables:
    `export IGNITE_HOME="$HOME/<apache-ignite-2.0.0-src location>"`
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

### Running java app, do the following:
* Stop all nodes/instances of apache-ignite
* Locate apache-ignite-2.0.0 directory, then run `./apache-ignite-2.0.0-src/bin/ignite.sh examples/config/example-cache.xml`
* Copy ApacheIgniteLoadinator project to IdeaProjects folder to run on IntelliJ 

Currently the loadinator spins up its own apache ignite node and once the load completes, the node is gracefully shutdown resulting in the data to be lost. However if another apache ignite node is running then the data will be distributed between the two nodes resulting in roughly half the data to be lost on shutdown of the application. So to use this application work needs to be done to update or fix app to not spin up its own node but distribute data to all other running nodes.

#### Ingest times using the java interface
Without indexing
* UKBAA01a = ?   seconds (   5 MB, rows    39,424)
* RGVA01   = ?   seconds (  79 MB, rows   652,158)
* ASHE07E  = 115 seconds ( 285 MB, rows 1,486,272)
Above times includes starting application and any logic within the loadinator application.
