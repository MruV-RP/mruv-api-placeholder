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
	return o.gen.FillWithTestData(&objects.CreateObjectResponse{}).(*objects.CreateObjectResponse), nil
}

func (o *ObjectsServer) GetObject(ctx context.Context, request *objects.GetObjectRequest) (*objects.GetObjectResponse, error) {
	return o.gen.FillWithTestData(&objects.GetObjectResponse{}).(*objects.GetObjectResponse), nil
}

func (o *ObjectsServer) UpdateObject(ctx context.Context, request *objects.UpdateObjectRequest) (*objects.UpdateObjectResponse, error) {
	return o.gen.FillWithTestData(&objects.UpdateObjectResponse{}).(*objects.UpdateObjectResponse), nil
}

func (o *ObjectsServer) DeleteObject(ctx context.Context, request *objects.DeleteObjectRequest) (*objects.DeleteObjectResponse, error) {
	return o.gen.FillWithTestData(&objects.DeleteObjectResponse{}).(*objects.DeleteObjectResponse), nil
}

func (o *ObjectsServer) AddObjectMaterial(ctx context.Context, request *objects.AddObjectMaterialRequest) (*objects.AddObjectMaterialResponse, error) {
	return o.gen.FillWithTestData(&objects.AddObjectMaterialResponse{}).(*objects.AddObjectMaterialResponse), nil
}

func (o *ObjectsServer) GetObjectMaterials(ctx context.Context, request *objects.GetObjectMaterialsRequest) (*objects.GetObjectMaterialsResponse, error) {
	return o.gen.FillWithTestData(&objects.GetObjectMaterialsResponse{}).(*objects.GetObjectMaterialsResponse), nil
}

func (o *ObjectsServer) DeleteObjectMaterial(ctx context.Context, request *objects.DeleteObjectMaterialRequest) (*objects.DeleteObjectMaterialResponse, error) {
	return o.gen.FillWithTestData(&objects.DeleteObjectMaterialResponse{}).(*objects.DeleteObjectMaterialResponse), nil
}

func (o *ObjectsServer) AddObjectMaterialText(ctx context.Context, request *objects.AddObjectMaterialTextRequest) (*objects.AddObjectMaterialTextResponse, error) {
	return o.gen.FillWithTestData(&objects.AddObjectMaterialTextResponse{}).(*objects.AddObjectMaterialTextResponse), nil
}

func (o *ObjectsServer) GetObjectMaterialTexts(ctx context.Context, request *objects.GetObjectMaterialTextsRequest) (*objects.GetObjectMaterialTextsResponse, error) {
	return o.gen.FillWithTestData(&objects.GetObjectMaterialTextsResponse{}).(*objects.GetObjectMaterialTextsResponse), nil
}

func (o *ObjectsServer) DeleteObjectMaterialText(ctx context.Context, request *objects.DeleteObjectMaterialTextRequest) (*objects.DeleteObjectMaterialTextResponse, error) {
	return o.gen.FillWithTestData(&objects.DeleteObjectMaterialTextResponse{}).(*objects.DeleteObjectMaterialTextResponse), nil
}

func (o *ObjectsServer) AddRemoveBuilding(ctx context.Context, request *objects.AddRemoveBuildingRequest) (*objects.AddRemoveBuildingResponse, error) {
	return o.gen.FillWithTestData(&objects.AddRemoveBuildingResponse{}).(*objects.AddRemoveBuildingResponse), nil
}

func (o *ObjectsServer) GetRemovedBuildings(ctx context.Context, request *objects.GetRemovedBuildingsRequest) (*objects.GetRemovedBuildingsResponse, error) {
	return o.gen.FillWithTestData(&objects.GetRemovedBuildingsResponse{}).(*objects.GetRemovedBuildingsResponse), nil
}

func (o *ObjectsServer) DeleteRemoveBuilding(ctx context.Context, request *objects.DeleteRemoveBuildingRequest) (*objects.DeleteRemoveBuildingResponse, error) {
	return o.gen.FillWithTestData(&objects.DeleteRemoveBuildingResponse{}).(*objects.DeleteRemoveBuildingResponse), nil
}

func (o *ObjectsServer) FetchAllObjects(request *objects.FetchAllObjectsRequest, server objects.MruVObjectsService_FetchAllObjectsServer) error {
	panic("implement me")
}
