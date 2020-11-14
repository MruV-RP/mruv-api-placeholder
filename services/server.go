package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/server"
)

type ServerServer struct {
	gen generator.IGenerator
}

func NewServerServer(gen generator.IGenerator) *ServerServer {
	return &ServerServer{gen: gen}
}

func (s *ServerServer) RegisterServer(ctx context.Context, info *server.ServerInfo) (*server.ServerID, error) {
	return s.gen.FillWithTestData(&server.ServerID{}).(*server.ServerID), nil
}

func (s *ServerServer) GetRegisteredServers(ctx context.Context, request *server.GetRegisteredServersRequest) (*server.GetRegisteredServersResponse, error) {
	return s.gen.FillWithTestData(&server.GetRegisteredServersResponse{}).(*server.GetRegisteredServersResponse), nil
}

func (s *ServerServer) GetServerInfo(ctx context.Context, id *server.ServerID) (*server.ServerInfo, error) {
	return s.gen.FillWithTestData(&server.ServerInfo{}).(*server.ServerInfo), nil
}

func (s *ServerServer) UpdateServerStatus(ctx context.Context, request *server.UpdateServerStatusRequest) (*server.UpdateServerStatusResponse, error) {
	return s.gen.FillWithTestData(&server.UpdateServerStatusResponse{}).(*server.UpdateServerStatusResponse), nil
}

func (s *ServerServer) ServerEventsStream(request *server.ServerEventsStreamRequest, server server.MruVServerService_ServerEventsStreamServer) error {
	panic("implement me")
}
