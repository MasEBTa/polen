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
	"time"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanUC usecase.LoanUseCase
	rg     *gin.RouterGroup
}

func (l *LoanController) createHandler(c *gin.Context) {
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
	if role != "peminjam" {
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
	var payload dto.LoanRequest
	payload.UserCredentialId = ucid
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	code, err := l.loanUC.Create(payload)
	if err != nil {
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "succes input data",
	})
}

func (l *LoanController) findById(c *gin.Context) {
	id := c.Param("id")

	data, err := l.loanUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(500, gin.H{
		"message": "Success getting data",
		"data":    data,
	})
}

func (l *LoanController) loanid(c *gin.Context) {
	id := c.Param("id")

	data, err := l.loanUC.FindByLoanId(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(500, gin.H{
		"message": "Success getting data",
		"data":    data,
	})
}

func (t *LoanController) Upload(c *gin.Context) {
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
	fmt.Println(role)

	// cek apakah user adalah admin
	if role != "peminjam" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}

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

	var payload dto.LoanInstallenmentResponse
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

	// payload
	payload.File = filePath
	payload.Id = id

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, err := t.loanUC.UploadFile(iduc, payload)
	if err != nil {
		_ = os.Remove(payload.File)
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success uploaded data"})
}

func (t *LoanController) GetUploaded(c *gin.Context) {
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
	fmt.Println(role)

	// cek apakah user adalah admin
	if role != "admin" {
		c.JSON(403, gin.H{
			"message": "you are not allowed",
		})
		return
	}
	data, err := t.loanUC.FindUploadedFile()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success grtting data",
		"data":    data,
	})
}

func (t *LoanController) confirm(c *gin.Context) {
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
	var payload dto.LoanInstallenmentResponse
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = t.loanUC.Accepted(dto.InstallenmentLoanByIdResponse{
		LoanInst: dto.LoanInstallenmentResponse{
			Id:      payload.Id,
			IsPayed: payload.IsPayed,
			Status:  payload.Status,
		},
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "updated",
	})
}

func (t *LoanController) latefee(c *gin.Context) {
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
	err = t.loanUC.UpdateLateFee()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "updated",
	})
}

func (l *LoanController) Route() {
	l.rg.POST("/loan", middleware.AuthMiddleware(), l.createHandler)
	l.rg.GET("/loan/installment/:id", middleware.AuthMiddleware(), l.findById)
	l.rg.POST("/loan/pay", middleware.AuthMiddleware(), l.Upload)
	l.rg.GET("/loan/updatedpayment", middleware.AuthMiddleware(), l.GetUploaded)
	l.rg.PUT("/loan/confirm", middleware.AuthMiddleware(), l.confirm)
	l.rg.PUT("/loan/updatelatefee", middleware.AuthMiddleware(), l.latefee)
	l.rg.GET("/loan/:id", middleware.AuthMiddleware(), l.loanid)
}

func NewLoanController(loanUC usecase.LoanUseCase, rg *gin.RouterGroup) *LoanController {
	return &LoanController{
		loanUC: loanUC,
		rg:     rg,
	}
}
