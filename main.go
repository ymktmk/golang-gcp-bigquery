package main

import (
	"context"
	"fmt"
	"cloud.google.com/go/bigquery"
)

type Item struct {
	Name string
	Age  int
}

func main() {
	projectID := "swift-delight-344206"
	datasetID := "qiita"
	tableID := "users"

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	inserter := client.Dataset(datasetID).Table(tableID).Inserter()
	items := []*Item{
		{Name: "Tomoki", Age: 10},
		{Name: "Yamaoka", Age: 11},
	}
	if err := inserter.Put(ctx, items); err != nil {
		fmt.Println(err)
	}
}