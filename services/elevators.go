package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/elevators"
)

type ElevatorsServer struct {
	gen generator.IGenerator
}

func NewElevatorsServer(gen generator.IGenerator) *ElevatorsServer {
	return &ElevatorsServer{gen: gen}
}

func (e *ElevatorsServer) CreateElevator(ctx context.Context, request *elevators.CreateElevatorRequest) (*elevators.CreateElevatorResponse, error) {
	return e.gen.FillWithTestData(&elevators.CreateElevatorResponse{}).(*elevators.CreateElevatorResponse), nil
}

func (e *ElevatorsServer) GetElevator(ctx context.Context, request *elevators.GetElevatorRequest) (*elevators.GetElevatorResponse, error) {
	return e.gen.FillWithTestData(&elevators.GetElevatorResponse{}).(*elevators.GetElevatorResponse), nil
}

func (e *ElevatorsServer) UpdateElevator(ctx context.Context, request *elevators.UpdateElevatorRequest) (*elevators.UpdateElevatorResponse, error) {
	return e.gen.FillWithTestData(&elevators.UpdateElevatorResponse{}).(*elevators.UpdateElevatorResponse), nil
}

func (e *ElevatorsServer) DeleteElevator(ctx context.Context, request *elevators.DeleteElevatorRequest) (*elevators.DeleteElevatorResponse, error) {
	return e.gen.FillWithTestData(&elevators.DeleteElevatorResponse{}).(*elevators.DeleteElevatorResponse), nil
}

func (e *ElevatorsServer) GetElevatorFloors(ctx context.Context, request *elevators.GetElevatorFloorsRequest) (*elevators.GetElevatorFloorsResponse, error) {
	return e.gen.FillWithTestData(&elevators.GetElevatorFloorsResponse{}).(*elevators.GetElevatorFloorsResponse), nil
}
