package mock

import (
	"polen/model"
	"polen/model/dto"
	"time"
)

var MockUserCred = model.UserCredential{
	Id:       "1",
	Username: "akbar",
	Email:    "akbarismail@gmail.com",
	Password: "123",
	VANumber: "Efvfdvfdhucsucuh",
	Role:     "borrower",
	IsActive: true,
}
var MockAuthReq = dto.AuthRequest{
	Username: "akbaris",
	Email:    "akbar@gmail.com",
	Password: "password",
	Role:     "borrower",
}
var MockAuthResponse = dto.AuthResponse{
	Username: MockAuthReq.Username,
	Token:    "",
}
var MockBiodata = model.BiodataUser{
	Id:             "1",
	UserCredential: model.UserCredential{Id: MockUserCred.Id},
	NamaLengkap:    "akbar ismail",
	Nik:            "32010",
	NomorTelepon:   "081287743960",
	Pekerjaan:      "IT",
	TempatLahir:    "jakarta",
	TanggalLahir:   time.Date(2000, time.December, 12, 0, 0, 0, 0, time.UTC),
	KodePos:        "1610",
	IsAglible:      false,
	StatusUpdate:   false,
	Information:    "biodata is not updated",
}
var MockSaldo = model.Saldo{
	Id: "1",
	UserCredential: model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Email:    "akbarismail@gmail.com",
		Password: "123",
		VANumber: "Efvfdvfdhucsucuh",
		Role:     "borrower",
		IsActive: true,
	},
	TotalSaving: 0,
}
var MockPageReq = dto.PageRequest{
	Page: 1,
	Size: 5,
}
var MockPaging = dto.Paging{
	Page:       1,
	Size:       5,
	TotalRows:  1,
	TotalPages: 1,
}
var MockUserCreds = []model.UserCredential{
	{
		Id:       "1",
		Username: "akbar",
		Email:    "akbarismail@gmail.com",
		Password: "123",
		VANumber: "Efvfdvfdhucsucuh",
		Role:     "borrower",
		IsActive: true,
	},
}
var MockBiodataResponse = dto.BiodataResponse{
	Id:          "1",
	NamaLengkap: "akbar ismail",
	UserCredential: dto.GetAuthResponse{
		Id:       "1",
		Username: "akbaris",
		Email:    "akbar@gmail.com",
		Password: "123",
		Role:     "peminjam",
		VaNumber: "bfdffbfvfhvf",
		IsActive: false,
	},
	Nik:          "32010",
	NomorTelepon: "081287743960",
	Pekerjaan:    "IT",
	TempatLahir:  "Jakarta",
	TanggalLahir: "2000-12-12",
	KodePos:      "1610",
	IsAglible:    false,
	StatusUpdate: false,
	Information:  "Additional",
}
var MockBiodataResponses = []dto.BiodataResponse{
	{
		Id:          "1",
		NamaLengkap: "akbar ismail",
		UserCredential: dto.GetAuthResponse{
			Id:       "1",
			Username: "akbaris",
			Email:    "akbar@gmail.com",
			Password: "123",
			Role:     "peminjam",
			VaNumber: "bfdffbfvfhvf",
			IsActive: false,
		},
		Nik:          "32010",
		NomorTelepon: "081287743960",
		Pekerjaan:    "IT",
		TempatLahir:  "Jakarta",
		TanggalLahir: "2000-12-12",
		KodePos:      "1610",
		IsAglible:    false,
		StatusUpdate: false,
		Information:  "Additional",
	},
}
var MockDepositeInterest = model.DepositeInterest{
	Id:           "1",
	CreateDate:   time.Now(),
	InterestDate: time.Now(),
	Duration:     30,
}
