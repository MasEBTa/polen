package usecase

import (
	"fmt"
	"polen/model"
	"polen/model/dto"
	"polen/repository"
	"polen/utils/common"
	"polen/utils/security"
	"regexp"
)

type UserUseCase interface {
	FindByUsername(username string) (model.UserCredential, error)
	Register(payload dto.AuthRequest) error

	FindById(id string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// FindById implements UserUseCase.
func (u *userUseCase) FindById(id string) (model.UserCredential, error) {
	return u.repo.FindById(id)
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.repo.FindByUsername(username)
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload dto.AuthRequest) error {
	if payload.Username == "" {
		return fmt.Errorf("username required")
	}
	if payload.Password == "" {
		return fmt.Errorf("password required")
	}
	if payload.Email == "" {
		return fmt.Errorf("email required")
	}
	if payload.Role == "" {
		return fmt.Errorf("role is required")
	}
	if payload.Role != "peminjam" && payload.Role != "pemodal" {
		return fmt.Errorf("role you has choose isnt available")
	}
	// Pola regex untuk memeriksa format email
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Mencocokkan alamat email dengan pola regex
	isValid := isValidEmail(emailPattern, payload.Email)
	if !isValid {
		return fmt.Errorf("is not valid email")
	}

	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	userCredential := model.UserCredential{
		Id:       common.GenerateID(),
		Username: payload.Username,
		Password: hashPassword,
		VANumber: common.GenerateID(),
		Role:     payload.Role,
	}

	if userCredential.Role == "pemodal" {
		saldoId := common.GenerateID()
		err = u.repo.Saldo(userCredential, saldoId)
	} else {
		err = u.repo.Save(userCredential)
	}
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}

func isValidEmail(pattern, email string) bool {
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
