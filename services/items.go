package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/common"
	"github.com/MruV-RP/mruv-pb-go/items"
)

type ItemsServer struct {
	gen generator.IGenerator
}

func NewItemsServer(gen generator.IGenerator) *ItemsServer {
	return &ItemsServer{gen: gen}
}

func (i *ItemsServer) CreateItem(ctx context.Context, item *items.Item) (*items.ItemID, error) {
	return i.gen.FillWithTestData(&items.ItemID{}).(*items.ItemID), nil
}

func (i *ItemsServer) GetItem(ctx context.Context, id *items.ItemID) (*items.Item, error) {
	return i.gen.FillWithTestData(&items.Item{}).(*items.Item), nil
}

func (i *ItemsServer) DeleteItem(ctx context.Context, id *items.ItemID) (*items.ItemID, error) {
	return i.gen.FillWithTestData(&items.ItemID{}).(*items.ItemID), nil
}

func (i *ItemsServer) GetItems(ctx context.Context, request *items.GetItemsRequest) (*items.GetItemsResponse, error) {
	return i.gen.FillWithTestData(&items.GetItemsResponse{}).(*items.GetItemsResponse), nil
}

func (i *ItemsServer) CreateItemType(ctx context.Context, itemType *items.ItemType) (*items.ItemTypeID, error) {
	return i.gen.FillWithTestData(&items.ItemTypeID{}).(*items.ItemTypeID), nil
}

func (i *ItemsServer) GetItemType(ctx context.Context, id *items.ItemTypeID) (*items.ItemType, error) {
	return i.gen.FillWithTestData(&items.ItemType{}).(*items.ItemType), nil
}

func (i *ItemsServer) DeleteItemType(ctx context.Context, id *items.ItemTypeID) (*items.ItemTypeID, error) {
	return i.gen.FillWithTestData(&items.ItemTypeID{}).(*items.ItemTypeID), nil
}

func (i *ItemsServer) GetItemTypes(ctx context.Context, request *items.GetItemTypesRequest) (*items.GetItemTypesResponse, error) {
	return i.gen.FillWithTestData(&items.GetItemTypesResponse{}).(*items.GetItemTypesResponse), nil
}

func (i *ItemsServer) CreateContainer(ctx context.Context, container *items.Container) (*items.ContainerID, error) {
	return i.gen.FillWithTestData(&items.ContainerID{}).(*items.ContainerID), nil
}

func (i *ItemsServer) GetContainer(ctx context.Context, id *items.ContainerID) (*items.Container, error) {
	return i.gen.FillWithTestData(&items.Container{}).(*items.Container), nil
}

func (i *ItemsServer) DeleteContainer(ctx context.Context, id *items.ContainerID) (*items.ContainerID, error) {
	return i.gen.FillWithTestData(&items.ContainerID{}).(*items.ContainerID), nil
}

func (i *ItemsServer) GetContainers(ctx context.Context, request *items.GetContainersRequest) (*items.GetContainersResponse, error) {
	return i.gen.FillWithTestData(&items.GetContainersResponse{}).(*items.GetContainersResponse), nil
}

func (i *ItemsServer) CreateContainerType(ctx context.Context, containerType *items.ContainerType) (*items.ContainerTypeID, error) {
	return i.gen.FillWithTestData(&items.ContainerTypeID{}).(*items.ContainerTypeID), nil
}

func (i *ItemsServer) GetContainerType(ctx context.Context, id *items.ContainerTypeID) (*items.ContainerType, error) {
	return i.gen.FillWithTestData(&items.ContainerType{}).(*items.ContainerType), nil
}

func (i *ItemsServer) DeleteContainerType(ctx context.Context, id *items.ContainerTypeID) (*items.ContainerTypeID, error) {
	return i.gen.FillWithTestData(&items.ContainerTypeID{}).(*items.ContainerTypeID), nil
}

func (i *ItemsServer) GetContainerTypes(ctx context.Context, request *items.GetContainerTypesRequest) (*items.GetContainerTypesResponse, error) {
	return i.gen.FillWithTestData(&items.GetContainerTypesResponse{}).(*items.GetContainerTypesResponse), nil
}

func (i *ItemsServer) GetContainerItems(ctx context.Context, request *items.GetContainerItemsRequest) (*items.GetContainerItemsResponse, error) {
	return i.gen.FillWithTestData(&items.GetContainerItemsResponse{}).(*items.GetContainerItemsResponse), nil
}

func (i *ItemsServer) PullItem(ctx context.Context, request *items.PullItemRequest) (*items.Item, error) {
	return i.gen.FillWithTestData(&items.Item{}).(*items.Item), nil
}

func (i *ItemsServer) PutItem(ctx context.Context, request *items.PutItemRequest) (*items.PutItemResponse, error) {
	return i.gen.FillWithTestData(&items.PutItemResponse{}).(*items.PutItemResponse), nil
}

func (i *ItemsServer) SortItems(ctx context.Context, request *items.SortItemsRequest) (*items.SortItemsResponse, error) {
	return i.gen.FillWithTestData(&items.SortItemsResponse{}).(*items.SortItemsResponse), nil
}

func (i *ItemsServer) GetNearestItems(ctx context.Context, request *items.GetNearestItemsRequest) (*items.GetNearestItemsResponse, error) {
	return i.gen.FillWithTestData(&items.GetNearestItemsResponse{}).(*items.GetNearestItemsResponse), nil
}

func (i *ItemsServer) UseItem(ctx context.Context, request *items.UseItemRequest) (*items.UseItemResponse, error) {
	return i.gen.FillWithTestData(&items.UseItemResponse{}).(*items.UseItemResponse), nil
}

func (i *ItemsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	return i.gen.FillWithTestData(&common.ServiceStatusResponse{}).(*common.ServiceStatusResponse), nil
}

func (i *ItemsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	return i.gen.FillWithTestData(&common.VersionResponse{}).(*common.VersionResponse), nil
}
