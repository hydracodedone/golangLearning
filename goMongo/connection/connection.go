package connection

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database
var mongoCliet *mongo.Client

var cancelFunc context.CancelFunc
var ctx context.Context

func init() {
	GetDB()
}

func GetDB() (*mongo.Database, context.Context) {
	if mongoDB != nil {
		return mongoDB, ctx
	}
	uri := "mongodb://localhost:27017"
	timeOut := time.Second * 5
	ctx, cancelFunc = context.WithTimeout(context.Background(), timeOut)
	option := options.Client().ApplyURI(uri)
	option.SetMaxPoolSize(2)
	option.SetMaxConnIdleTime(time.Second * 30)

	var err error
	mongoCliet, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatalln(err)
	}
	if err := mongoCliet.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}
	DBName := "demo"
	mongoDB = mongoCliet.Database(DBName, options.Database())
	return mongoDB, ctx
}
func CloseConnection() {
	defer cancelFunc()

	if mongoCliet != nil {
		ctx, timeOutCancelFunc := context.WithTimeout(context.Background(), 1)
		defer timeOutCancelFunc()
		err := mongoCliet.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
