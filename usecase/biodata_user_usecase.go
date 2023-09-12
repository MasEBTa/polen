package usecase

import (
	"fmt"
	"polen/model"
	"polen/repository"
)

type BiodataUserUseCase interface {
	CreateNew(payload model.BiodataUser) error

	FindById(id string) (model.BiodataUser, error)
	FindByNIK(nik string) (model.BiodataUser, error)
	FindAll() ([]model.BiodataUser, error)
	Update(payload model.BiodataUser) error
	Delete(id string) error
}

type biodataUserUseCase struct {
	repo   repository.BiodataUser
	userUC UserUseCase
}

// FindById implements BiodataUserUseCase.
func (b *biodataUserUseCase) FindById(id string) (model.BiodataUser, error) {
	biodata, err := b.repo.FindById(id)
	if err != nil {
		return model.BiodataUser{}, err
	}
	return biodata, nil
}

// CreateNew implements BiodataUserUseCase.
func (b *biodataUserUseCase) CreateNew(payload model.BiodataUser) error {
	if payload.Nik == "" {
		return fmt.Errorf("nik is required")
	}
	if payload.UserCredential.Id == "" {
		return fmt.Errorf("id user is required")
	}

	if payload.NamaLengkap == "" {
		return fmt.Errorf("name is required")
	}

	if payload.NomorTelepon == "" {
		return fmt.Errorf("phone is required")
	}

	if payload.Pekerjaan == "" {
		return fmt.Errorf("job is required")
	}
	if payload.TempatLahir == "" {
		return fmt.Errorf("birth place is required")
	}

	_, err := b.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}

	err = b.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new Biodata: %v", err)
	}

	return nil
}

// Delete implements BiodataUserUseCase.
func (b *biodataUserUseCase) Delete(id string) error {
	biodata, err := b.repo.FindById(id)
	if err != nil {
		return err
	}

	err = b.repo.DeleteById(biodata.Id)
	if err != nil {
		return fmt.Errorf("failed to delete biodata: %v", err)
	}

	return nil
}

// FindAll implements BiodataUserUseCase.
func (b *biodataUserUseCase) FindAll() ([]model.BiodataUser, error) {
	return b.repo.FindAll()
}

// FindByNIK implements BiodataUserUseCase.
func (b *biodataUserUseCase) FindByNIK(nik string) (model.BiodataUser, error) {
	biodata, err := b.repo.FindByNIK(nik)
	if err != nil {
		return model.BiodataUser{}, err
	}
	return biodata, nil
}

// Update implements BiodataUserUseCase.
func (b *biodataUserUseCase) Update(payload model.BiodataUser) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Nik == "" {
		return fmt.Errorf("nik is required")
	}
	if payload.UserCredential.Id == "" {
		return fmt.Errorf("id user is required")
	}

	if payload.NamaLengkap == "" {
		return fmt.Errorf("name is required")
	}

	if payload.NomorTelepon == "" {
		return fmt.Errorf("phone is required")
	}

	if payload.Pekerjaan == "" {
		return fmt.Errorf("job is required")
	}
	if payload.TempatLahir == "" {
		return fmt.Errorf("birth place is required")
	}

	_, err := b.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}

	_, err = b.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = b.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update biodata: %v", err)
	}

	return nil
}

func NewBiodataUserUseCase(repo repository.BiodataUser, userUC UserUseCase) BiodataUserUseCase {
	return &biodataUserUseCase{
		repo:   repo,
		userUC: userUC,
	}
}
