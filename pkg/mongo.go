package tigbra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FeedMongoDB(docs []Doc, host string, name string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return err
	}
	collection := client.Database("tigbra").Collection(name)
	var iface []interface{}
	for _, t := range docs {
		iface = append(iface, t)
	}
	_, err = collection.InsertMany(ctx, iface, &options.InsertManyOptions{})
	return err
}
