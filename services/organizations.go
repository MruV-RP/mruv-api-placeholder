package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/organizations"
)

type OrganizationsServer struct {
	gen generator.IGenerator
}

func NewOrganizationsServer(gen generator.IGenerator) *OrganizationsServer {
	return &OrganizationsServer{gen: gen}
}

func (o *OrganizationsServer) CreateOrganization(ctx context.Context, request *organizations.CreateOrganizationRequest) (*organizations.CreateOrganizationResponse, error) {
	panic("implement me")
}

func (o *OrganizationsServer) GetOrganization(ctx context.Context, request *organizations.GetOrganizationRequest) (*organizations.GetOrganizationResponse, error) {
	panic("implement me")
}

func (o *OrganizationsServer) UpdateOrganization(ctx context.Context, request *organizations.UpdateOrganizationRequest) (*organizations.UpdateOrganizationResponse, error) {
	panic("implement me")
}

func (o *OrganizationsServer) DeleteOrganization(ctx context.Context, request *organizations.DeleteOrganizationRequest) (*organizations.DeleteOrganizationResponse, error) {
	panic("implement me")
}

func (o *OrganizationsServer) AssignLeader(ctx context.Context, request *organizations.AssignLeaderRequest) (*organizations.AssignLeaderResponse, error) {
	panic("implement me")
}

func (o *OrganizationsServer) UnassignLeader(ctx context.Context, request *organizations.UnassignLeaderRequest) (*organizations.UnassignLeaderResponse, error) {
	panic("implement me")
}
