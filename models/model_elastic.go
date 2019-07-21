package models

import (
	"context"
	"fmt"
	conf "kmp-news-v2/config"
	dts "kmp-news-v2/datastruct"
	"reflect"
	"strconv"

	"gopkg.in/olivere/elastic.v7" //
)

//GetNewsElastic ..
func GetNewsElastic(conn *conf.Connection, client *elastic.Client, index string, page string, perpage int) ([]dts.NewsDetail, error) {
	var offset int

	iPage, _ := strconv.Atoi(page)
	if iPage == 1 {
		offset = 0
	} else {
		offset = iPage + (iPage - 1)
	}

	searchResult, err := client.Search().
		Index(index).               // search in index "belajar"
		Sort("id", false).          // sort by "id" field, des
		From(offset).Size(perpage). // take documents 0-x
		Pretty(true).               // pretty print request and response JSON
		Do(context.Background())    // execute

	if err != nil {
		// Handle error
		return nil, err
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var news dts.News
	arrNews := []dts.NewsDetail{}

	for _, item := range searchResult.Each(reflect.TypeOf(news)) {
		t := item.(dts.News)
		fmt.Printf("Response  %v: %s\n", t.ID, t.Created)
		// query detail news in mysql db by ID
		res, _ := GetNewsDetail(conn, t.ID)

		fmt.Println(res[0])

		arrNews = append(arrNews, res[0])
	}

	//fmt.Println(arrNews)
	return arrNews, nil
}
