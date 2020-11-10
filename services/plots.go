package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/plots"
)

type PlotsServer struct {
}

func (p *PlotsServer) CreatePlot(ctx context.Context, request *plots.CreatePlotRequest) (*plots.CreatePlotResponse, error) {
	panic("implement me")
}

func (p *PlotsServer) GetPlot(ctx context.Context, request *plots.GetPlotRequest) (*plots.GetPlotResponse, error) {
	panic("implement me")
}

func (p *PlotsServer) UpdatePlot(ctx context.Context, request *plots.UpdatePlotRequest) (*plots.UpdatePlotResponse, error) {
	panic("implement me")
}

func (p *PlotsServer) DeletePlot(ctx context.Context, request *plots.DeletePlotRequest) (*plots.DeletePlotResponse, error) {
	panic("implement me")
}

func NewPlotsServer() *PlotsServer {
	return &PlotsServer{}
}
