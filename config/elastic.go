package config

import (
	"gopkg.in/olivere/elastic.v7"
)

//ElasticConn ..
func ElasticConn(url string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	return client, err
}
