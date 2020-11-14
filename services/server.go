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
	panic("implement me")
}

func (s *ServerServer) GetRegisteredServers(ctx context.Context, request *server.GetRegisteredServersRequest) (*server.GetRegisteredServersResponse, error) {
	panic("implement me")
}

func (s *ServerServer) GetServerInfo(ctx context.Context, id *server.ServerID) (*server.ServerInfo, error) {
	panic("implement me")
}

func (s *ServerServer) UpdateServerStatus(ctx context.Context, request *server.UpdateServerStatusRequest) (*server.UpdateServerStatusResponse, error) {
	panic("implement me")
}

func (s *ServerServer) ServerEventsStream(request *server.ServerEventsStreamRequest, server server.MruVServerService_ServerEventsStreamServer) error {
	panic("implement me")
}
