package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// logger.LoggingSetting("./log/")
	// cfg := config.NewConfig("../../documents.config.json")
	// db := database.NewDB(cfg)
	// conn, ctx, cancel := db.NewMongoDBConnection()
	// defer cancel()

	// collection := conn.Collection("test")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "foo", "value": 123})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// id := res.InsertedID

	// log.Println(id)

	engine := gin.Default()
	engine.Any("/*proxyPath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": c.Request.Method,
			"path":   c.Param("proxyPath"),
		})
	})
	engine.Run(":9000")
}
