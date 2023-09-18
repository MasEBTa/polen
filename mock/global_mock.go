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
	Id:    "1",
	UcId:  "1",
	Total: 0,
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
var MockDepositeDto = []dto.Deposite{
	{
		Id:             "1",
		DepositeAmount: 100,
		InterestRate:   1,
		TaxRate:        1,
		DurationMounth: 12,
		CreateDate:     time.Now(),
		MaturityDate:   time.Now(),
		Status:         "pending",
		GrossProfit:    12000,
		Tax:            10,
		NetProfit:      2000,
		TotalReturn:    10000,
	},
}
var MockTopUpByUser = dto.TopUpByUser{
	UserCredential: dto.GetAuthResponse{
		Id:       "1",
		Username: "akbar",
		Email:    "akbarismail@gmail.com",
		Password: "123",
		VaNumber: "Efvfdvfdhucsucuh",
		Role:     "borrower",
		IsActive: true,
	},
	UserBio: dto.BiodataRequest{
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081287743960",
		Pekerjaan:    "IT",
		TempatLahir:  "Jakarta",
		TanggalLahir: "2000-12-12",
		KodePos:      "1610",
	},
	TopUp: []dto.TopUp{
		{
			Id:                    "1",
			VaNumber:              "Efvfdvfdhucsucuh",
			TopUpAmount:           100,
			MaturityTime:          time.Now(),
			AcceptedTime:          time.Now(),
			Accepted:              false,
			Status:                "pending",
			TransferConfirmRecipe: false,
			File:                  "",
		},
	},
}
var MockTopUpId = dto.TopUpById{
	UserCredential: dto.GetAuthResponse{
		Id:       "1",
		Username: "akbar",
		Email:    "akbarismail@gmail.com",
		Password: "123",
		VaNumber: "Efvfdvfdhucsucuh",
		Role:     "borrower",
		IsActive: true,
	},
	UserBio: dto.BiodataRequest{
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081287743960",
		Pekerjaan:    "IT",
		TempatLahir:  "Jakarta",
		TanggalLahir: "2000-12-12",
		KodePos:      "1610",
	},
	TopUp: dto.TopUp{
		Id:                    "1",
		VaNumber:              "Efvfdvfdhucsucuh",
		TopUpAmount:           100,
		MaturityTime:          time.Now(),
		AcceptedTime:          time.Now(),
		Accepted:              false,
		Status:                "pending",
		TransferConfirmRecipe: false,
		File:                  "",
	},
}
var MockTopUp = model.TopUp{
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
	TopUpAmount:           100,
	MaturityTime:          time.Now(),
	AcceptedTime:          time.Now(),
	Accepted:              false,
	Status:                "pending",
	TransferConfirmRecipe: false,
	File:                  "",
}
var MockListTopUp = []dto.TopUp{
	{
		Id:                    "1",
		VaNumber:              "vfdbvhfdbdhf",
		TopUpAmount:           1000,
		MaturityTime:          time.Now(),
		AcceptedTime:          time.Now(),
		Accepted:              false,
		Status:                "not accepted",
		TransferConfirmRecipe: false,
		File:                  "",
	},
}
var MockDeposites = []dto.DepositeInterestRequest{
	{
		Id:             "1",
		InterestRate:   1,
		TaxRate:        1,
		DurationMounth: 12,
	},
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
	Id:             "1",
	CreateDate:     time.Now(),
	InterestRate:   1,
	TaxRate:        1,
	DurationMounth: 12,
}
var MockDepositeInterestReq = dto.DepositeInterestRequest{
	Id:             "1",
	InterestRate:   1,
	TaxRate:        1,
	DurationMounth: 12,
}
var MockDepositeByIdResponse = dto.DepositeByIdResponse{
	BioUser: dto.BiodataResponse{
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
	Deposite: dto.Deposite{
		Id:             "1",
		DepositeAmount: 100,
		InterestRate:   1,
		TaxRate:        1,
		DurationMounth: 12,
		CreateDate:     time.Now(),
		MaturityDate:   time.Now(),
		Status:         "pending",
		GrossProfit:    12000,
		Tax:            10,
		NetProfit:      2000,
		TotalReturn:    10000,
	},
}
