package api

import (
	"fmt"
	"net/http"
	"polen/model"
	"polen/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BiodataUserController struct {
	biodataUC usecase.BiodataUserUseCase
	rg        *gin.RouterGroup
}

func (b *BiodataUserController) createHandler(c *gin.Context) {
	var biodata model.BiodataUser
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	biodata.Id = uuid.NewString()
	if err := b.biodataUC.CreateNew(biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, biodata)
}

func (b *BiodataUserController) listHandler(c *gin.Context) {
	biodata, err := b.biodataUC.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (b *BiodataUserController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	biodata, err := b.biodataUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (b *BiodataUserController) updateHandler(c *gin.Context) {
	var biodata model.BiodataUser
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := b.biodataUC.Update(biodata)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update biodata",
	})
}

func (b *BiodataUserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := b.biodataUC.Delete(id); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfully delete biodata with id %s", id)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (b *BiodataUserController) Route() {
	b.rg.GET("/biodata/create", b.createHandler)
	b.rg.GET("/biodata", b.listHandler)
	b.rg.GET("/biodata/:id", b.getByIdHandler)
	b.rg.PUT("/biodata/update", b.updateHandler)
	b.rg.DELETE("/biodata/:id", b.deleteHandler)
}

func NewBiodataController(biodataUC usecase.BiodataUserUseCase, rg *gin.RouterGroup) *BiodataUserController {
	return &BiodataUserController{
		biodataUC: biodataUC,
		rg:        rg,
	}
}
