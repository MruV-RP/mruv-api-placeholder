package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/entrances"
)

type EntrancesServer struct {
	gen generator.IGenerator
}

func NewEntrancesServer(gen generator.IGenerator) *EntrancesServer {
	return &EntrancesServer{gen: gen}
}

func (e *EntrancesServer) CreateEntrance(ctx context.Context, request *entrances.CreateEntranceRequest) (*entrances.CreateEntranceResponse, error) {
	return e.gen.FillWithTestData(&entrances.CreateEntranceResponse{}).(*entrances.CreateEntranceResponse), nil
}

func (e *EntrancesServer) GetEntrance(ctx context.Context, request *entrances.GetEntranceRequest) (*entrances.GetEntranceResponse, error) {
	return e.gen.FillWithTestData(&entrances.GetEntranceResponse{}).(*entrances.GetEntranceResponse), nil
}

func (e *EntrancesServer) UpdateEntrance(ctx context.Context, request *entrances.UpdateEntranceRequest) (*entrances.UpdateEntranceResponse, error) {
	return e.gen.FillWithTestData(&entrances.UpdateEntranceResponse{}).(*entrances.UpdateEntranceResponse), nil
}

func (e *EntrancesServer) DeleteEntrance(ctx context.Context, request *entrances.DeleteEntranceRequest) (*entrances.DeleteEntranceResponse, error) {
	return e.gen.FillWithTestData(&entrances.DeleteEntranceResponse{}).(*entrances.DeleteEntranceResponse), nil
}

func (e *EntrancesServer) Lock(ctx context.Context, request *entrances.LockRequest) (*entrances.LockResponse, error) {
	return e.gen.FillWithTestData(&entrances.LockResponse{}).(*entrances.LockResponse), nil
}

func (e *EntrancesServer) Unlock(ctx context.Context, request *entrances.UnlockRequest) (*entrances.UnlockResponse, error) {
	return e.gen.FillWithTestData(&entrances.UnlockResponse{}).(*entrances.UnlockResponse), nil
}

func (e *EntrancesServer) FindNearestEntrance(ctx context.Context, request *entrances.FindNearestEntranceRequest) (*entrances.FindNearestEntranceResponse, error) {
	return e.gen.FillWithTestData(&entrances.FindNearestEntranceResponse{}).(*entrances.FindNearestEntranceResponse), nil
}

func (e *EntrancesServer) Enter(ctx context.Context, request *entrances.EnterRequest) (*entrances.EnterResponse, error) {
	return e.gen.FillWithTestData(&entrances.EnterResponse{}).(*entrances.EnterResponse), nil
}

func (e *EntrancesServer) Exit(ctx context.Context, request *entrances.ExitRequest) (*entrances.ExitResponse, error) {
	return e.gen.FillWithTestData(&entrances.ExitResponse{}).(*entrances.ExitResponse), nil
}

func (e *EntrancesServer) FetchAll(request *entrances.FetchAllEntrancesRequest, server entrances.MruVEntrancesService_FetchAllServer) error {
	panic("implement me")
}
