package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/offers"
)

type OffersServer struct {
	gen generator.IGenerator
}

func NewOffersServer(gen generator.IGenerator) *OffersServer {
	return &OffersServer{gen: gen}
}

func (o *OffersServer) CreateOffer(ctx context.Context, request *offers.CreateOfferRequest) (*offers.CreateOfferResponse, error) {
	panic("implement me")
}

func (o *OffersServer) GetOffer(ctx context.Context, request *offers.GetOfferRequest) (*offers.GetOfferResponse, error) {
	panic("implement me")
}

func (o *OffersServer) UpdateOffer(ctx context.Context, request *offers.UpdateOfferRequest) (*offers.UpdateOfferResponse, error) {
	panic("implement me")
}

func (o *OffersServer) DeleteOffer(ctx context.Context, request *offers.DeleteOfferRequest) (*offers.DeleteOfferResponse, error) {
	panic("implement me")
}

func (o *OffersServer) AcceptOffer(ctx context.Context, request *offers.AcceptOfferRequest) (*offers.AcceptOfferResponse, error) {
	panic("implement me")
}
