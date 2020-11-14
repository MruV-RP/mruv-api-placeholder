package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/vehicles"
)

type VehiclesServer struct {
	gen generator.IGenerator
}

func NewVehiclesServer(gen generator.IGenerator) *VehiclesServer {
	return &VehiclesServer{gen: gen}
}

func (v *VehiclesServer) CreateVehicle(ctx context.Context, request *vehicles.CreateVehicleRequest) (*vehicles.CreateVehicleResponse, error) {
	return v.gen.FillWithTestData(&vehicles.CreateVehicleResponse{}).(*vehicles.CreateVehicleResponse), nil
}

func (v *VehiclesServer) GetVehicle(ctx context.Context, request *vehicles.GetVehicleRequest) (*vehicles.GetVehicleResponse, error) {
	return v.gen.FillWithTestData(&vehicles.GetVehicleResponse{}).(*vehicles.GetVehicleResponse), nil
}

func (v *VehiclesServer) UpdateVehicle(ctx context.Context, request *vehicles.UpdateVehicleRequest) (*vehicles.UpdateVehicleResponse, error) {
	return v.gen.FillWithTestData(&vehicles.UpdateVehicleResponse{}).(*vehicles.UpdateVehicleResponse), nil
}

func (v *VehiclesServer) DeleteVehicle(ctx context.Context, request *vehicles.DeleteVehicleRequest) (*vehicles.DeleteVehicleResponse, error) {
	return v.gen.FillWithTestData(&vehicles.DeleteVehicleResponse{}).(*vehicles.DeleteVehicleResponse), nil
}
