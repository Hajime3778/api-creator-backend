package main

import (
	"log"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// logger.LoggingSetting("./log/")
	cfg := config.NewConfig("../../documents.config.json")
	db := database.NewDB(cfg)
	conn, ctx, cancel := db.NewMongoDBConnection()
	defer cancel()

	collection := conn.Collection("test")
	res, err := collection.InsertOne(ctx, bson.M{"name": "foo", "value": 123})
	if err != nil {
		log.Fatalln(err)
		return
	}
	id := res.InsertedID

	log.Println(id)
}
