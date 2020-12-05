package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/api/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// APIHandler APIに対するリクエストハンドラ
type APIHandler struct {
	usecase usecase.APIUsecase
}

// NewAPIHandler APIHandlerを作成します
func NewAPIHandler(r *gin.RouterGroup, u usecase.APIUsecase) {
	handler := &APIHandler{
		usecase: u,
	}
	apiRoutes := r.Group("/apis")
	{
		apiRoutes.GET("", handler.GetAll)
		apiRoutes.GET("/:id", handler.GetByID)
		apiRoutes.POST("", handler.Create)
		apiRoutes.PUT("", handler.Update)
		apiRoutes.DELETE("/:id", handler.Delete)
	}
}

// GetAll 複数のAPIを取得します
func (h *APIHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		}
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID APIを1件取得します
func (h *APIHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.usecase.GetByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		}
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// Create APIを作成します
func (h *APIHandler) Create(c *gin.Context) {
	var api domain.API
	c.BindJSON(&api)

	status, id, err := h.usecase.Create(api)
	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// Update APIを更新します。
func (h *APIHandler) Update(c *gin.Context) {
	var api domain.API
	c.BindJSON(&api)

	status, err := h.usecase.Update(api)

	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Delete APIを削除します
func (h *APIHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	status, err := h.usecase.Delete(id)
	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
