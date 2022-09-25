package service

import (
	"context"

	"github.com/dorm-parcel-manager/dpm/pkg/api"
	"github.com/dorm-parcel-manager/dpm/pkg/user/model"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type userServiceServer struct {
	api.UnimplementedUserServiceServer

	db *gorm.DB
}

func NewUserServiceServer(db *gorm.DB) api.UserServiceServer {
	db.AutoMigrate(&model.User{})
	return &userServiceServer{
		db: db,
	}
}

func (s *userServiceServer) Hello(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func (s *userServiceServer) GetUserForAuth(ctx context.Context, in *api.GetUserForAuthRequest) (*api.User, error) {
	var user model.User
	result := s.db.Where(&model.User{OauthID: in.OauthId}).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(result.Error)
		}
		user.OauthID = in.OauthId
		user.Email = in.Email
		user.FirstName = in.FirstName
		user.LastName = in.LastName
		result = s.db.Create(&user)
		if result.Error != nil {
			return nil, errors.WithStack(result.Error)
		}
	}
	return mapModelToApi(&user), nil
}

func (s *userServiceServer) GetUsers(ctx context.Context, in *api.Empty) (*api.UserList, error) {
	var users []model.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiUsers []*api.User
	for _, user := range users {
		apiUsers = append(apiUsers, mapModelToApi(&user))
	}
	return &api.UserList{Users: apiUsers}, nil
}

func mapModelToApi(user *model.User) *api.User {
	return &api.User{
		Id:        int32(user.ID),
		OauthId:   user.OauthID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
