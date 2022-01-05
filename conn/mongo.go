package conn

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 5
)

type MongoDatabase struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}

var MongoDB *MongoDatabase

func init() {
	// load .env
	godotenv.Load()

	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s", username, password, clusterEndpoint)
	dbname := os.Getenv("MONGODB_DATABASE")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		panic(err)
	}
	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	db := client.Database(dbname)
	MongoDB = &MongoDatabase{DB: db, Client: client, Context: ctx}
	fmt.Println("mongodb connected")
}

func (d *MongoDatabase) Close() {
	d.Client.Disconnect(d.Context)
}
