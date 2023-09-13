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

func (u *TopUpController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	biodata, err := u.topupUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (s *TopUpController) createHandler(c *gin.Context) {
	var topup model.TopUp
	if err := c.ShouldBindJSON(&topup); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	topup.Id = uuid.NewString()
	if err := s.topupUC.CreateNew(topup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, topup)
}

func (t *TopUpController) Route() {
	t.rg.GET("/topup/create", t.createHandler)
	t.rg.GET("/topup/:id", t.getByIdHandler)

}

func NewTopUpController(topupUC usecase.TopUpUseCase, rg *gin.RouterGroup) *TopUpController {
	return &TopUpController{
		topupUC: topupUC,
		rg:      rg,
	}
}
