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
	return g.gen.FillWithTestData(&groups.CreateGroupResponse{}).(*groups.CreateGroupResponse), nil
}

func (g *GroupsServer) GetGroup(ctx context.Context, request *groups.GetGroupRequest) (*groups.GetGroupResponse, error) {
	return g.gen.FillWithTestData(&groups.GetGroupResponse{}).(*groups.GetGroupResponse), nil
}

func (g *GroupsServer) UpdateGroup(ctx context.Context, request *groups.UpdateGroupRequest) (*groups.UpdateGroupResponse, error) {
	return g.gen.FillWithTestData(&groups.UpdateGroupResponse{}).(*groups.UpdateGroupResponse), nil
}

func (g *GroupsServer) DeleteGroup(ctx context.Context, request *groups.DeleteGroupRequest) (*groups.DeleteGroupResponse, error) {
	return g.gen.FillWithTestData(&groups.DeleteGroupResponse{}).(*groups.DeleteGroupResponse), nil
}

func (g *GroupsServer) GetGroups(ctx context.Context, request *groups.GetGroupsRequest) (*groups.GetGroupsResponse, error) {
	return g.gen.FillWithTestData(&groups.GetGroupsResponse{}).(*groups.GetGroupsResponse), nil
}

func (g *GroupsServer) AssignOwner(ctx context.Context, request *groups.AssignOwnerRequest) (*groups.AssignOwnerResponse, error) {
	return g.gen.FillWithTestData(&groups.AssignOwnerResponse{}).(*groups.AssignOwnerResponse), nil
}

func (g *GroupsServer) GetOwner(ctx context.Context, request *groups.GetOwnerRequest) (*groups.GetOwnerResponse, error) {
	return g.gen.FillWithTestData(&groups.GetOwnerResponse{}).(*groups.GetOwnerResponse), nil
}

func (g *GroupsServer) AddMember(ctx context.Context, request *groups.AddMemberRequest) (*groups.AddMemberResponse, error) {
	return g.gen.FillWithTestData(&groups.AddMemberResponse{}).(*groups.AddMemberResponse), nil
}

func (g *GroupsServer) GetMembers(ctx context.Context, request *groups.GetMembersRequest) (*groups.GetMembersResponse, error) {
	return g.gen.FillWithTestData(&groups.GetMembersResponse{}).(*groups.GetMembersResponse), nil
}

func (g *GroupsServer) RemoveMember(ctx context.Context, request *groups.RemoveMemberRequest) (*groups.RemoveMemberResponse, error) {
	return g.gen.FillWithTestData(&groups.RemoveMemberResponse{}).(*groups.RemoveMemberResponse), nil
}

func (g *GroupsServer) AddPermission(ctx context.Context, request *groups.AddPermissionRequest) (*groups.AddPermissionResponse, error) {
	return g.gen.FillWithTestData(&groups.AddPermissionResponse{}).(*groups.AddPermissionResponse), nil
}

func (g *GroupsServer) GetPermissions(ctx context.Context, request *groups.GetPermissionsRequest) (*groups.GetPermissionsResponse, error) {
	return g.gen.FillWithTestData(&groups.GetPermissionsResponse{}).(*groups.GetPermissionsResponse), nil
}

func (g *GroupsServer) RemovePermission(ctx context.Context, request *groups.RemovePermissionRequest) (*groups.RemovePermissionResponse, error) {
	return g.gen.FillWithTestData(&groups.RemovePermissionResponse{}).(*groups.RemovePermissionResponse), nil
}

func (g *GroupsServer) AddSubgroup(ctx context.Context, request *groups.AddSubgroupRequest) (*groups.AddSubgroupResponse, error) {
	return g.gen.FillWithTestData(&groups.AddSubgroupResponse{}).(*groups.AddSubgroupResponse), nil
}

func (g *GroupsServer) GetSubgroups(ctx context.Context, request *groups.GetSubgroupsRequest) (*groups.GetSubgroupsResponse, error) {
	return g.gen.FillWithTestData(&groups.GetSubgroupsResponse{}).(*groups.GetSubgroupsResponse), nil
}

func (g *GroupsServer) RemoveSubgroup(ctx context.Context, request *groups.RemoveSubgroupRequest) (*groups.RemoveSubgroupResponse, error) {
	return g.gen.FillWithTestData(&groups.RemoveSubgroupResponse{}).(*groups.RemoveSubgroupResponse), nil
}

func (g *GroupsServer) IsPermitted(ctx context.Context, request *groups.IsPermittedRequest) (*groups.IsPermittedResponse, error) {
	return g.gen.FillWithTestData(&groups.IsPermittedResponse{}).(*groups.IsPermittedResponse), nil
}

func (g *GroupsServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	return g.gen.FillWithTestData(&common.ServiceStatusResponse{}).(*common.ServiceStatusResponse), nil
}

func (g *GroupsServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	return g.gen.FillWithTestData(&common.VersionResponse{}).(*common.VersionResponse), nil
}
