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
	return o.gen.FillWithTestData(&organizations.CreateOrganizationResponse{}).(*organizations.CreateOrganizationResponse), nil
}

func (o *OrganizationsServer) GetOrganization(ctx context.Context, request *organizations.GetOrganizationRequest) (*organizations.GetOrganizationResponse, error) {
	return o.gen.FillWithTestData(&organizations.GetOrganizationResponse{}).(*organizations.GetOrganizationResponse), nil
}

func (o *OrganizationsServer) UpdateOrganization(ctx context.Context, request *organizations.UpdateOrganizationRequest) (*organizations.UpdateOrganizationResponse, error) {
	return o.gen.FillWithTestData(&organizations.UpdateOrganizationResponse{}).(*organizations.UpdateOrganizationResponse), nil
}

func (o *OrganizationsServer) DeleteOrganization(ctx context.Context, request *organizations.DeleteOrganizationRequest) (*organizations.DeleteOrganizationResponse, error) {
	return o.gen.FillWithTestData(&organizations.DeleteOrganizationResponse{}).(*organizations.DeleteOrganizationResponse), nil
}

func (o *OrganizationsServer) AssignLeader(ctx context.Context, request *organizations.AssignLeaderRequest) (*organizations.AssignLeaderResponse, error) {
	return o.gen.FillWithTestData(&organizations.AssignLeaderResponse{}).(*organizations.AssignLeaderResponse), nil
}

func (o *OrganizationsServer) UnassignLeader(ctx context.Context, request *organizations.UnassignLeaderRequest) (*organizations.UnassignLeaderResponse, error) {
	return o.gen.FillWithTestData(&organizations.UnassignLeaderResponse{}).(*organizations.UnassignLeaderResponse), nil
}
