package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/elevators"
)

type ElevatorsServer struct {
}

func (e *ElevatorsServer) CreateElevator(ctx context.Context, request *elevators.CreateElevatorRequest) (*elevators.CreateElevatorResponse, error) {
	panic("implement me")
}

func (e *ElevatorsServer) GetElevator(ctx context.Context, request *elevators.GetElevatorRequest) (*elevators.GetElevatorResponse, error) {
	panic("implement me")
}

func (e *ElevatorsServer) UpdateElevator(ctx context.Context, request *elevators.UpdateElevatorRequest) (*elevators.UpdateElevatorResponse, error) {
	panic("implement me")
}

func (e *ElevatorsServer) DeleteElevator(ctx context.Context, request *elevators.DeleteElevatorRequest) (*elevators.DeleteElevatorResponse, error) {
	panic("implement me")
}

func (e *ElevatorsServer) GetElevatorFloors(ctx context.Context, request *elevators.GetElevatorFloorsRequest) (*elevators.GetElevatorFloorsResponse, error) {
	panic("implement me")
}

func NewElevatorsServer() *ElevatorsServer {
	return &ElevatorsServer{}
}
