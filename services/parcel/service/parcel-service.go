package service

import (
	"context"
	"database/sql"
	"fmt"
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

func (s *parcelServiceServer) GetParcels(ctx context.Context, in *pb.GetParcelsRequest) (*pb.GetParcelsResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	data := in.Data
	queryStatement := ""

	var parcels []model.Parcel
	result := s.db.WithContext(ctx).Where(queryStatement)

	TransportCompany := data.TransportCompany
	if TransportCompany != nil {
		result = result.Where("transport_company LIKE ? ", *TransportCompany+"%")
	}

	TrackingNumber := data.TrackingNumber
	if TrackingNumber != nil {
		result = result.Where("tracking_number LIKE ? ", *TrackingNumber+"%")
	}

	Sender := data.Sender
	if Sender != nil {
		result = result.Where("sender LIKE ? ", *Sender+"%")
	}

	result = result.Where("status", pb.ParcelStatus_PARCEL_REGISTERED)
	result = result.Find(&parcels)

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
	err := appCtx.RequireLoggedIn()
	if err != nil {
		return nil, err
	}

	var parcel model.Parcel
	where := &model.Parcel{ID: uint(in.Id)}
	if appCtx.GetUserType() == pb.UserType_TYPE_STUDENT {
		where.OwnerID = appCtx.GetUserID()
	}
	result := s.db.WithContext(ctx).Where(where).First(&parcel)
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
		CreatedAt:        sql.NullTime{Time: time.Now(), Valid: true},
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
		ID:        uint(in.Id),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	OwnerID := data.OwnerId
	if OwnerID != nil {
		parcel.OwnerID = uint(*OwnerID)
	}

	ArrivalDate := data.ArrivalDate
	if ArrivalDate != nil {
		parcel.ArrivalDate = sql.NullTime{Time: data.ArrivalDate.AsTime(), Valid: true}
	}

	TransportCompany := data.TransportCompany
	if TransportCompany != nil {
		parcel.TransportCompany = *data.TransportCompany
	}

	TrackingNumber := data.TrackingNumber
	if TrackingNumber != nil {
		parcel.TrackingNumber = *TrackingNumber
	}

	Sender := data.Sender
	if Sender != nil {
		parcel.Sender = *Sender
	}

	Status := data.Status
	if Status != nil {
		parcel.Status = *Status
	}

	Description := data.Description
	if Description != nil {
		parcel.Description = *Description
	}

	result := s.db.WithContext(ctx).Model(&parcel).Updates(parcel)

	if result.Error != nil {
		fmt.Println(result.Error)
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

	parcel, err := s.localGetParcel(ctx, &model.Parcel{ID: uint(in.Id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	updatedParcel := &model.Parcel{
		ID:          uint(in.Id),
		Status:      pb.ParcelStatus_PARCEL_ARRIVED,
		ArrivalDate: sql.NullTime{Time: time.Now(), Valid: true},
		Description: in.Data.Description,
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	fmt.Println(updatedParcel)

	result := s.db.WithContext(ctx).Model(&updatedParcel).Select(
		"Status", "ArrivalDate", "Description", "UpdatedAt",
	).Updates(updatedParcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}

	body := rabbitmq.NotificationBody{
		Title:   "Delivery arrival notification",
		Message: fmt.Sprintf("%s have  arrived.", parcel.Name),
		Link:    fmt.Sprintf("/parcels/%d", in.Id),
		UserID:  parcel.OwnerID,
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

	_, err = s.localGetParcel(ctx, &model.Parcel{ID: uint(in.Id), OwnerID: uint(in.Context.UserId)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	parcel := &model.Parcel{
		ID:           uint(in.Id),
		Status:       pb.ParcelStatus_PARCEL_PICKED_UP,
		PickedUpDate: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"Status",
		"PickedUpDate",
		"UpdatedAt",
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
		Description:      parcel.Description,
		CreatedAt:        utils.NullTimeToTimestampPb(parcel.CreatedAt),
		UpdatedAt:        utils.NullTimeToTimestampPb(parcel.UpdatedAt),
	}
}

func (s *parcelServiceServer) localGetParcel(ctx context.Context, where *model.Parcel) (*model.Parcel, error) {
	var parcel model.Parcel
	result := s.db.WithContext(ctx).Where(where).First(&parcel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "parcel id %v not found", where.ID)
		}
	}
	return &parcel, nil
}
