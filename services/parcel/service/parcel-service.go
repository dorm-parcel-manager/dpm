package service

import (
	"context"
	"fmt"

	"github.com/dorm-parcel-manager/dpm/common/appcontext"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/services/parcel/model"
	"github.com/pkg/errors"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type parcelServiceServer struct {
	pb.UnimplementedParcelServiceServer

	userService pb.UserServiceClient
	db *gorm.DB
}

func NewParcelServiceServer(db *gorm.DB, userService pb.UserServiceClient) (pb.ParcelServiceServer, error) {
	err := db.AutoMigrate(&model.Parcel{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &parcelServiceServer{
		db: db,
		userService: userService,
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

//Consult context of two roles with pro
func (s *parcelServiceServer) GetParcels(ctx context.Context, in *pb.GetParcelsRequest) (*pb.GetParcelsResponse, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	var parcels []model.Parcel;
	result := s.db.WithContext(ctx).Find(&parcels);
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}


	var apiParcels []*pb.Parcel
	for _, parcel := range parcels {
		apiParcels = append(apiParcels, mapModelToApi(&parcel))
	}
	return &pb.GetParcelsResponse{Parcels: apiParcels}, nil
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

func (s *parcelServiceServer) CreteParcel(ctx context.Context, in *pb.CreateParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStudent()
	if err != nil {
		return nil, err
	}

	data := in.Data

	var parcel = &model.Parcel{
		Owner_ID:			uint(data.OwnerId),
		Arrival_Date:		data.ArrivalDate.AsTime(),
		Transport_Company:  data.TransportCompany,
		Tracking_Number:   	data.TrackingNumber,
		Sender:      		data.Sender,
		Status: 			pb.ParcelStatus_PARCEL_INAWAIT,
	}

	result := s.db.WithContext(ctx).Create(&parcel);
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
		ID:        			uint(in.Id),
		Owner_ID:			uint(data.OwnerId),
		Arrival_Date:		data.ArrivalDate.AsTime(),
		Transport_Company:  data.TransportCompany,
		Tracking_Number:   	data.TrackingNumber,
		Sender:      		data.Sender,
		Status: 			data.Status,
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"ID",
	).Updates(parcel)

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

	parcel := &model.Parcel{
		ID:        			uint(in.Id),
		Status: 			pb.ParcelStatus_PARCEL_ACCEPTED,
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"ID",
	).Updates(parcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) StudentClaimParcel(ctx context.Context, in *pb.StudentClaimParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStudent()
	if err != nil {
		return nil, err
	}

	parcel := &model.Parcel{
		ID:        			uint(in.Id),
		Status: 			pb.ParcelStatus_PARCEL_STUDENT_CLAIMED,
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"ID",
	).Updates(parcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}

func (s *parcelServiceServer) StaffConfirmClaimParcel(ctx context.Context, in *pb.StaffConfirmClaimParcelRequest) (*pb.Empty, error) {
	appCtx := appcontext.NewAppContext(in.Context)
	err := appCtx.RequireStaff()
	if err != nil {
		return nil, err
	}

	parcel := &model.Parcel{
		ID:        			uint(in.Id),
		Status: 			pb.ParcelStatus_PARCEL_STAFF_CONFIRM_CLAIMED,
	}

	result := s.db.WithContext(ctx).Model(&parcel).Select(
		"ID",
	).Updates(parcel)

	if result.Error != nil {
		return nil, errors.WithStack(err)
	}
	return &pb.Empty{}, nil
}


func mapModelToApi(parcel *model.Parcel) *pb.Parcel {
	return &pb.Parcel{
		Id:        			int32(parcel.ID),
		OwnerId:          	int32(parcel.Owner_ID),             
		ArrivalDate:      	timestamppb.New(parcel.Arrival_Date),
		TransportCompany:  	parcel.Transport_Company,           
		TrackingNumber: 	parcel.Tracking_Number,                 
		Sender:	           	parcel.Sender,                 
		Status: 	        parcel.Status,          
	}
}

