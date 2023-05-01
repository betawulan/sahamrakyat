package service

import (
	"context"
	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func (u userService) ReadByID(ctx context.Context, ID int64) (model.User, error) {
	user, err := u.userRepo.ReadByID(ctx, ID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u userService) Update(ctx context.Context, ID int64, user model.User) error {
	err := u.userRepo.Update(ctx, ID, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) Delete(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Create(ctx context.Context, user model.User) (model.User, error) {
	user, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{
		userRepo: userRepository,
	}
}
