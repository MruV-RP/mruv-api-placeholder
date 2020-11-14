package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/gates"
)

type GatesServer struct {
	gen generator.IGenerator
}

func NewGatesServer(gen generator.IGenerator) *GatesServer {
	return &GatesServer{gen: gen}
}

func (g *GatesServer) CreateGate(ctx context.Context, request *gates.CreateGateRequest) (*gates.CreateGateResponse, error) {
	panic("implement me")
}

func (g *GatesServer) GetGate(ctx context.Context, request *gates.GetGateRequest) (*gates.GetGateResponse, error) {
	panic("implement me")
}

func (g *GatesServer) UpdateGate(ctx context.Context, request *gates.UpdateGateRequest) (*gates.UpdateGateResponse, error) {
	panic("implement me")
}

func (g *GatesServer) DeleteGate(ctx context.Context, request *gates.DeleteGateRequest) (*gates.DeleteGateResponse, error) {
	panic("implement me")
}

func (g *GatesServer) Lock(ctx context.Context, request *gates.LockRequest) (*gates.LockResponse, error) {
	panic("implement me")
}

func (g *GatesServer) Unlock(ctx context.Context, request *gates.UnlockRequest) (*gates.UnlockResponse, error) {
	panic("implement me")
}

func (g *GatesServer) Open(ctx context.Context, request *gates.OpenRequest) (*gates.OpenResponse, error) {
	panic("implement me")
}

func (g *GatesServer) Close(ctx context.Context, request *gates.CloseRequest) (*gates.CloseResponse, error) {
	panic("implement me")
}

func (g *GatesServer) FindNearestGate(ctx context.Context, request *gates.FindNearestGateRequest) (*gates.FindNearestGateResponse, error) {
	panic("implement me")
}

func (g *GatesServer) FetchAll(request *gates.FetchAllGatesRequest, server gates.MruVGatesService_FetchAllServer) error {
	panic("implement me")
}
