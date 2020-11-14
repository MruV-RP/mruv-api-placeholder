package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/plots"
)

type PlotsServer struct {
	gen generator.IGenerator
}

func NewPlotsServer(gen generator.IGenerator) *PlotsServer {
	return &PlotsServer{gen: gen}
}

func (p *PlotsServer) CreatePlot(ctx context.Context, request *plots.CreatePlotRequest) (*plots.CreatePlotResponse, error) {
	return p.gen.FillWithTestData(&plots.CreatePlotResponse{}).(*plots.CreatePlotResponse), nil
}

func (p *PlotsServer) GetPlot(ctx context.Context, request *plots.GetPlotRequest) (*plots.GetPlotResponse, error) {
	return p.gen.FillWithTestData(&plots.GetPlotResponse{}).(*plots.GetPlotResponse), nil
}

func (p *PlotsServer) UpdatePlot(ctx context.Context, request *plots.UpdatePlotRequest) (*plots.UpdatePlotResponse, error) {
	return p.gen.FillWithTestData(&plots.UpdatePlotResponse{}).(*plots.UpdatePlotResponse), nil
}

func (p *PlotsServer) DeletePlot(ctx context.Context, request *plots.DeletePlotRequest) (*plots.DeletePlotResponse, error) {
	return p.gen.FillWithTestData(&plots.DeletePlotResponse{}).(*plots.DeletePlotResponse), nil
}
