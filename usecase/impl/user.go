package impl

import (
	"errors"

	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
	"github.com/uzimaru0000/poker/usecase"
)

type userUsecase struct {
	userRepo      repository.UserRepository
	uuidGenerator lib.UUIDGenerator
}

func NewUserUsecase(userRepo repository.UserRepository, uuidGenerator lib.UUIDGenerator) usecase.UserUsecase {
	return &userUsecase{
		userRepo:      userRepo,
		uuidGenerator: uuidGenerator,
	}
}

func (uu *userUsecase) CreateUser(user *model.User) (*model.User, error) {
	id, err := uu.uuidGenerator.Generate()
	if err != nil {
		return nil, err
	}

	user.ID = id
	err = uu.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) GetUser(user *model.User) (*model.User, error) {
	return uu.userRepo.GetUser(user)
}

func (uu *userUsecase) UpdateUser(user *model.User) (*model.User, error) {
	return nil, errors.New("Failed")
}

func (uu *userUsecase) DeleteUser(user *model.User) error {
	return errors.New("Failed")
}
