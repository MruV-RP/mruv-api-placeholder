package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/common"
	"github.com/MruV-RP/mruv-pb-go/items"
)

type ItemsServer struct {
}

func (i *ItemsServer) CreateItem(ctx context.Context, item *items.Item) (*items.ItemID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetItem(ctx context.Context, id *items.ItemID) (*items.Item, error) {
	panic("implement me")
}

func (i *ItemsServer) DeleteItem(ctx context.Context, id *items.ItemID) (*items.ItemID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetItems(ctx context.Context, request *items.GetItemsRequest) (*items.GetItemsResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) CreateItemType(ctx context.Context, itemType *items.ItemType) (*items.ItemTypeID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetItemType(ctx context.Context, id *items.ItemTypeID) (*items.ItemType, error) {
	panic("implement me")
}

func (i *ItemsServer) DeleteItemType(ctx context.Context, id *items.ItemTypeID) (*items.ItemTypeID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetItemTypes(ctx context.Context, request *items.GetItemTypesRequest) (*items.GetItemTypesResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) CreateContainer(ctx context.Context, container *items.Container) (*items.ContainerID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetContainer(ctx context.Context, id *items.ContainerID) (*items.Container, error) {
	panic("implement me")
}

func (i *ItemsServer) DeleteContainer(ctx context.Context, id *items.ContainerID) (*items.ContainerID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetContainers(ctx context.Context, request *items.GetContainersRequest) (*items.GetContainersResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) CreateContainerType(ctx context.Context, containerType *items.ContainerType) (*items.ContainerTypeID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetContainerType(ctx context.Context, id *items.ContainerTypeID) (*items.ContainerType, error) {
	panic("implement me")
}

func (i *ItemsServer) DeleteContainerType(ctx context.Context, id *items.ContainerTypeID) (*items.ContainerTypeID, error) {
	panic("implement me")
}

func (i *ItemsServer) GetContainerTypes(ctx context.Context, request *items.GetContainerTypesRequest) (*items.GetContainerTypesResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) GetContainerItems(ctx context.Context, request *items.GetContainerItemsRequest) (*items.GetContainerItemsResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) PullItem(ctx context.Context, request *items.PullItemRequest) (*items.Item, error) {
	panic("implement me")
}

func (i *ItemsServer) PutItem(ctx context.Context, request *items.PutItemRequest) (*items.PutItemResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) SortItems(ctx context.Context, request *items.SortItemsRequest) (*items.SortItemsResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) GetNearestItems(ctx context.Context, request *items.GetNearestItemsRequest) (*items.GetNearestItemsResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) UseItem(ctx context.Context, request *items.UseItemRequest) (*items.UseItemResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	panic("implement me")
}

func (i *ItemsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	panic("implement me")
}

func NewItemsServer() *ItemsServer {
	return &ItemsServer{}
}
