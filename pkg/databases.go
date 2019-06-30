package dante

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties": {
			"category": {
				"type":"text"
			},
			"theme": {
				"type":"text"
			},
			"quote": {
				"type":"text"
			},
			"page": {
				"type":"integer"	
			}
		}
	}
}`

// Stores Docs in mongo DB, connects to the :host, creates the collection :name
func MongoDB(docs []Doc, host string, name string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return err
	}
	collection := client.Database("dante").Collection(name)
	var iface []interface{}
	for _, t := range docs {
		iface = append(iface, t)
	}
	_, err = collection.InsertMany(ctx, iface, &options.InsertManyOptions{})
	return err
}

// Stores Docs in
func ElasticSearch(docs []Doc, host string, name string) error {
	fmt.Printf("Trying to reach ElasticSearch at %s...\n", host)
	retries := 0
	for retries < 5 {
		_, err := elastic.NewClient(elastic.SetURL(host))
		if err == nil {
			break
		}
		retries += 1
		fmt.Println("Retrying #", retries)
		time.Sleep(time.Second * 5)
	}
	client, err := elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx := context.Background()
	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	fmt.Printf("Creating index %s\n", name)
	_, err = client.CreateIndex(name).BodyString(mapping).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, item := range docs {
		put, err := client.Index().Index(name).BodyJson(item).Do(ctx)
		if err != nil {
			return err
		}
		fmt.Println(put)
	}
	return nil
}
