package main

import (
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/handler"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// #region MongoDBテスト用のコード
	// logger.LoggingSetting("./log/")
	// cfg := config.NewConfig("../../apiserver.config.json")
	// db := database.NewDB(cfg)
	// conn, ctx, cancel := db.NewMongoDBConnection()
	// defer cancel()
	//
	// collection := conn.Collection("test")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "foo", "value": 123})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// id := res.InsertedID
	//
	// log.Println(id)
	// #endregion

	cfg := config.NewConfig("../../admin.config.json")
	mysqlDB := database.NewDB(cfg)
	mysqlConn := mysqlDB.NewMysqlConnection()
	apiRepository := _apiRepository.NewAPIRepository(mysqlConn)
	methodRepository := _methodRepository.NewMethodRepository(mysqlConn)

	engine := gin.Default()

	apiserverUsecase := usecase.NewAPIServerUsecase(apiRepository, methodRepository)
	handler.NewAPIServerHandler(engine, apiserverUsecase)

	engine.Run(":9000")
}
