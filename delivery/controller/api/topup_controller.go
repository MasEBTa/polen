package api

import (
	"net/http"
	"polen/model"
	"polen/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TopUpController struct {
	topupUC usecase.TopUpUseCase
	rg      *gin.RouterGroup
}

func (t *TopUpController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	biodata, err := t.topupUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (t *TopUpController) createHandler(c *gin.Context) {
	var topup model.TopUp
	if err := c.ShouldBindJSON(&topup); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	topup.Id = uuid.NewString()
	if err := t.topupUC.CreateNew(c, topup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, topup)
}

func (t *TopUpController) updateHandler(c *gin.Context) {
	var topup model.TopUp
	if err := c.ShouldBindJSON(&topup); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := t.topupUC.Update(c, topup)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update topup",
	})
}

func (t *TopUpController) Route() {
	t.rg.GET("/topup/create", t.createHandler)
	t.rg.GET("/topup/:id", t.getByIdHandler)
	t.rg.PUT("/topup/update", t.updateHandler)

}

func NewTopUpController(topupUC usecase.TopUpUseCase, rg *gin.RouterGroup) *TopUpController {
	return &TopUpController{
		topupUC: topupUC,
		rg:      rg,
	}
}
