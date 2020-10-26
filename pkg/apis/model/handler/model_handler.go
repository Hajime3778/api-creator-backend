package handler

import (
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/model/usecase"
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
	// apiに紐づいたmodelのルート(ルーティングまとめる箇所の検討余地あり)
	r.GET("/apis/:id/models", handler.GetListByAPIID)
}

// GetAll 複数のModelを取得します
func (h *ModelHandler) GetAll(c *gin.Context) {
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

// GetByID Modelを1件取得します
func (h *ModelHandler) GetByID(c *gin.Context) {
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

// GetListByAPIID ModelをAPIIDで複数取得します
func (h *ModelHandler) GetListByAPIID(c *gin.Context) {
	apiID := c.Param("id")

	result, err := h.usecase.GetListByAPIID(apiID)

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

// Create Modelを作成します
func (h *ModelHandler) Create(c *gin.Context) {
	var model domain.Model
	c.BindJSON(&model)

	id, err := h.usecase.Create(model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// Update Modelを更新します
func (h *ModelHandler) Update(c *gin.Context) {
	var model domain.Model
	c.BindJSON(&model)

	err := h.usecase.Update(model)

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

// Delete Modelを削除します
func (h *ModelHandler) Delete(c *gin.Context) {
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
