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
	return g.gen.FillWithTestData(&gates.CreateGateResponse{}).(*gates.CreateGateResponse), nil
}

func (g *GatesServer) GetGate(ctx context.Context, request *gates.GetGateRequest) (*gates.GetGateResponse, error) {
	return g.gen.FillWithTestData(&gates.GetGateResponse{}).(*gates.GetGateResponse), nil
}

func (g *GatesServer) UpdateGate(ctx context.Context, request *gates.UpdateGateRequest) (*gates.UpdateGateResponse, error) {
	return g.gen.FillWithTestData(&gates.UpdateGateResponse{}).(*gates.UpdateGateResponse), nil
}

func (g *GatesServer) DeleteGate(ctx context.Context, request *gates.DeleteGateRequest) (*gates.DeleteGateResponse, error) {
	return g.gen.FillWithTestData(&gates.DeleteGateResponse{}).(*gates.DeleteGateResponse), nil
}

func (g *GatesServer) Lock(ctx context.Context, request *gates.LockRequest) (*gates.LockResponse, error) {
	return g.gen.FillWithTestData(&gates.LockResponse{}).(*gates.LockResponse), nil
}

func (g *GatesServer) Unlock(ctx context.Context, request *gates.UnlockRequest) (*gates.UnlockResponse, error) {
	return g.gen.FillWithTestData(&gates.UnlockResponse{}).(*gates.UnlockResponse), nil
}

func (g *GatesServer) Open(ctx context.Context, request *gates.OpenRequest) (*gates.OpenResponse, error) {
	return g.gen.FillWithTestData(&gates.OpenResponse{}).(*gates.OpenResponse), nil
}

func (g *GatesServer) Close(ctx context.Context, request *gates.CloseRequest) (*gates.CloseResponse, error) {
	return g.gen.FillWithTestData(&gates.CloseResponse{}).(*gates.CloseResponse), nil
}

func (g *GatesServer) FindNearestGate(ctx context.Context, request *gates.FindNearestGateRequest) (*gates.FindNearestGateResponse, error) {
	return g.gen.FillWithTestData(&gates.FindNearestGateResponse{}).(*gates.FindNearestGateResponse), nil
}

func (g *GatesServer) FetchAll(request *gates.FetchAllGatesRequest, server gates.MruVGatesService_FetchAllServer) error {
	panic("implement me")
}
