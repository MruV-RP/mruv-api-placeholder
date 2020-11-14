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
	return h.gen.FillWithTestData(&houses.CreateHouseResponse{}).(*houses.CreateHouseResponse), nil
}

func (h *HousesServer) GetHouse(ctx context.Context, request *houses.GetHouseRequest) (*houses.GetHouseResponse, error) {
	return h.gen.FillWithTestData(&houses.GetHouseResponse{}).(*houses.GetHouseResponse), nil
}

func (h *HousesServer) UpdateHouse(ctx context.Context, request *houses.UpdateHouseRequest) (*houses.UpdateHouseResponse, error) {
	return h.gen.FillWithTestData(&houses.UpdateHouseResponse{}).(*houses.UpdateHouseResponse), nil
}

func (h *HousesServer) DeleteHouse(ctx context.Context, request *houses.DeleteHouseRequest) (*houses.DeleteHouseResponse, error) {
	return h.gen.FillWithTestData(&houses.DeleteHouseResponse{}).(*houses.DeleteHouseResponse), nil
}
