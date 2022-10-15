package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/dorm-parcel-manager/dpm/common/appcontext"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/common/utils"
	"github.com/dorm-parcel-manager/dpm/services/parcel/model"

	"github.com/pkg/errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type parcelServiceServer struct {
	pb.UnimplementedParcelServiceServer
	rabbitmqChannel *amqp.Channel
	userService     pb.UserServiceClient
	db              *gorm.DB
}

func NewParcelServiceServer(db *gorm.DB, userService pb.UserServiceClient, rabbitmqChannel *amqp.Channel) (pb.ParcelServiceServer, error) {
	err := db.AutoMigrate(&model.Parcel{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &parcelServiceServer{
		db:              db,
		userService:     userService,
		rabbitmqChannel: rabbitmqChannel,
	}, nil
}

func (s *parcelServiceServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	response, err := s.userService.Hello(ctx, in)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	message := fmt.Sprintf("User service: %v", response.Message)
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

// Consult context of two roles with pro
func (s *parcelServiceServer) GetParcels(ctx context.Context, in *pb.GetParcelsRequest) (*pb.GetParcelsResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	data := in.Data
	queryStatement := &model.Parcel{
		OwnerID:          uint(*data.OwnerId),
		ArrivalDate:      sql.NullTime{Time: data.ArrivalDate.AsTime(), Valid: true},
		Name:             *data.Name,
		TransportCompany: *data.TransportCompany,
		TrackingNumber:   *data.TrackingNumber,
		Sender:           *data.Sender,
		Status:           *data.Status,
	}

	var parcels []model.Parcel
	result := s.db.WithContext(ctx).Where(queryStatement).Find(&parcels)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiParcels []*pb.Parcel
	for _, parcel := range parcels {
		apiParcels = append(apiParcels, mapModelToApi(&parcel))
	}
	return &pb.GetParcelsResponse{Parcels: apiParcels}, nil
}

func (s *parcelServiceServer) StudentGetParcels(ctx context.Context, in *pb.StudentGetParcelsRequest) (*pb.StudentGetParcelsResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStudent()
	if err != nil {
		return nil, err
	}

	var parcels []model.Parcel
	id := in.Context.UserId
	result := s.db.WithContext(ctx).Where(&model.Parcel{OwnerID: uint(id)}).Find(&parcels)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var apiParcels []*pb.Parcel
	for _, parcel := range parcels {
		apiParcels = append(apiParcels, mapModelToApi(&parcel))
	}
	return &pb.StudentGetParcelsResponse{Parcels: apiParcels}, nil
}

func (s *parcelServiceServer) GetParcel(ctx context.Context, in *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	var parcel model.Parcel
	result := s.db.WithContext(ctx).Where(&model.Parcel{ID: uint(in.Id)}).First(&parcel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "parcel id %v not found", in.Id)
		}
	}

	return &pb.GetParcelResponse{
		Parcel: mapModelToApi(&parcel),
	}, nil
}

func (s *parcelServiceServer) CreateParcel(ctx context.Context, in *pb.CreateParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStudent()
	if err != nil {
		return nil, err
	}

	data := in.Data

	var parcel = &model.Parcel{
		OwnerID:          uint(data.OwnerId),
		Name:             data.Name,
		TransportCompany: data.TransportCompany,
		TrackingNumber:   data.TrackingNumber,
		Sender:           data.Sender,
		Status:           pb.ParcelStatus_PARCEL_REGISTERED,
	}

	result := s.db.WithContext(ctx).Create(&parcel)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) UpdateParcel(ctx context.Context, in *pb.UpdateParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	data := in.Data
	parcel := &model.Parcel{
		ID:               uint(in.Id),
		OwnerID:          uint(data.OwnerId),
		ArrivalDate:      sql.NullTime{},
		TransportCompany: data.TransportCompany,
		TrackingNumber:   data.TrackingNumber,
		Sender:           data.Sender,
		Status:           data.Status,
		Description:      data.Description,
	}

	result := s.db.WithContext(ctx).Model(&parcel).Updates(parcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) DeleteParcel(ctx context.Context, in *pb.DeleteParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireAdmin()
	if err != nil {
		return nil, err
	}

	result := s.db.WithContext(ctx).Delete(&model.Parcel{}, in.Id)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) StaffAcceptDelivery(ctx context.Context, in *pb.StaffAcceptDeliveryRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	parcel, err := s.localGetParcel(ctx, uint(in.Id))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	updatedParcel := &model.Parcel{
		ID:          uint(in.Id),
		Status:      pb.ParcelStatus_PARCEL_ARRIVED,
		ArrivalDate: sql.NullTime{Time: time.Now(), Valid: true},
		Description: in.Data.Description,
	}

	result := s.db.WithContext(ctx).Model(&updatedParcel).Select(
		"Status", "ArrivalDate", "Description",
	).Updates(updatedParcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}

	body := rabbitmq.NotificationBody{
		Title:   "Delivery arrival notification",
		Message: fmt.Sprintf("Your parcel %s have been accepted to our system.", parcel.TrackingNumber),
		Link:    "ABCDEF",
		UserID:  strconv.Itoa(int(parcel.OwnerID)),
	}

	rabbitmq.PublishNotification(ctx, s.rabbitmqChannel, &body)
	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) StudentClaimParcel(ctx context.Context, in *pb.StudentClaimParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStudent()
	if err != nil {
		return nil, err
	}

	_, err = s.localGetParcel(ctx, uint(in.Id))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	parcel := &model.Parcel{
		ID:           uint(in.Id),
		Status:       pb.ParcelStatus_PARCEL_PICKED_UP,
		PickedUpDate: sql.NullTime{Time: time.Now(), Valid: true},
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"Status",
		"PickedUpDate",
	).Updates(parcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}

func mapModelToApi(parcel *model.Parcel) *pb.Parcel {
	return &pb.Parcel{
		Id:               int32(parcel.ID),
		OwnerId:          int32(parcel.OwnerID),
		ArrivalDate:      utils.NullTimeToTimestampPb(parcel.ArrivalDate),
		PickedUpDate:     utils.NullTimeToTimestampPb(parcel.PickedUpDate),
		Name:             parcel.Name,
		TransportCompany: parcel.TransportCompany,
		TrackingNumber:   parcel.TrackingNumber,
		Sender:           parcel.Sender,
		Status:           parcel.Status,
	}
}

func (s *parcelServiceServer) localGetParcel(ctx context.Context, id uint) (*model.Parcel, error) {
	var parcel model.Parcel
	result := s.db.WithContext(ctx).Where(&model.Parcel{ID: id}).First(&parcel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "parcel id %v not found", id)
		}
	}
	return &parcel, nil
}
