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

type LatePaymentFeeController struct {
	LatePaymentFeeUC usecase.LatePaymentFeeUsecase
	rg               *gin.RouterGroup
}

func (p *LatePaymentFeeController) createHandler(c *gin.Context) {
	var app model.LatePaymentFee
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
	code, err := p.LatePaymentFeeUC.CreateNew(app)
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

func (p *LatePaymentFeeController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := p.LatePaymentFeeUC.Pagging(payload)
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

func (p *LatePaymentFeeController) Route() {
	p.rg.POST("/latepaymentfee", middleware.AuthMiddleware(), p.createHandler)
	p.rg.GET("/latepaymentfee/list/:page/:size", middleware.AuthMiddleware(), p.paggingHandler)
	p.rg.PUT("/latepaymentfee/", middleware.AuthMiddleware(), p.updateHandler)
	p.rg.DELETE("/latepaymentfee/:id", middleware.AuthMiddleware(), p.deleteHandler)

}

func (p *LatePaymentFeeController) deleteHandler(c *gin.Context) {
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

	err = p.LatePaymentFeeUC.DeleteById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully delete cost",
	})
}

func (p *LatePaymentFeeController) updateHandler(c *gin.Context) {
	var app model.LatePaymentFee
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

	err = p.LatePaymentFeeUC.Update(app)
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

func NewLatePaymentFeeController(LatePaymentFeeUC usecase.LatePaymentFeeUsecase, rg *gin.RouterGroup) *LatePaymentFeeController {
	return &LatePaymentFeeController{
		LatePaymentFeeUC: LatePaymentFeeUC,
		rg:               rg,
	}
}
