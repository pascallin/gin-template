package conn

import (
	"context"
	"sync"
	"time"

	"github.com/pascallin/gin-template/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 5
)

type Mongo struct {
	DB     *mongo.Database
	Client *mongo.Client
}

var (
	mgOnce sync.Once
	_mongo *Mongo
)

func GetMongo(ctx context.Context) *Mongo {
	mgOnce.Do(func() {
		connectionURI := config.GetMongoConfig().URI

		opts := options.Client().ApplyURI(connectionURI).SetConnectTimeout(connectTimeout)
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}

		db := client.Database(config.GetMongoConfig().DATABASE)

		_mongo = &Mongo{DB: db, Client: client}
	})

	return _mongo
}

func Ping() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongo := GetMongo(ctx)
	err := mongo.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Error(err)
		return "fail"
	}

	return "ok"
}
