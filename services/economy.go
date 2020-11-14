package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/economy"
)

type EconomyServer struct {
	gen generator.IGenerator
}

func NewEconomyServer(gen generator.IGenerator) *EconomyServer {
	return &EconomyServer{gen: gen}
}

func (e *EconomyServer) RegisterProduct(ctx context.Context, request *economy.RegisterProductRequest) (*economy.RegisterProductResponse, error) {
	return e.gen.FillWithTestData(&economy.RegisterProductResponse{}).(*economy.RegisterProductResponse), nil
}

func (e *EconomyServer) GetProduct(ctx context.Context, request *economy.GetProductRequest) (*economy.GetProductResponse, error) {
	return e.gen.FillWithTestData(&economy.GetProductResponse{}).(*economy.GetProductResponse), nil
}

func (e *EconomyServer) UpdateProduct(ctx context.Context, request *economy.UpdateProductRequest) (*economy.UpdateProductResponse, error) {
	return e.gen.FillWithTestData(&economy.UpdateProductResponse{}).(*economy.UpdateProductResponse), nil
}

func (e *EconomyServer) DeleteProduct(ctx context.Context, request *economy.DeleteProductRequest) (*economy.DeleteProductResponse, error) {
	return e.gen.FillWithTestData(&economy.DeleteProductResponse{}).(*economy.DeleteProductResponse), nil
}

func (e *EconomyServer) UpdatePrice(ctx context.Context, request *economy.UpdatePriceRequest) (*economy.UpdatePriceResponse, error) {
	return e.gen.FillWithTestData(&economy.UpdatePriceResponse{}).(*economy.UpdatePriceResponse), nil
}

func (e *EconomyServer) GetPrice(ctx context.Context, request *economy.GetPriceRequest) (*economy.GetPriceResponse, error) {
	return e.gen.FillWithTestData(&economy.GetPriceResponse{}).(*economy.GetPriceResponse), nil
}

func (e *EconomyServer) BuyProduct(ctx context.Context, request *economy.BuyProductRequest) (*economy.BuyProductResponse, error) {
	return e.gen.FillWithTestData(&economy.BuyProductResponse{}).(*economy.BuyProductResponse), nil
}

func (e *EconomyServer) WatchProduct(request *economy.WatchProductRequest, server economy.MruVEconomyService_WatchProductServer) error {
	panic("implement me")
}

func (e *EconomyServer) WatchPrice(request *economy.WatchPriceRequest, server economy.MruVEconomyService_WatchPriceServer) error {
	panic("implement me")
}
