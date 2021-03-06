package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/method/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MethodHandler MethodAPIに対するリクエストハンドラ
type MethodHandler struct {
	usecase usecase.MethodUsecase
}

// NewMethodHandler MethodHandlerを作成します
func NewMethodHandler(r *gin.RouterGroup, u usecase.MethodUsecase) {
	handler := &MethodHandler{
		usecase: u,
	}
	methodRoutes := r.Group("/methods")
	{
		methodRoutes.GET("", handler.GetAll)
		methodRoutes.GET("/:id", handler.GetByID)
		methodRoutes.POST("", handler.Create)
		methodRoutes.PUT("", handler.Update)
		methodRoutes.DELETE("/:id", handler.Delete)
	}
	// apiに紐づいたmethodのルート(ルーティングまとめる箇所の検討余地あり)
	r.GET("/apis/:id/methods", handler.GetListByAPIID)
	r.POST("/apis/:id/create-default-methods", handler.CreateDefaultMethod)
}

// GetAll 複数のMethodを取得します
func (h *MethodHandler) GetAll(c *gin.Context) {
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

// GetByID Methodを1件取得します
func (h *MethodHandler) GetByID(c *gin.Context) {
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

// GetListByAPIID MethodをAPIIDで複数取得します
func (h *MethodHandler) GetListByAPIID(c *gin.Context) {
	apiID := c.Param("id")

	result, err := h.usecase.GetListByAPIID(apiID)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// Create Methodを作成します
func (h *MethodHandler) Create(c *gin.Context) {
	var method domain.Method
	c.BindJSON(&method)

	status, id, err := h.usecase.Create(method)
	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// CreateDefaultMethod デフォルトのCRUDMethodを作成します
func (h *MethodHandler) CreateDefaultMethod(c *gin.Context) {
	apiID := c.Param("id")

	status, methods, err := h.usecase.CreateDefaultMethods(apiID)
	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, methods)
}

// Update Methodを更新します
func (h *MethodHandler) Update(c *gin.Context) {
	var method domain.Method
	c.BindJSON(&method)

	status, err := h.usecase.Update(method)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		}
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Delete Methodを削除します
func (h *MethodHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.Delete(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		}
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
