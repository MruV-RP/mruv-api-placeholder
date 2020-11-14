package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/common"
	"github.com/MruV-RP/mruv-pb-go/devtools"
)

type DevtoolsServer struct {
	gen generator.IGenerator
}

func NewDevtoolsServer(gen generator.IGenerator) *DevtoolsServer {
	return &DevtoolsServer{gen: gen}
}

func (d *DevtoolsServer) GetPositions(ctx context.Context, request *devtools.GetPositionsRequest) (*devtools.GetPositionsResponse, error) {
	return d.gen.FillWithTestData(&devtools.GetPositionsResponse{}).(*devtools.GetPositionsResponse), nil
}

func (d *DevtoolsServer) SavePosition(ctx context.Context, request *devtools.SavePositionRequest) (*devtools.SavePositionResponse, error) {
	return d.gen.FillWithTestData(&devtools.SavePositionResponse{}).(*devtools.SavePositionResponse), nil
}

func (d *DevtoolsServer) DeletePosition(ctx context.Context, request *devtools.DeletePositionRequest) (*devtools.DeletePositionResponse, error) {
	return d.gen.FillWithTestData(&devtools.DeletePositionResponse{}).(*devtools.DeletePositionResponse), nil
}

func (d *DevtoolsServer) GetOutfits(ctx context.Context, request *devtools.GetOutfitsRequest) (*devtools.GetOutfitsResponse, error) {
	return d.gen.FillWithTestData(&devtools.GetOutfitsResponse{}).(*devtools.GetOutfitsResponse), nil
}

func (d *DevtoolsServer) SaveOutfit(ctx context.Context, request *devtools.SaveOutfitRequest) (*devtools.SaveOutfitResponse, error) {
	return d.gen.FillWithTestData(&devtools.SaveOutfitResponse{}).(*devtools.SaveOutfitResponse), nil
}

func (d *DevtoolsServer) GetAnimations(ctx context.Context, request *devtools.GetAnimationsRequest) (*devtools.GetAnimationsResponse, error) {
	return d.gen.FillWithTestData(&devtools.GetAnimationsResponse{}).(*devtools.GetAnimationsResponse), nil
}

func (d *DevtoolsServer) GetAnimation(ctx context.Context, request *devtools.GetAnimationRequest) (*devtools.GetAnimationResponse, error) {
	return d.gen.FillWithTestData(&devtools.GetAnimationResponse{}).(*devtools.GetAnimationResponse), nil
}

func (d *DevtoolsServer) SaveAnimation(ctx context.Context, request *devtools.SaveAnimationRequest) (*devtools.SaveAnimationResponse, error) {
	return d.gen.FillWithTestData(&devtools.SaveAnimationResponse{}).(*devtools.SaveAnimationResponse), nil
}

func (d *DevtoolsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	return d.gen.FillWithTestData(&common.ServiceStatusResponse{}).(*common.ServiceStatusResponse), nil
}

func (d *DevtoolsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	return d.gen.FillWithTestData(&common.VersionResponse{}).(*common.VersionResponse), nil
}
