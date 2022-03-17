package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func main() {
	elasticUrl := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(elasticUrl).Do(context.Background())
	if err != nil {
		panic(err)
	}
	exists, _ := client.IndexExists("db").Do(context.Background())
	if exists {
		type BookStore struct {
			Author string `json:"author"`
			Book   string `json:"book"`
		}
		// POST
		//bs := BookStore{
		//	Author: "李莫",
		//	Book:   "AK47",
		//}
		//_, err = client.Index().Index("db").BodyJson(bs).Do(context.Background())
		//fmt.Println(err)
		_, _ = client.Search().Index("db").Do(context.Background())
		//matchQuery := elastic.NewMatchPhrasePrefixQuery("author", "Ma")
		// 只能小写
		matchQuery := elastic.NewTermQuery("author", "JoMo")
		matchQuery.CaseInsensitive(true)
		searchResult, err := client.Search().Index("db").Query(matchQuery).Do(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println(searchResult.TotalHits())
		for _, item := range searchResult.Each(reflect.TypeOf(BookStore{})) {
			fmt.Println(item.(BookStore))
		}
	} else {
		indicesCreateResult, err := client.CreateIndex("db").Do(context.Background())
		if err != nil {
			return
		}
		if !indicesCreateResult.Acknowledged {
			fmt.Println("Not Created")
			return
		}
	}
}
