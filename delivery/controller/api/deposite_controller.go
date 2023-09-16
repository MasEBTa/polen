package api

import (
	"fmt"
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"
	"polen/utils/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepositeController struct {
	depoUc usecase.DepositeUseCase
	rg     *gin.RouterGroup
}

func (d *DepositeController) createHandler(c *gin.Context) {
	// get credential
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
	if role != "pemodal" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	ucid, err := common.GetId(c)
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
	// getting input
	var deposite dto.DepositeRequest
	if err := c.ShouldBindJSON(&deposite); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// creating payload
	var payload dto.DepositeDto
	payload.UserCredential.Id = ucid
	payload.InterestRate.Id = deposite.InterestRateId
	payload.DepositeAmount = deposite.Amount
	// go to usecase
	code, err := d.depoUc.CreateDeposite(payload)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success creating data",
	})
}

func (d *DepositeController) getDepositeByUserHandler(c *gin.Context) {
	id, err := common.GetId(c)
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
	code, result, err := d.depoUc.FindByUcId(id)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    result,
	})
}

func (d *DepositeController) getDepositeByIdHandler(c *gin.Context) {
	// getting input
	id := c.Param("id")

	code, result, err := d.depoUc.FindById(id)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    result,
	})
}

func (d *DepositeController) getDepositeByUserIdHandler(c *gin.Context) {
	// getting input
	id := c.Param("id")

	// get credential
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

	code, result, err := d.depoUc.FindByUcId(id)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    result,
	})
}

func (d *DepositeController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := d.depoUc.Pagging(payload)
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

func (d *DepositeController) updateHandler(c *gin.Context) {
	// get credential
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
	// update
	fmt.Println("Melakukan pembaruan database...")
	defer fmt.Println("Selesai melakukan pembaruan data...")
	err = d.depoUc.Update()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(200, gin.H{
				"message": "Selesai melakukan pembaruan data",
			})
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (d *DepositeController) Route() {
	d.rg.POST("/deposite", middleware.AuthMiddleware(), d.createHandler)
	d.rg.GET("/deposite/user", middleware.AuthMiddleware(), d.getDepositeByUserHandler)
	d.rg.GET("/deposite/user/:id", middleware.AuthMiddleware(), d.getDepositeByUserIdHandler)
	d.rg.GET("/deposite/:id", middleware.AuthMiddleware(), d.getDepositeByIdHandler)
	d.rg.GET("/deposite/list/:page/:size", middleware.AuthMiddleware(), d.paggingHandler)
	d.rg.PUT("/deposite/update", middleware.AuthMiddleware(), d.updateHandler)
}

func NewDepositeController(depoUc usecase.DepositeUseCase, rg *gin.RouterGroup) *DepositeController {
	return &DepositeController{
		depoUc: depoUc,
		rg:     rg,
	}
}
