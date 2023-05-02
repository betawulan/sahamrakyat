package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository/mocks"
)

func Test_UserService_Create(t *testing.T) {
	ctx := context.Background()

	type create struct {
		ctx  context.Context
		user model.User
		resp model.User
		err  error
	}

	tests := []struct {
		name       string
		argCtx     context.Context
		argUser    model.User
		createUser create
		expResp    model.User
		expErr     error
	}{
		{
			name:   "success",
			argCtx: ctx,
			argUser: model.User{
				ID:         1,
				FullName:   "John Wick",
				FirstOrder: "First Order"},
			createUser: create{
				ctx: ctx,
				user: model.User{
					ID:         1,
					FullName:   "John Wick",
					FirstOrder: "First Order",
				},
				resp: model.User{
					ID:         1,
					FullName:   "John Wick",
					FirstOrder: "First Order",
				},
				err: nil,
			},
			expResp: model.User{
				ID:         1,
				FullName:   "John Wick",
				FirstOrder: "First Order",
			},
			expErr: nil,
		},
		{
			name:   "error",
			argCtx: ctx,
			argUser: model.User{
				ID:         1,
				FullName:   "John Wick",
				FirstOrder: "First Order"},
			createUser: create{
				ctx: ctx,
				user: model.User{
					ID:         1,
					FullName:   "John Wick",
					FirstOrder: "First Order",
				},
				resp: model.User{
					ID:         1,
					FullName:   "John Wick",
					FirstOrder: "First Order",
				},
				err: errors.New("error"),
			},
			expResp: model.User{
				ID:         1,
				FullName:   "John Wick",
				FirstOrder: "First Order",
			},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)

			userRepoMock.On("Create", test.createUser.ctx, test.createUser.user).
				Return(test.createUser.resp, test.createUser.err).
				Once()

			userService := NewUserService(userRepoMock)
			response, err := userService.Create(test.argCtx, test.argUser)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expResp, response)
		})
	}
}

func Test_UserService_Update(t *testing.T) {
	ctx := context.Background()
	ID := int64(1)

	type update struct {
		ctx  context.Context
		ID   int64
		user model.User
		err  error
	}

	tests := []struct {
		name       string
		argCtx     context.Context
		argID      int64
		argUser    model.User
		updateUser update
		expErr     error
	}{
		{
			name:   "success",
			argCtx: ctx,
			argID:  ID,
			argUser: model.User{
				ID:         1,
				FullName:   "John W",
				FirstOrder: "First Order"},
			updateUser: update{
				ctx: ctx,
				ID:  ID,
				user: model.User{
					ID:         1,
					FullName:   "John W",
					FirstOrder: "First Order",
				},
				err: nil,
			},
			expErr: nil,
		},
		{
			name:   "error",
			argCtx: ctx,
			argID:  ID,
			argUser: model.User{
				ID:         1,
				FullName:   "John W",
				FirstOrder: "First Order"},
			updateUser: update{
				ctx: ctx,
				ID:  ID,
				user: model.User{
					ID:         1,
					FullName:   "John W",
					FirstOrder: "First Order",
				},
				err: errors.New("error"),
			},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)

			userRepoMock.On("Update", test.updateUser.ctx, test.updateUser.ID, test.updateUser.user).
				Return(test.updateUser.err).
				Once()

			userService := NewUserService(userRepoMock)
			err := userService.Update(test.argCtx, test.argID, test.argUser)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)

		})
	}
}

func Test_UserService_ReadByID(t *testing.T) {
	ctx := context.Background()
	ID := int64(1)

	type read struct {
		ctx  context.Context
		ID   int64
		resp model.User
		err  error
	}

	tests := []struct {
		name     string
		argCtx   context.Context
		argID    int64
		readByID read
		expResp  model.User
		expErr   error
	}{
		{
			name:   "success",
			argCtx: ctx,
			argID:  ID,
			readByID: read{
				ctx: ctx,
				ID:  1,
				resp: model.User{
					ID:         1,
					FullName:   "John W",
					FirstOrder: "First Order",
				},
				err: nil,
			},
			expResp: model.User{
				ID:         1,
				FullName:   "John W",
				FirstOrder: "First Order",
			},
			expErr: nil,
		},
		{
			name:   "error",
			argCtx: ctx,
			argID:  ID,
			readByID: read{
				ctx: ctx,
				ID:  1,
				resp: model.User{
					ID:         1,
					FullName:   "John W",
					FirstOrder: "First Order",
				},
				err: errors.New("error"),
			},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)

			userRepoMock.On("ReadByID", test.readByID.ctx, test.readByID.ID).
				Return(test.readByID.resp, test.readByID.err).
				Once()

			userService := NewUserService(userRepoMock)
			response, err := userService.ReadByID(test.argCtx, test.argID)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expResp, response)
		})
	}
}

func Test_UserService_Publish(t *testing.T) {
	ctx := context.Background()
	ID := int64(1)

	type publish struct {
		ctx context.Context
		ID  int64
		err error
	}

	tests := []struct {
		name        string
		argCtx      context.Context
		argID       int64
		publishUser publish
		expErr      error
	}{
		{
			name:   "success",
			argCtx: ctx,
			argID:  ID,
			publishUser: publish{
				ctx: ctx,
				ID:  ID,
				err: nil,
			},
			expErr: nil,
		},
		{
			name:   "error",
			argCtx: ctx,
			argID:  ID,
			publishUser: publish{
				ctx: ctx,
				ID:  ID,
				err: errors.New("error"),
			},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)

			userRepoMock.On("Publish", test.publishUser.ctx, test.publishUser.ID).
				Return(test.publishUser.err).
				Once()

			userService := NewUserService(userRepoMock)
			err := userService.Publish(test.argCtx, test.argID)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)

		})
	}
}

func Test_UserService_UnPublish(t *testing.T) {
	ctx := context.Background()
	ID := int64(1)

	type unpublish struct {
		ctx context.Context
		ID  int64
		err error
	}

	tests := []struct {
		name          string
		argCtx        context.Context
		argID         int64
		unpublishUser unpublish
		expErr        error
	}{
		{
			name:   "success",
			argCtx: ctx,
			argID:  ID,
			unpublishUser: unpublish{
				ctx: ctx,
				ID:  ID,
				err: nil,
			},
			expErr: nil,
		},
		{
			name:   "error",
			argCtx: ctx,
			argID:  ID,
			unpublishUser: unpublish{
				ctx: ctx,
				ID:  ID,
				err: errors.New("error"),
			},
			expErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)

			userRepoMock.On("UnPublish", test.unpublishUser.ctx, test.unpublishUser.ID).
				Return(test.unpublishUser.err).
				Once()

			userService := NewUserService(userRepoMock)
			err := userService.UnPublish(test.argCtx, test.argID)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)

		})
	}
}
