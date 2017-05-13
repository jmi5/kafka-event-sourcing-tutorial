### Kafka Stuff

https://semaphoreci.com/community/tutorials/writing-and-testing-an-event-sourcing-microservice-with-kafka-and-go

If we compare Kafka to a database, a table in a database is a topic in Kafka. Each table can have data expressed as a row, while in Kafka, data is simply expressed as a commit log, which is a string. Each of the commit logs has an index, aka an offset. In Kafka, the order of commit logs is important, so each one of them has an ever-increasing index number used as an offset.


Stored in: `~/personal/TechnicalLearning`

He uses the Apache Zookeeper download of Kafka (the confluent one doesn't come with a config/ folder) 
- https://zookeeper.apache.org/
- https://www.apache.org/dyn/closer.cgi?path=/kafka/0.10.2.0/kafka_2.11-0.10.2.0.tgz

### Starting Zoookeeper

cd ~/personal/Technical\ Learning
bin/zookeeper-server-start.sh config/zookeeper.properties

Edited line 48 of `kafka-run-class.sh` to be the right `base_dir="/Users/jizzard/personal/TechnicalLearning/kafka_2.11-0.10.2.0"`.
But then took that edit out - the problem was a janky file path. 

### Starting Kafka

bin/kafka-server-start.sh config/server.properties
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partition 1 --topic xbanku-transactions-t1


`cp config/server.properties config/server-1.properties`
Add these changes
broker.id=1
listeners=PLAINTEXT://:9093
log.dir=/tmp/kafka-logs-1



### Getting Govendor


- Install go: https://golang.org/doc/install?download=go1.8.1.darwin-amd64.pkg
	- Install govendor https://github.com/kardianos/govendor: `go get -u github.com/kardianos/govendor`
- Adjust your path: go puts all its stuff in a ~/go folder:
	
	export PATH="/usr/local/go/bin:$PATH"
	export GOPATH="$HOME/go"
	export PATH="$GOPATH:$PATH"
- Go's directory structure is 
	- ~/go/src/project1
	- ~/go/src/project2
	- etc...

### Compiling Banku for 1st Time

Change `package banku_test` to `package main_test`
	- Source: http://stackoverflow.com/questions/27596539/golang-testing-cant-load-package

Ran go get github.com/satori/go.uuid (inside the package so govendor picks it up) to work around the most recent error `undefined: uuid in uuid.NewV4`

### Redis stuff

Get the Redis thing and Sarama:
- pwd: `$GOPATH/src/banku` <-> `/Users/jizzard/go/src/banku`
- `go get github.com/Shopify/sarama`
- `go get github.com/go-redis/redis`


### To Start Everything Up
- `bin/zookeeper-server-start.sh config/zookeeper.properties`
- `bin/kafka-server-start.sh config/server.properties`
	- `bin/kafka-server-start.sh config/server-1.properties`

### Knobs
- kafka.go, main.go in banku
- config/server.properties in kafka home

### Final Putting It all Together
- Need a redis server running
	- `redis-server`
- Need Kafka and Zookeeper running
	- `cd ~/personal/TechnicalLearning/kafka-event-sourcing-tutorial/kafka_2.11-0.10.1.0`
	- `bin/zookeeper-server-start.sh config/zookeeper.properties`
	- `bin/kafka-server-start.sh config/server.properties`
- Run banku
	- `cd ~/go/src/banku`
	- `./banku --act=consumer`
	- `./banku`



### Links

- https://github.com/saveav/banku
- Getting syntax highlighting vim https://superuser.com/questions/252865/vim-with-colors-on-mac-command-line
