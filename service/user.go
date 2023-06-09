package service

import (
	"context"
	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func (u userService) Publish(ctx context.Context, ID int64) error {
	err := u.userRepo.Publish(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) UnPublish(ctx context.Context, ID int64) error {
	err := u.userRepo.UnPublish(ctx, ID)
	if err != nil {
		return err
	}

	return nil
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

func (u userService) Create(ctx context.Context, user model.User) (model.User, error) {
	user, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u userService) Read(ctx context.Context, filter model.UserFilter) (model.UserResponse, error) {
	users, err := u.userRepo.Read(ctx, filter)
	if err != nil {
		return model.UserResponse{}, err
	}

	response := model.UserResponse{
		Users: users,
	}

	return response, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return userService{
		userRepo: userRepository,
	}
}
