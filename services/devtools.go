package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/common"
	"github.com/MruV-RP/mruv-pb-go/devtools"
)

type DevtoolsServer struct {
}

func (d *DevtoolsServer) GetPositions(ctx context.Context, request *devtools.GetPositionsRequest) (*devtools.GetPositionsResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) SavePosition(ctx context.Context, request *devtools.SavePositionRequest) (*devtools.SavePositionResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) DeletePosition(ctx context.Context, request *devtools.DeletePositionRequest) (*devtools.DeletePositionResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) GetOutfits(ctx context.Context, request *devtools.GetOutfitsRequest) (*devtools.GetOutfitsResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) SaveOutfit(ctx context.Context, request *devtools.SaveOutfitRequest) (*devtools.SaveOutfitResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) GetAnimations(ctx context.Context, request *devtools.GetAnimationsRequest) (*devtools.GetAnimationsResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) GetAnimation(ctx context.Context, request *devtools.GetAnimationRequest) (*devtools.GetAnimationResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) SaveAnimation(ctx context.Context, request *devtools.SaveAnimationRequest) (*devtools.SaveAnimationResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	panic("implement me")
}

func (d *DevtoolsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	panic("implement me")
}

func NewDevtoolsServer() *DevtoolsServer {
	return &DevtoolsServer{}
}
