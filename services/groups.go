package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/common"
	"github.com/MruV-RP/mruv-pb-go/groups"
)

type GroupsServer struct {
	gen generator.IGenerator
}

func NewGroupsServer(gen generator.IGenerator) *GroupsServer {
	return &GroupsServer{gen: gen}
}

func (g *GroupsServer) CreateGroup(ctx context.Context, request *groups.CreateGroupRequest) (*groups.CreateGroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetGroup(ctx context.Context, request *groups.GetGroupRequest) (*groups.GetGroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) UpdateGroup(ctx context.Context, request *groups.UpdateGroupRequest) (*groups.UpdateGroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) DeleteGroup(ctx context.Context, request *groups.DeleteGroupRequest) (*groups.DeleteGroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetGroups(ctx context.Context, request *groups.GetGroupsRequest) (*groups.GetGroupsResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) AssignOwner(ctx context.Context, request *groups.AssignOwnerRequest) (*groups.AssignOwnerResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetOwner(ctx context.Context, request *groups.GetOwnerRequest) (*groups.GetOwnerResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) AddMember(ctx context.Context, request *groups.AddMemberRequest) (*groups.AddMemberResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetMembers(ctx context.Context, request *groups.GetMembersRequest) (*groups.GetMembersResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) RemoveMember(ctx context.Context, request *groups.RemoveMemberRequest) (*groups.RemoveMemberResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) AddPermission(ctx context.Context, request *groups.AddPermissionRequest) (*groups.AddPermissionResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetPermissions(ctx context.Context, request *groups.GetPermissionsRequest) (*groups.GetPermissionsResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) RemovePermission(ctx context.Context, request *groups.RemovePermissionRequest) (*groups.RemovePermissionResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) AddSubgroup(ctx context.Context, request *groups.AddSubgroupRequest) (*groups.AddSubgroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetSubgroups(ctx context.Context, request *groups.GetSubgroupsRequest) (*groups.GetSubgroupsResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) RemoveSubgroup(ctx context.Context, request *groups.RemoveSubgroupRequest) (*groups.RemoveSubgroupResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) IsPermitted(ctx context.Context, request *groups.IsPermittedRequest) (*groups.IsPermittedResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	panic("implement me")
}

func (g *GroupsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	panic("implement me")
}
