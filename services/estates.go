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
	panic("implement me")
}

func (e *EstatesServer) GetEstate(ctx context.Context, request *estates.GetEstateRequest) (*estates.Estate, error) {
	panic("implement me")
}

func (e *EstatesServer) UpdateEstate(ctx context.Context, request *estates.UpdateEstateRequest) (*estates.UpdateEstateResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) DeleteEstate(ctx context.Context, request *estates.DeleteEstateRequest) (*estates.DeleteEstateResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) GetEstates(ctx context.Context, request *estates.GetEstatesRequest) (*estates.GetEstatesResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) AddGate(ctx context.Context, request *estates.AddGateRequest) (*estates.AddGateResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) RemoveGate(ctx context.Context, request *estates.RemoveGateRequest) (*estates.RemoveGateResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) GetEstateGates(ctx context.Context, request *estates.GetEstateGatesRequest) (*estates.GetEstateGatesResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) AddEntrance(ctx context.Context, request *estates.AddEntranceRequest) (*estates.AddEntranceResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) RemoveEntrance(ctx context.Context, request *estates.RemoveEntranceRequest) (*estates.RemoveEntranceResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) GetEstateEntrances(ctx context.Context, request *estates.GetEstateEntrancesRequest) (*estates.GetEstateEntrancesResponse, error) {
	panic("implement me")
}

func (e *EstatesServer) FetchAll(request *estates.FetchAllEstatesRequest, server estates.MruVEstateService_FetchAllServer) error {
	panic("implement me")
}
