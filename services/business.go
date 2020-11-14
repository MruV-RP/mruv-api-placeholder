package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/business"
	"github.com/MruV-RP/mruv-pb-go/economy"
)

type BusinessServer struct {
	gen generator.IGenerator
}

func NewBusinessServer() *BusinessServer {
	return &BusinessServer{gen: generator.SimpleGenerator{}}
}

func (b *BusinessServer) RegisterProduct(ctx context.Context, request *economy.RegisterProductRequest) (*economy.RegisterProductResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) GetProduct(ctx context.Context, request *economy.GetProductRequest) (*economy.GetProductResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) UpdateProduct(ctx context.Context, request *economy.UpdateProductRequest) (*economy.UpdateProductResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) DeleteProduct(ctx context.Context, request *economy.DeleteProductRequest) (*economy.DeleteProductResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) UpdatePrice(ctx context.Context, request *economy.UpdatePriceRequest) (*economy.UpdatePriceResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) GetPrice(ctx context.Context, request *economy.GetPriceRequest) (*economy.GetPriceResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) BuyProduct(ctx context.Context, request *economy.BuyProductRequest) (*economy.BuyProductResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) WatchProduct(request *economy.WatchProductRequest, server economy.MruVEconomyService_WatchProductServer) error {
	panic("implement me")
}

func (b *BusinessServer) WatchPrice(request *economy.WatchPriceRequest, server economy.MruVEconomyService_WatchPriceServer) error {
	panic("implement me")
}

func (b *BusinessServer) CreateBusiness(ctx context.Context, request *business.CreateBusinessRequest) (*business.CreateBusinessResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) GetBusiness(ctx context.Context, request *business.GetBusinessRequest) (*business.Business, error) {
	panic("implement me")
}

func (b *BusinessServer) UpdateBusiness(ctx context.Context, request *business.UpdateBusinessRequest) (*business.Business, error) {
	panic("implement me")
}

func (b *BusinessServer) DeleteBusiness(ctx context.Context, request *business.DeleteBusinessRequest) (*business.DeleteBusinessResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) AssignOwner(ctx context.Context, request *business.AssignOwnerRequest) (*business.AssignOwnerResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) AssignEstate(ctx context.Context, request *business.AssignEstateRequest) (*business.AssignEstateResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) UnassignEstate(ctx context.Context, request *business.UnassignEstateRequest) (*business.UnassignEstateResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) BuyBusiness(ctx context.Context, request *business.BuyBusinessRequest) (*business.BuyBusinessResponse, error) {
	panic("implement me")
}

func (b *BusinessServer) WatchBusiness(request *business.WatchBusinessRequest, server business.MruVBusinessService_WatchBusinessServer) error {
	panic("implement me")
}

func (b *BusinessServer) WatchBusinesses(request *business.WatchBusinessesRequest, server business.MruVBusinessService_WatchBusinessesServer) error {
	panic("implement me")
}
