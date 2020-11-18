package handler

import (
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
	body, _ := c.GetRawData()

	response, httpStatus, err := h.usecase.RequestDocumentServer(httpMethod, url, body)

	if err != nil {
		c.JSON(httpStatus, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(httpStatus, response)
}
