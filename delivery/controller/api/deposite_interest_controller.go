package api

import (
	"net/http"
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"
	"polen/utils/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositeInterestController struct {
	depositeUC usecase.DepositeInterestUseCase
	rg         *gin.RouterGroup
}

func (d *DepositeInterestController) createHandler(c *gin.Context) {
	var deposite dto.DepositeInterestRequest
	if err := c.ShouldBindJSON(&deposite); err != nil {
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
	deposite.Id = uuid.NewString()
	code, err := d.depositeUC.CreateNew(deposite)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success creating data",
		"data":    deposite,
	})
}

func (d *DepositeInterestController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := d.depositeUC.Pagging(payload)
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

func (d *DepositeInterestController) Route() {
	d.rg.POST("/depositeinterest", middleware.AuthMiddleware(), d.createHandler)
	d.rg.GET("/depositeinterest/list/:page/:size", middleware.AuthMiddleware(), d.paggingHandler)
	d.rg.PUT("/depositeinterest/", middleware.AuthMiddleware(), d.updateHandler)
	d.rg.DELETE("/depositeinterest/:id", middleware.AuthMiddleware(), d.deleteHandler)

}

func (d *DepositeInterestController) deleteHandler(c *gin.Context) {
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

	err = d.depositeUC.DeleteById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully delete deposite",
	})
}

func (d *DepositeInterestController) updateHandler(c *gin.Context) {
	var deposite dto.DepositeInterestRequest
	if err := c.ShouldBindJSON(&deposite); err != nil {
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

	err = d.depositeUC.Update(deposite)
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
func NewDepositeInterestController(depositeinterestUC usecase.DepositeInterestUseCase, rg *gin.RouterGroup) *DepositeInterestController {
	return &DepositeInterestController{
		depositeUC: depositeinterestUC,
		rg:         rg,
	}
}
