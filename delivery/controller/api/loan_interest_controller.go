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

type LoanInterestController struct {
	loanInterestUC usecase.LoanInterestUseCase
	rg             *gin.RouterGroup
}

func (l *LoanInterestController) createHandler(c *gin.Context) {
	var loanInterest model.LoanInterest
	if err := c.ShouldBindJSON(&loanInterest); err != nil {
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
	loanInterest.Id = uuid.NewString()
	code, err := l.loanInterestUC.CreateNew(loanInterest)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success creating data",
		"data":    loanInterest,
	})
}

func (l *LoanInterestController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := l.loanInterestUC.Pagging(payload)
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

func (l *LoanInterestController) Route() {
	l.rg.POST("/loaninterest", middleware.AuthMiddleware(), l.createHandler)
	l.rg.GET("/loaninterest/list/:page/:size", middleware.AuthMiddleware(), l.paggingHandler)
	l.rg.PUT("/loaninterest/", middleware.AuthMiddleware(), l.updateHandler)
	l.rg.DELETE("/loaninterest/:id", middleware.AuthMiddleware(), l.deleteHandler)

}

func (l *LoanInterestController) deleteHandler(c *gin.Context) {
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

	err = l.loanInterestUC.DeleteById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully delete loan interest",
	})
}

func (l *LoanInterestController) updateHandler(c *gin.Context) {
	var loan model.LoanInterest
	if err := c.ShouldBindJSON(&loan); err != nil {
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

	err = l.loanInterestUC.Update(loan)
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

func NewLoanInterestController(loanInterestUC usecase.LoanInterestUseCase, rg *gin.RouterGroup) *LoanInterestController {
	return &LoanInterestController{
		loanInterestUC: loanInterestUC,
		rg:             rg,
	}
}
