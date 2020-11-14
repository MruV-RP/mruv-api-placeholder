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
	return o.gen.FillWithTestData(&offers.CreateOfferResponse{}).(*offers.CreateOfferResponse), nil
}

func (o *OffersServer) GetOffer(ctx context.Context, request *offers.GetOfferRequest) (*offers.GetOfferResponse, error) {
	return o.gen.FillWithTestData(&offers.GetOfferResponse{}).(*offers.GetOfferResponse), nil
}

func (o *OffersServer) UpdateOffer(ctx context.Context, request *offers.UpdateOfferRequest) (*offers.UpdateOfferResponse, error) {
	return o.gen.FillWithTestData(&offers.UpdateOfferResponse{}).(*offers.UpdateOfferResponse), nil
}

func (o *OffersServer) DeleteOffer(ctx context.Context, request *offers.DeleteOfferRequest) (*offers.DeleteOfferResponse, error) {
	return o.gen.FillWithTestData(&offers.DeleteOfferResponse{}).(*offers.DeleteOfferResponse), nil
}

func (o *OffersServer) AcceptOffer(ctx context.Context, request *offers.AcceptOfferRequest) (*offers.AcceptOfferResponse, error) {
	return o.gen.FillWithTestData(&offers.AcceptOfferResponse{}).(*offers.AcceptOfferResponse), nil
}
