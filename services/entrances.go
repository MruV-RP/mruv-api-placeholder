package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/entrances"
)

type EntrancesServer struct {
}

func (e *EntrancesServer) CreateEntrance(ctx context.Context, request *entrances.CreateEntranceRequest) (*entrances.CreateEntranceResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) GetEntrance(ctx context.Context, request *entrances.GetEntranceRequest) (*entrances.GetEntranceResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) UpdateEntrance(ctx context.Context, request *entrances.UpdateEntranceRequest) (*entrances.UpdateEntranceResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) DeleteEntrance(ctx context.Context, request *entrances.DeleteEntranceRequest) (*entrances.DeleteEntranceResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) Lock(ctx context.Context, request *entrances.LockRequest) (*entrances.LockResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) Unlock(ctx context.Context, request *entrances.UnlockRequest) (*entrances.UnlockResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) FindNearestEntrance(ctx context.Context, request *entrances.FindNearestEntranceRequest) (*entrances.FindNearestEntranceResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) Enter(ctx context.Context, request *entrances.EnterRequest) (*entrances.EnterResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) Exit(ctx context.Context, request *entrances.ExitRequest) (*entrances.ExitResponse, error) {
	panic("implement me")
}

func (e *EntrancesServer) FetchAll(request *entrances.FetchAllEntrancesRequest, server entrances.MruVEntrancesService_FetchAllServer) error {
	panic("implement me")
}

func NewEntrancesServer() *EntrancesServer {
	return &EntrancesServer{}
}
