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

func (s *userServiceServer) GetUsers(ctx context.Context, in *api.Empty) (*api.UserList, error) {
	var users []model.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiUsers []*api.User
	for _, user := range users {
		apiUsers = append(apiUsers, &api.User{
			Id:        int32(user.ID),
			OauthId:   user.OauthID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	return &api.UserList{Users: apiUsers}, nil
}
