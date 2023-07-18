package models

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

var Mongo = InitMongo()
var RDB = InitRedis()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://192.168.12.133:27017"))

	if err != nil {
		log.Printf("Connect Mongo Error:", err)
	}
	return client.Database("im")
}
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "192.168.12.133:6379",
	})
}
