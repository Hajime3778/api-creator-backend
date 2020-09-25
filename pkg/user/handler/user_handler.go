package handler

import (
	"log"
	"net/http"
	"strconv"

	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserHandler handler for user request
type UserHandler struct {
	usecase usecase.UserUsecase
}

// NewUserHandler is init for UserHandler
func NewUserHandler(r *gin.RouterGroup, u usecase.UserUsecase) {
	handler := &UserHandler{
		usecase: u,
	}
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", handler.GetAll)
		userRoutes.GET("/:id", handler.GetByID)
		userRoutes.POST("", handler.Create)
		userRoutes.PUT("", handler.Update)
		userRoutes.DELETE("/:id", handler.Delete)
	}
}

// GetAll 複数のUserを取得します
func (h *UserHandler) GetAll(c *gin.Context) {
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

// GetByID 1件のUserを取得します
func (h *UserHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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

// Create Userを作成します
func (h *UserHandler) Create(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	id, err := h.usecase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, domain.CreatedResponse{ID: id})
}

// Update Userを更新します。
func (h *UserHandler) Update(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	err := h.usecase.Update(user)

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

// Delete Userを削除します
func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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
