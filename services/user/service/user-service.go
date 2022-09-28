package service

import (
	"context"
	"fmt"
	"github.com/dorm-parcel-manager/dpm/common/appcontext"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/services/user/model"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer

	db *gorm.DB
}

func NewUserServiceServer(db *gorm.DB) (pb.UserServiceServer, error) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &userServiceServer{
		db: db,
	}, nil
}

func (s *userServiceServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello %v!", in.Name)
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

func (s *userServiceServer) GetUserForAuth(ctx context.Context, in *pb.GetUserForAuthRequest) (*pb.User, error) {
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

func (s *userServiceServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var users []model.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiUsers []*pb.User
	for _, user := range users {
		apiUsers = append(apiUsers, mapModelToApi(&user))
	}
	return &pb.GetUsersResponse{Users: apiUsers}, nil
}

func mapModelToApi(user *model.User) *pb.User {
	return &pb.User{
		Id:        int32(user.ID),
		OauthId:   user.OauthID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
