package service

import (
	"context"
	"fmt"

	"github.com/dorm-parcel-manager/dpm/common/appcontext"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/services/user/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	result := s.db.WithContext(ctx).Where(&model.User{OauthID: in.OauthId}).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(result.Error)
		}
		user.OauthID = in.OauthId
		user.Email = in.Email
		user.FirstName = in.FirstName
		user.LastName = in.LastName
		user.Picture = in.Picture
		result = s.db.WithContext(ctx).Create(&user)
		if result.Error != nil {
			return nil, errors.WithStack(result.Error)
		}
	}
	return mapModelToApi(&user), nil
}

func (s *userServiceServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.UserInfo, error) {
	var user model.User
	result := s.db.WithContext(ctx).Where(&model.User{ID: uint(in.Id)}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user id %v not found", in.Id)
		}
		return nil, result.Error
	}
	return &pb.UserInfo{
		Id:        uint32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
		Type:      user.Type,
	}, nil
}

func (s *userServiceServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	var users []model.User
	result := s.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiUsers []*pb.User
	for _, user := range users {
		apiUsers = append(apiUsers, mapModelToApi(&user))
	}
	return &pb.GetUsersResponse{Users: apiUsers}, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	var user model.User
	result := s.db.WithContext(ctx).Where(&model.User{ID: uint(in.Id)}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user id %v not found", in.Id)
		}
	}

	return &pb.GetUserResponse{
		User: mapModelToApi(&user),
	}, nil
}

func (s *userServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	data := in.Data
	user := &model.User{
		ID:        uint(in.Id),
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Picture:   data.Picture,
		Type:      data.Type,
	}
	result := s.db.WithContext(ctx).Model(&user).Select(
		"Email", "FirstName", "LastName", "Type",
	).Updates(user)
	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}

func (s *userServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	result := s.db.WithContext(ctx).Delete(&model.User{}, in.Id)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return &pb.Empty{}, nil
}

func mapModelToApi(user *model.User) *pb.User {
	return &pb.User{
		Id:        int32(user.ID),
		OauthId:   user.OauthID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
		Type:      user.Type,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
