package api

import (
	"fmt"
	"net/http"
	"polen/model"
	"polen/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositeInterestController struct {
	depositeUC usecase.DepositeInterestUseCase
	rg         *gin.RouterGroup
}

func (d *DepositeInterestController) createHandler(c *gin.Context) {
	var deposite model.DepositeInterest
	if err := c.ShouldBindJSON(&deposite); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	deposite.Id = uuid.NewString()
	if err := d.depositeUC.CreateNew(deposite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, deposite)
}
func (d *DepositeInterestController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := d.depositeUC.DeleteById(id); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfully delete deposite with id %s", id)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (d *DepositeInterestController) Route() {
	d.rg.GET("/depositeinterest/create", d.createHandler)
	d.rg.GET("/depositeinterest/:id", d.getByIdHandler)
	d.rg.PUT("/depositeinterest/update", d.updateHandler)
	d.rg.DELETE("/depositeinterest/:id", d.deleteHandler)

}

func (d *DepositeInterestController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	deposite, err := d.depositeUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, deposite)
}

func (d *DepositeInterestController) updateHandler(c *gin.Context) {
	var deposite model.DepositeInterest
	if err := c.ShouldBindJSON(&deposite); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := d.depositeUC.Update(deposite)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update deposite",
	})
}
func NewDepositeInterestController(depositeinterestUC usecase.DepositeInterestUseCase, rg *gin.RouterGroup) *DepositeInterestController {
	return &DepositeInterestController{
		depositeUC: depositeinterestUC,
		rg:         rg,
	}
}
