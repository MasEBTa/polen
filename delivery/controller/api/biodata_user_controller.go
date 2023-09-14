package api

import (
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"
	"polen/utils/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BiodataUserController struct {
	biodataUC usecase.BiodataUserUseCase
	rg        *gin.RouterGroup
}

// func (s *BiodataUserController) createHandler(c *gin.Context) {
// 	var biodata dto.BiodataRequest
// 	if err := c.ShouldBindJSON(&biodata); err != nil {
// 		c.JSON(400, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	if err := s.biodataUC.CreateNew(biodata, c); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, biodata)
// }

func (u *BiodataUserController) listHandler(c *gin.Context) {
	biodata, err := u.biodataUC.FindByUserCredential(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, biodata)
}

func (u *BiodataUserController) listUserUpdated(c *gin.Context) {
	role, err := common.GetRole(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	biodata, err := u.biodataUC.FindUserUpdated()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}

	c.JSON(200, biodata)
}

func (u *BiodataUserController) updateAdmin(c *gin.Context) {
	var req dto.UpdateBioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	role, err := common.GetRole(c)
	if err != nil {
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
	code, err := u.biodataUC.AdminUpdate(req, c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "Internal Server Error",
		})
		return
	}
	c.JSON(code, gin.H{
		"message": "successfully update verification user",
	})
}

func (u *BiodataUserController) updateUser(c *gin.Context) {
	var biodata dto.BiodataRequest
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	code, err := u.biodataUC.UserUpdate(biodata, c)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{
		"message": "successfully update biodata",
	})
}

func (b *BiodataUserController) paggingBiodataHandler(c *gin.Context) {
	// get role
	role, err := common.GetRole(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
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

	model, pagereturn, err := b.biodataUC.Paging(payload)
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

func (b *BiodataUserController) Route() {
	// b.rg.PUT("/biodata/update", middleware.AuthMiddleware(), b.createHandler)
	b.rg.GET("/biodata", middleware.AuthMiddleware(), b.listHandler)
	b.rg.GET("/biodata/updated", middleware.AuthMiddleware(), b.listUserUpdated)
	b.rg.PUT("/biodata/update", middleware.AuthMiddleware(), b.updateUser)
	b.rg.PUT("/biodata/verified", middleware.AuthMiddleware(), b.updateAdmin)
	b.rg.GET("/biodata/list/:page/:size", middleware.AuthMiddleware(), b.paggingBiodataHandler)
}

func NewBiodataController(biodataUC usecase.BiodataUserUseCase, rg *gin.RouterGroup) *BiodataUserController {
	return &BiodataUserController{
		biodataUC: biodataUC,
		rg:        rg,
	}
}
