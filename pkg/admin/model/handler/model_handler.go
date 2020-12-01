package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/admin/model/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ModelHandler ModelAPIに対するリクエストハンドラ
type ModelHandler struct {
	usecase usecase.ModelUsecase
}

// NewModelHandler ModelHandlerを作成します
func NewModelHandler(r *gin.RouterGroup, u usecase.ModelUsecase) {
	handler := &ModelHandler{
		usecase: u,
	}
	modelRoutes := r.Group("/models")
	{
		modelRoutes.GET("", handler.GetAll)
		modelRoutes.GET("/:id", handler.GetByID)
		modelRoutes.POST("", handler.Create)
		modelRoutes.PUT("", handler.Update)
		modelRoutes.DELETE("/:id", handler.Delete)
	}
	// apiに紐づいたmodel(ルーティングまとめる箇所の検討余地あり)
	r.GET("/apis/:id/model", handler.GetByAPIID)
}

// GetAll 複数のModelを取得します
func (h *ModelHandler) GetAll(c *gin.Context) {
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

// GetByID Modelを1件取得します
func (h *ModelHandler) GetByID(c *gin.Context) {
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

// GetByAPIID APIIDからModelを1件取得します
func (h *ModelHandler) GetByAPIID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.usecase.GetByAPIID(id)

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

// Create Modelを作成します
func (h *ModelHandler) Create(c *gin.Context) {
	var model domain.Model
	c.BindJSON(&model)

	status, id, err := h.usecase.Create(model)
	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// Update Modelを更新します
func (h *ModelHandler) Update(c *gin.Context) {
	var model domain.Model
	c.BindJSON(&model)

	status, err := h.usecase.Update(model)

	if err != nil {
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Delete Modelを削除します
func (h *ModelHandler) Delete(c *gin.Context) {
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
