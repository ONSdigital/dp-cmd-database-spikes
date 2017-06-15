
## Neo4j notes

### Installation / getting started

#### Run Neo4j in Docker
Docker related docs: https://neo4j.com/developer/docker/

Docker: https://hub.docker.com/_/neo4j/

```
docker run --rm \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data \
    --env=NEO4J_AUTH=none \
    --env=NEO4J_dbms_memory_pagecache_size=1G \
    --env=NEO4J_dbms_memory_heap_maxSize=512M \
    --name neo neo4j:3.2.1
```
( change memory related environment variables as required )

Navigate to http://localhost:7474/browser/ to use the browser based interface.

#### Run the importer

- cd import
- unzip the file you want to import in the ../input-files directory
- `go build`
- `./import ../../input-files/{filename}`

### Go client / driver
 - https://neo4j.com/developer/go/

Bolt - https://github.com/johnnadratowski/golang-neo4j-bolt-driver
    - uses binary bolt protocol which is more performant
    - no obvious way to batch create = slow batch inserting
    - not officially supported
    - not well documented
CQ (not evaluated) - https://github.com/go-cq/cq
    - Does not support the current version with bolt protocol

### Performance considerations
 - Does not have indexes in the traditional RDBMS sense.
 - Indexes exist in Neo4J, but are only used to find the initial node to start traversing the graph.
 
### Scalability considerations
 - clustering / sharding
 - Read perf cache based sharding - route different dataset requests to different nodes. : http://info.neo4j.com/rs/neotechnology/images/Understanding%20Neo4j%20Scalability(2).pdf
 - write perf - "Neo4j HA makes use of a single master to coordinate all write operations, and is thus limited to the write throughput of a single machine. Despite this, write throughput can still be very high." - use a queue to buffer write operations.

### Common queries

Remove all nodes from the DB: `MATCH (n) DETACH DELETE n`
Count all nodes `MATCH (n) RETURN count(n)`

### Example queries for ASHE07E dataset

```
MATCH (ds:Dataset)<-[r:isObservationOf]-(ob:Observation)
WHERE r.Geography="K02000001" AND r.Year="2015" 
AND r.Sex="CI_0006618" AND r.`Working pattern`="CI_0006618" 
AND r.`Earnings`="CI_0021537" AND r.`Earnings statistics`="CI_0006603"
RETURN r,ds,ob LIMIT 100
```