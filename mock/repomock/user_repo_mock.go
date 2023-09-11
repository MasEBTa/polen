package repomock

import (
	"polen/model"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) FindByUsername(username string) (model.UserCredential, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}
func (u *UserRepoMock) Save(payload model.UserCredential) error {
	return u.Called(payload).Error(0)
}
