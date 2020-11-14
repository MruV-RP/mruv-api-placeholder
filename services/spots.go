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
	return s.gen.FillWithTestData(&spots.CreateSpotResponse{}).(*spots.CreateSpotResponse), nil
}

func (s *SpotsServer) GetSpot(ctx context.Context, request *spots.GetSpotRequest) (*spots.GetSpotResponse, error) {
	return s.gen.FillWithTestData(&spots.GetSpotResponse{}).(*spots.GetSpotResponse), nil
}

func (s *SpotsServer) UpdateSpot(ctx context.Context, request *spots.UpdateSpotRequest) (*spots.UpdateSpotResponse, error) {
	return s.gen.FillWithTestData(&spots.UpdateSpotResponse{}).(*spots.UpdateSpotResponse), nil
}

func (s *SpotsServer) DeleteSpot(ctx context.Context, request *spots.DeleteSpotRequest) (*spots.DeleteSpotResponse, error) {
	return s.gen.FillWithTestData(&spots.DeleteSpotResponse{}).(*spots.DeleteSpotResponse), nil
}

func (s *SpotsServer) FetchAll(request *spots.FetchAllSpotsRequest, server spots.MruVSpotsService_FetchAllServer) error {
	panic("implement me")
}
