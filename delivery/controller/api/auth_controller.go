package api

import (
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var dto dto.AuthRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	authResponse, err := a.authUC.Login(dto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Failed",
		})
		return
	}

	response := gin.H{
		"message": "successfully login",
		"data":    authResponse,
	}

	c.JSON(200, response)
}

func (a *AuthController) registerHandler(c *gin.Context) {
	var dto dto.AuthRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userUC.Register(dto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "successfully register",
	}

	c.JSON(200, response)
}

func (a *AuthController) showUserHandler(c *gin.Context) {
	name := c.Param("name")

	model, err := a.userUC.FindByUsername(name, c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "successfully getting data",
		"data": gin.H{
			"id":       model.Id,
			"username": model.Username,
			"role":     model.Role,
		},
	}

	c.JSON(200, response)
}

func (a *AuthController) paggingUserHandler(c *gin.Context) {
	// Mengambil parameter dari URL
	page, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(c.Param("size"))

	// Memberikan nilai default jika parameter kosong
	if page == 0 {
		page = 1 // Nilai default untuk page
	}

	if size == 0 {
		size = 10 // Nilai default untuk size
	}
	payload := dto.PageRequest{
		Page: page,
		Size: size,
	}

	model, pagereturn, err := a.userUC.Paging(payload, c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "Success getting data",
		"data":    model,
		"paging":  pagereturn,
	}

	c.JSON(200, response)
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
	a.rg.GET("/usercred/:name", middleware.AuthMiddleware(), a.showUserHandler)
	a.rg.GET("/user/:page/:size", middleware.AuthMiddleware(), a.paggingUserHandler)
}

func NewAuthController(userUC usecase.UserUseCase, authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		userUC: userUC,
		authUC: authUC,
		rg:     rg,
	}
}
