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

func (s *BiodataUserController) createHandler(c *gin.Context) {
	var biodata model.BiodataUser
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	biodata.Id = uuid.NewString()
	if err := s.biodataUC.CreateNew(biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, biodata)
}

func (u *BiodataUserController) listHandler(c *gin.Context) {
	biodata, err := u.biodataUC.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (u *BiodataUserController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	biodata, err := u.biodataUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (u *BiodataUserController) updateHandler(c *gin.Context) {
	var biodata model.BiodataUser
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := u.biodataUC.Update(biodata)
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

func (u *BiodataUserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.biodataUC.Delete(id); err != nil {
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

	b.rg.GET("/biodata", b.listHandler)
	b.rg.GET("/biodata/:id", b.getByIdHandler)
	b.rg.PUT("/biodata", b.updateHandler)
	b.rg.DELETE("/biodata/:id", b.deleteHandler)
}

func NewBiodataController(biodataUC usecase.BiodataUserUseCase, rg *gin.RouterGroup) *BiodataUserController {
	return &BiodataUserController{
		biodataUC: biodataUC,
		rg:        rg,
	}
}
