package api

// import (
// 	"net/http"
// 	"polen/model"
// 	"polen/usecase"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type DepositeInterestController struct {
// 	depositeUC usecase.DepositeInterestUseCase
// 	rg         *gin.RouterGroup
// }

// func (s *DepositeInterestController) createHandler(c *gin.Context) {
// 	var deposite model.DepositeInterest
// 	if err := c.ShouldBindJSON(&deposite); err != nil {
// 		c.JSON(400, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	deposite.Id = uuid.NewString()
// 	if err := s.depositeUC.CreateNew(deposite); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, deposite)

// }
// func (d *DepositeInterestController) Route() {
// 	d.rg.GET("/depositeinterest/create", d.createHandler)

// }
// func NewDepositeInterestController(depositeinterestUC usecase.DepositeInterestUseCase, rg *gin.RouterGroup) *DepositeInterestController {
// 	return &DepositeInterestController{
// 		depositeUC: depositeinterestUC,
// 		rg:         &gin.RouterGroup{},
// 	}
// }
