package main

import (
	"github.com/Hajime3778/api-creator-backend/cmd/logger"
	"github.com/Hajime3778/api-creator-backend/cmd/server"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting("./log/")
	cfg := config.NewConfig("./config.json")
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	// collection := client.Database("api-creator-documents").Collection("test")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// id := res.InsertedID

	// log.Println(id)
}
