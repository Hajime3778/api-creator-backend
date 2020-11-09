package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// logger.LoggingSetting("./log/")
	// cfg := config.NewConfig("./documents.config.json")
	// db := database.DB{
	// 	Host:       cfg.DataBase.Host,
	// 	Port:       cfg.DataBase.Port,
	// 	Username:   cfg.DataBase.User,
	// 	Password:   cfg.DataBase.Password,
	// 	DBName:     cfg.DataBase.Database,
	// 	Connection: cfg.DataBase.Host,
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	collection := client.Database("api-creator-documents").Collection("test")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatalln(err)
		return
	}
	id := res.InsertedID

	log.Println(id)
}
