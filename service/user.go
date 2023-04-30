package service

import (
	"context"
	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func (u userService) Create(ctx context.Context, user model.User) (model.User, error) {
	user, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
