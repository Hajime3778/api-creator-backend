package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/method/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MethodHandler handler for method request
type MethodHandler struct {
	usecase usecase.MethodUsecase
}

// NewMethodHandler is init for MethodHandler
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
}

// GetAll 複数のMethodを取得します
func (h *MethodHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID 1件のMethodを取得します
func (h *MethodHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.usecase.GetByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Create Methodを作成します
func (h *MethodHandler) Create(c *gin.Context) {
	var method domain.Method
	c.BindJSON(&method)

	id, err := h.usecase.Create(method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// Update Methodを更新します。
func (h *MethodHandler) Update(c *gin.Context) {
	var method domain.Method
	c.BindJSON(&method)

	err := h.usecase.Update(method)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
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
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
