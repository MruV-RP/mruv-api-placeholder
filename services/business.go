package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/business"
)

type BusinessServer struct {
	gen generator.IGenerator
}

func NewBusinessServer(gen generator.IGenerator) *BusinessServer {
	return &BusinessServer{gen: gen}
}

func (b *BusinessServer) CreateBusiness(ctx context.Context, request *business.CreateBusinessRequest) (*business.CreateBusinessResponse, error) {
	return b.gen.FillWithTestData(&business.CreateBusinessResponse{}).(*business.CreateBusinessResponse), nil
}

func (b *BusinessServer) GetBusiness(ctx context.Context, request *business.GetBusinessRequest) (*business.Business, error) {
	return b.gen.FillWithTestData(&business.Business{}).(*business.Business), nil
}

func (b *BusinessServer) UpdateBusiness(ctx context.Context, request *business.UpdateBusinessRequest) (*business.Business, error) {
	return b.gen.FillWithTestData(&business.Business{}).(*business.Business), nil
}

func (b *BusinessServer) DeleteBusiness(ctx context.Context, request *business.DeleteBusinessRequest) (*business.DeleteBusinessResponse, error) {
	return b.gen.FillWithTestData(&business.DeleteBusinessResponse{}).(*business.DeleteBusinessResponse), nil
}

func (b *BusinessServer) AssignOwner(ctx context.Context, request *business.AssignOwnerRequest) (*business.AssignOwnerResponse, error) {
	return b.gen.FillWithTestData(&business.AssignOwnerResponse{}).(*business.AssignOwnerResponse), nil
}

func (b *BusinessServer) AssignEstate(ctx context.Context, request *business.AssignEstateRequest) (*business.AssignEstateResponse, error) {
	return b.gen.FillWithTestData(&business.AssignEstateResponse{}).(*business.AssignEstateResponse), nil
}

func (b *BusinessServer) UnassignEstate(ctx context.Context, request *business.UnassignEstateRequest) (*business.UnassignEstateResponse, error) {
	return b.gen.FillWithTestData(&business.UnassignEstateResponse{}).(*business.UnassignEstateResponse), nil
}

func (b *BusinessServer) BuyBusiness(ctx context.Context, request *business.BuyBusinessRequest) (*business.BuyBusinessResponse, error) {
	return b.gen.FillWithTestData(&business.BuyBusinessResponse{}).(*business.BuyBusinessResponse), nil
}

func (b *BusinessServer) WatchBusiness(request *business.WatchBusinessRequest, server business.MruVBusinessService_WatchBusinessServer) error {
	panic("implement me")
}

func (b *BusinessServer) WatchBusinesses(request *business.WatchBusinessesRequest, server business.MruVBusinessService_WatchBusinessesServer) error {
	panic("implement me")
}
