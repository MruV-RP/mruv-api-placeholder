package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/objects"
)

type ObjectsServer struct {
	gen generator.IGenerator
}

func NewObjectsServer(gen generator.IGenerator) *ObjectsServer {
	return &ObjectsServer{gen: gen}
}

func (o *ObjectsServer) CreateObject(ctx context.Context, request *objects.CreateObjectRequest) (*objects.CreateObjectResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) GetObject(ctx context.Context, request *objects.GetObjectRequest) (*objects.GetObjectResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) UpdateObject(ctx context.Context, request *objects.UpdateObjectRequest) (*objects.UpdateObjectResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) DeleteObject(ctx context.Context, request *objects.DeleteObjectRequest) (*objects.DeleteObjectResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) AddObjectMaterial(ctx context.Context, request *objects.AddObjectMaterialRequest) (*objects.AddObjectMaterialResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) GetObjectMaterials(ctx context.Context, request *objects.GetObjectMaterialsRequest) (*objects.GetObjectMaterialsResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) DeleteObjectMaterial(ctx context.Context, request *objects.DeleteObjectMaterialRequest) (*objects.DeleteObjectMaterialResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) AddObjectMaterialText(ctx context.Context, request *objects.AddObjectMaterialTextRequest) (*objects.AddObjectMaterialTextResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) GetObjectMaterialTexts(ctx context.Context, request *objects.GetObjectMaterialTextsRequest) (*objects.GetObjectMaterialTextsResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) DeleteObjectMaterialText(ctx context.Context, request *objects.DeleteObjectMaterialTextRequest) (*objects.DeleteObjectMaterialTextResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) AddRemoveBuilding(ctx context.Context, request *objects.AddRemoveBuildingRequest) (*objects.AddRemoveBuildingResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) GetRemovedBuildings(ctx context.Context, request *objects.GetRemovedBuildingsRequest) (*objects.GetRemovedBuildingsResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) DeleteRemoveBuilding(ctx context.Context, request *objects.DeleteRemoveBuildingRequest) (*objects.DeleteRemoveBuildingResponse, error) {
	panic("implement me")
}

func (o *ObjectsServer) FetchAllObjects(request *objects.FetchAllObjectsRequest, server objects.MruVObjectsService_FetchAllObjectsServer) error {
	panic("implement me")
}
