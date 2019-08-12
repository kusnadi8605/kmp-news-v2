# kmp-news
Golang &amp; Kafka

## Add Library
go get -v github.com/go-sql-driver/mysql  
go get -v gopkg.in/olivere/elastic.v7  
go get -v github.com/segmentio/kafka-go  
go get -v kmp-news-consumer/parser  
go get -v github.com/go-yaml/yaml  

# Running Kafka
## Running Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
## Running Kafka Server
bin/kafka-server-start.sh config/server.properties

# Running Redis
services start redis

# Running App
go run main.go

## Get News
curl -X GET 'http://localhost:3000/api/news?page=1'   

curl -X GET 'http://localhost:3000/api/news?page=2'   
