package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/spots"
)

type SpotsServer struct {
	gen generator.IGenerator
}

func NewSpotsServer(gen generator.IGenerator) *SpotsServer {
	return &SpotsServer{gen: gen}
}

func (s *SpotsServer) CreateSpot(ctx context.Context, request *spots.CreateSpotRequest) (*spots.CreateSpotResponse, error) {
	panic("implement me")
}

func (s *SpotsServer) GetSpot(ctx context.Context, request *spots.GetSpotRequest) (*spots.GetSpotResponse, error) {
	panic("implement me")
}

func (s *SpotsServer) UpdateSpot(ctx context.Context, request *spots.UpdateSpotRequest) (*spots.UpdateSpotResponse, error) {
	panic("implement me")
}

func (s *SpotsServer) DeleteSpot(ctx context.Context, request *spots.DeleteSpotRequest) (*spots.DeleteSpotResponse, error) {
	panic("implement me")
}

func (s *SpotsServer) FetchAll(request *spots.FetchAllSpotsRequest, server spots.MruVSpotsService_FetchAllServer) error {
	panic("implement me")
}
