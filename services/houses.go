package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/houses"
)

type HousesServer struct {
	gen generator.IGenerator
}

func NewHousesServer(gen generator.IGenerator) *HousesServer {
	return &HousesServer{gen: gen}
}

func (h *HousesServer) CreateHouse(ctx context.Context, request *houses.CreateHouseRequest) (*houses.CreateHouseResponse, error) {
	panic("implement me")
}

func (h *HousesServer) GetHouse(ctx context.Context, request *houses.GetHouseRequest) (*houses.GetHouseResponse, error) {
	panic("implement me")
}

func (h *HousesServer) UpdateHouse(ctx context.Context, request *houses.UpdateHouseRequest) (*houses.UpdateHouseResponse, error) {
	panic("implement me")
}

func (h *HousesServer) DeleteHouse(ctx context.Context, request *houses.DeleteHouseRequest) (*houses.DeleteHouseResponse, error) {
	panic("implement me")
}
