package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"
	"polen/utils/common"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TopUpController struct {
	topupUC usecase.TopUpUseCase
	bioUc   usecase.BiodataUserUseCase
	rg      *gin.RouterGroup
}

func (t *TopUpController) getByIdUserId(c *gin.Context) {
	id := c.Param("id")
	// ambil role
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
	// jika bukan pemodal maka tidak boleh
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	data, err := t.topupUC.FindByIdUser(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    data,
	})
}

func (t *TopUpController) getById(c *gin.Context) {
	id := c.Param("id")
	data, err := t.topupUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    data,
	})
}

func (t *TopUpController) ConfirmUpload(c *gin.Context) {
	// ambil role
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

	// cek apakah user adalah admin
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

	// ambil data request
	var topup dto.TopUpUser
	if err := c.ShouldBindJSON(&topup); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	code, err := t.topupUC.ConfimUploadFile(topup)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success updating data",
	})
}

func (t *TopUpController) UploadedFile(c *gin.Context) {
	// ambil role
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

	// cek apakah user adalah admin
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

	data, err := t.topupUC.FindUploadedFile()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
	}

	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    data,
	})
}

func (t *TopUpController) createHandler(c *gin.Context) {
	// ambil id
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

	// ambil role
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

	// cek apakah user adalah pemodal
	if role != "pemodal" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

	// cek apakah user sudah terkonfirmasi boleh melakukan top up
	bio, err := t.bioUc.FindByUserCredential(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(bio.IsAglible, bio.Information)
	if !bio.IsAglible {
		c.JSON(403, gin.H{
			"message": []string{
				"you are not allowed to di this transaction",
				bio.Information,
			},
		})
		return
	}

	var topup dto.TopUpUser
	if err := c.ShouldBindJSON(&topup); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	topup.Id = uuid.NewString()
	topup.UserCredential.Id = ucid
	result, err := t.topupUC.CreateNew(topup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success creating data",
		"data":    result,
	})
}

func (t *TopUpController) getByIdUserLoginHandler(c *gin.Context) {
	// ambil role
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
	// jika bukan pemodal maka tidak boleh
	if role != "pemodal" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	// user id
	ucid, err := common.GetId(c)
	fmt.Println(ucid)
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
	// get data
	data, err := t.topupUC.FindByIdUser(ucid)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			// "message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success getting data",
		"data":    data,
	})
}

func (t *TopUpController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat nama file unik dengan menambahkan timestamp di belakangnya
	fileName := filepath.Base(file.Filename)
	fileExt := filepath.Ext(fileName)
	timestamp := time.Now().Format("20060102150405") // Format timestamp yang diinginkan (YYYYMMDDHHmmss)
	uniqueFileName := fileName[:len(fileName)-len(fileExt)] + "_" + timestamp + fileExt

	// Menggunakan path file yang aman
	filePath := "uploads/" + uniqueFileName

	id := c.PostForm("id")

	// get user credential
	iduc, err := common.GetId(c)
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

	// payload
	payload := dto.TopUpUser{
		Id:             id,
		File:           filePath,
		UserCredential: dto.GetAuthResponse{Id: iduc},
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, err := t.topupUC.UploadFile(payload)
	if err != nil {
		_ = os.Remove(payload.File)
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success uploaded data"})
}

func (t *TopUpController) paggingHandler(c *gin.Context) {
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

	model, pagereturn, err := t.topupUC.Pagging(payload)
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

func (t *TopUpController) Route() {
	t.rg.POST("/topup/", middleware.AuthMiddleware(), t.createHandler)
	t.rg.POST("/topup/upload", middleware.AuthMiddleware(), t.UploadFile)
	t.rg.PUT("/topup/confirm", middleware.AuthMiddleware(), t.ConfirmUpload)
	t.rg.GET("/topup/uploaded", middleware.AuthMiddleware(), t.UploadedFile)
	t.rg.GET("/topup/user/:id", middleware.AuthMiddleware(), t.getByIdUserId)
	t.rg.GET("/topup/:id", middleware.AuthMiddleware(), t.getById)
	t.rg.GET("/topup/user", middleware.AuthMiddleware(), t.getByIdUserLoginHandler)
	t.rg.GET("/topup/list/:page/:size", middleware.AuthMiddleware(), t.paggingHandler)
	// t.rg.PUT("/topup/update", t.updateHandler)

}

func NewTopUpController(topupUC usecase.TopUpUseCase, bioUc usecase.BiodataUserUseCase, rg *gin.RouterGroup) *TopUpController {
	return &TopUpController{
		topupUC: topupUC,
		rg:      rg,
		bioUc:   bioUc,
	}
}
