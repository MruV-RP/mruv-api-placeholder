package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/estates"
)

type EstatesServer struct {
	gen generator.IGenerator
}

func NewEstatesServer(gen generator.IGenerator) *EstatesServer {
	return &EstatesServer{gen: gen}
}

func (e *EstatesServer) CreateEstate(ctx context.Context, request *estates.CreateEstateRequest) (*estates.CreateEstateResponse, error) {
	return e.gen.FillWithTestData(&estates.CreateEstateResponse{}).(*estates.CreateEstateResponse), nil
}

func (e *EstatesServer) GetEstate(ctx context.Context, request *estates.GetEstateRequest) (*estates.Estate, error) {
	return e.gen.FillWithTestData(&estates.Estate{}).(*estates.Estate), nil
}

func (e *EstatesServer) UpdateEstate(ctx context.Context, request *estates.UpdateEstateRequest) (*estates.UpdateEstateResponse, error) {
	return e.gen.FillWithTestData(&estates.UpdateEstateResponse{}).(*estates.UpdateEstateResponse), nil
}

func (e *EstatesServer) DeleteEstate(ctx context.Context, request *estates.DeleteEstateRequest) (*estates.DeleteEstateResponse, error) {
	return e.gen.FillWithTestData(&estates.DeleteEstateResponse{}).(*estates.DeleteEstateResponse), nil
}

func (e *EstatesServer) GetEstates(ctx context.Context, request *estates.GetEstatesRequest) (*estates.GetEstatesResponse, error) {
	return e.gen.FillWithTestData(&estates.GetEstatesResponse{}).(*estates.GetEstatesResponse), nil
}

func (e *EstatesServer) AddGate(ctx context.Context, request *estates.AddGateRequest) (*estates.AddGateResponse, error) {
	return e.gen.FillWithTestData(&estates.AddGateResponse{}).(*estates.AddGateResponse), nil
}

func (e *EstatesServer) RemoveGate(ctx context.Context, request *estates.RemoveGateRequest) (*estates.RemoveGateResponse, error) {
	return e.gen.FillWithTestData(&estates.RemoveGateResponse{}).(*estates.RemoveGateResponse), nil
}

func (e *EstatesServer) GetEstateGates(ctx context.Context, request *estates.GetEstateGatesRequest) (*estates.GetEstateGatesResponse, error) {
	return e.gen.FillWithTestData(&estates.GetEstateGatesResponse{}).(*estates.GetEstateGatesResponse), nil
}

func (e *EstatesServer) AddEntrance(ctx context.Context, request *estates.AddEntranceRequest) (*estates.AddEntranceResponse, error) {
	return e.gen.FillWithTestData(&estates.AddEntranceResponse{}).(*estates.AddEntranceResponse), nil
}

func (e *EstatesServer) RemoveEntrance(ctx context.Context, request *estates.RemoveEntranceRequest) (*estates.RemoveEntranceResponse, error) {
	return e.gen.FillWithTestData(&estates.RemoveEntranceResponse{}).(*estates.RemoveEntranceResponse), nil
}

func (e *EstatesServer) GetEstateEntrances(ctx context.Context, request *estates.GetEstateEntrancesRequest) (*estates.GetEstateEntrancesResponse, error) {
	return e.gen.FillWithTestData(&estates.GetEstateEntrancesResponse{}).(*estates.GetEstateEntrancesResponse), nil
}

func (e *EstatesServer) FetchAll(request *estates.FetchAllEstatesRequest, server estates.MruVEstateService_FetchAllServer) error {
	panic("implement me")
}
