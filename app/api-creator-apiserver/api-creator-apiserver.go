package main

import (
	"log"
	"net/http"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MongoDB↓
	// logger.LoggingSetting("./log/")
	// cfg := config.NewConfig("../../apiserver.config.json")
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

	// MySQL↓

	engine := gin.Default()
	engine.Any("/*proxyPath", func(c *gin.Context) {
		httpMethod := c.Request.Method
		// 最初の文字は/なので削除する
		url := c.Param("proxyPath")[1:]

		api, method := getRequestedAPIAndMethod(httpMethod, url)
		// リクエストされたパラメータを取得
		param := strings.Replace(url, api.URL, "", 1)

		if param != "" {
			param = param[1:]
		}

		c.JSON(http.StatusOK, gin.H{
			"param":  param,
			"method": method,
		})
	})
	engine.Run(":9000")
}

func getRequestedAPIAndMethod(httpMethod string, url string) (domain.API, domain.Method) {
	cfg := config.NewConfig("../../admin.config.json")
	mysqlDB := database.NewDB(cfg)
	mysqlConn := mysqlDB.NewMysqlConnection()
	apiRepository := _apiRepository.NewAPIRepository(mysqlConn)
	methodRepository := _methodRepository.NewMethodRepository(mysqlConn)

	//httpMethod := "GET"
	//requestURL := "my-project/api/users"
	api, err := apiRepository.GetByURL(url)
	if err != nil {
		log.Fatal(err)
	}

	methods, err := methodRepository.GetListByAPIID(api.ID)
	if err != nil {
		log.Fatal(err)
	}

	// MethodのURL部分を抽出
	requestedMethodURL := strings.Replace(url, api.URL, "", 1)

	// 区切り文字で分割する
	//params := regexp.MustCompile("[/?]").Split(methodURL, -1)
	requestedSlashCount := strings.Count(requestedMethodURL, "/")

	var returnMethod domain.Method
	// ※パラメータが2つ以上やクエリストリングは現在の仕様にないので今は考えない！
	for _, method := range methods {
		if method.Type == httpMethod {
			// methodにURLがないパターン
			if requestedMethodURL == "" && method.URL == "" {
				returnMethod = method
				break
			}

			// リクエストとMethod.URLの/数が同じものを検索する
			if requestedSlashCount == strings.Count(method.URL, "/") {
				returnMethod = method
				break
			}
		}
	}
	return api, returnMethod
}
