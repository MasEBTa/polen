package api

import (
	"net/http"
	"polen/delivery/middleware"
	"polen/model"
	"polen/model/dto"
	"polen/usecase"
	"polen/utils/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppHandlingCostController struct {
	appHandlingCostUC usecase.AppHandlingCostUsecase
	rg                *gin.RouterGroup
}

func (p *AppHandlingCostController) createHandler(c *gin.Context) {
	var app model.AppHandlingCost
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	role, err := common.GetRole(c)
	if err != nil {
		if err.Error() == "unautorized" {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	app.Id = uuid.NewString()
	code, err := p.appHandlingCostUC.CreateNew(app)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success creating data",
		"data":    app,
	})
}

func (p *AppHandlingCostController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := p.appHandlingCostUC.Pagging(payload)
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

func (p *AppHandlingCostController) Route() {
	p.rg.POST("/apphandlingcost", middleware.AuthMiddleware(), p.createHandler)
	p.rg.GET("/apphandlingcost/list/:page/:size", middleware.AuthMiddleware(), p.paggingHandler)
	p.rg.PUT("/apphandlingcost/", middleware.AuthMiddleware(), p.updateHandler)
	p.rg.DELETE("/apphandlingcost/:id", middleware.AuthMiddleware(), p.deleteHandler)

}

func (p *AppHandlingCostController) deleteHandler(c *gin.Context) {
	id := c.Param("id")

	role, err := common.GetRole(c)
	if err != nil {
		if err.Error() == "unautorized" {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

	err = p.appHandlingCostUC.DeleteById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully delete app handling cost",
	})
}

func (p *AppHandlingCostController) updateHandler(c *gin.Context) {
	var app model.AppHandlingCost
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	role, err := common.GetRole(c)
	if err != nil {
		if err.Error() == "unautorized" {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

	err = p.appHandlingCostUC.Update(app)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update",
	})
}

func NewAppHandlingCostController(appHandlingCostUC usecase.AppHandlingCostUsecase, rg *gin.RouterGroup) *AppHandlingCostController {
	return &AppHandlingCostController{
		appHandlingCostUC: appHandlingCostUC,
		rg:                rg,
	}
}
