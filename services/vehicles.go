package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/vehicles"
)

type VehiclesServer struct {
}

func (v *VehiclesServer) CreateVehicle(ctx context.Context, request *vehicles.CreateVehicleRequest) (*vehicles.CreateVehicleResponse, error) {
	panic("implement me")
}

func (v *VehiclesServer) GetVehicle(ctx context.Context, request *vehicles.GetVehicleRequest) (*vehicles.GetVehicleResponse, error) {
	panic("implement me")
}

func (v *VehiclesServer) UpdateVehicle(ctx context.Context, request *vehicles.UpdateVehicleRequest) (*vehicles.UpdateVehicleResponse, error) {
	panic("implement me")
}

func (v *VehiclesServer) DeleteVehicle(ctx context.Context, request *vehicles.DeleteVehicleRequest) (*vehicles.DeleteVehicleResponse, error) {
	panic("implement me")
}

func NewVehiclesServer() *VehiclesServer {
	return &VehiclesServer{}
}
