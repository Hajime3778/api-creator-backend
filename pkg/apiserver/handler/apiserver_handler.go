package handler

import (
	"net/http"
	"strings"

	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/usecase"

	"github.com/gin-gonic/gin"
)

// APIServerHandler APIServerAPIに対するリクエストハンドラ
type APIServerHandler struct {
	usecase usecase.APIServerUsecase
}

// NewAPIServerHandler APIServerHandlerを作成します
func NewAPIServerHandler(engine *gin.Engine, u usecase.APIServerUsecase) {
	handler := &APIServerHandler{
		usecase: u,
	}
	engine.Any("/*proxyPath", handler.RequestDocumentServer)
}

// RequestDocumentServer リクエスト情報からAPIServerを特定し、ドキュメントに対してCRUDします
func (h *APIServerHandler) RequestDocumentServer(c *gin.Context) {

	httpMethod := c.Request.Method
	// 最初の文字は/なので削除する
	url := c.Param("proxyPath")[1:]

	api, method := h.usecase.RequestDocumentServer(httpMethod, url)
	// リクエストされたパラメータを取得
	param := strings.Replace(url, api.URL, "", 1)

	if param != "" {
		param = param[1:]
	}

	c.JSON(http.StatusOK, gin.H{
		"param":  param,
		"method": method,
	})
}
