package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/punishments"
)

type PunishmentsServer struct {
	gen generator.IGenerator
}

func NewPunishmentsServer(gen generator.IGenerator) *PunishmentsServer {
	return &PunishmentsServer{gen: gen}
}

func (p *PunishmentsServer) Ban(ctx context.Context, request *punishments.BanRequest) (*punishments.BanResponse, error) {
	return p.gen.FillWithTestData(&punishments.BanResponse{}).(*punishments.BanResponse), nil
}

func (p *PunishmentsServer) Block(ctx context.Context, request *punishments.BlockRequest) (*punishments.BlockResponse, error) {
	return p.gen.FillWithTestData(&punishments.BlockResponse{}).(*punishments.BlockResponse), nil
}

func (p *PunishmentsServer) Warn(ctx context.Context, request *punishments.WarnRequest) (*punishments.WarnResponse, error) {
	return p.gen.FillWithTestData(&punishments.WarnResponse{}).(*punishments.WarnResponse), nil
}

func (p *PunishmentsServer) AdminJail(ctx context.Context, request *punishments.AdminJailRequest) (*punishments.AdminJailResponse, error) {
	return p.gen.FillWithTestData(&punishments.AdminJailResponse{}).(*punishments.AdminJailResponse), nil
}

func (p *PunishmentsServer) MuteGlobalChats(ctx context.Context, request *punishments.MuteGlobalChatsRequest) (*punishments.MuteGlobalChatsResponse, error) {
	return p.gen.FillWithTestData(&punishments.MuteGlobalChatsResponse{}).(*punishments.MuteGlobalChatsResponse), nil
}

func (p *PunishmentsServer) UnBan(ctx context.Context, request *punishments.UnBanRequest) (*punishments.UnBanResponse, error) {
	return p.gen.FillWithTestData(&punishments.UnBanResponse{}).(*punishments.UnBanResponse), nil
}

func (p *PunishmentsServer) UnBlock(ctx context.Context, request *punishments.UnBlockRequest) (*punishments.UnBlockResponse, error) {
	return p.gen.FillWithTestData(&punishments.UnBlockResponse{}).(*punishments.UnBlockResponse), nil
}

func (p *PunishmentsServer) UnWarn(ctx context.Context, request *punishments.UnWarnRequest) (*punishments.UnWarnResponse, error) {
	return p.gen.FillWithTestData(&punishments.UnWarnResponse{}).(*punishments.UnWarnResponse), nil
}

func (p *PunishmentsServer) UnAdminJail(ctx context.Context, request *punishments.UnAdminJailRequest) (*punishments.UnAdminJailResponse, error) {
	return p.gen.FillWithTestData(&punishments.UnAdminJailResponse{}).(*punishments.UnAdminJailResponse), nil
}

func (p *PunishmentsServer) UnMuteGlobalChats(ctx context.Context, request *punishments.UnMuteGlobalChatsRequest) (*punishments.UnMuteGlobalChatsResponse, error) {
	return p.gen.FillWithTestData(&punishments.UnMuteGlobalChatsResponse{}).(*punishments.UnMuteGlobalChatsResponse), nil
}

func (p *PunishmentsServer) GetPlayerBans(ctx context.Context, request *punishments.GetPlayerBansRequest) (*punishments.GetPlayerBansResponse, error) {
	return p.gen.FillWithTestData(&punishments.GetPlayerBansResponse{}).(*punishments.GetPlayerBansResponse), nil
}

func (p *PunishmentsServer) GetPlayerWarns(ctx context.Context, request *punishments.GetPlayerWarnsRequest) (*punishments.GetPlayerWarnsResponse, error) {
	return p.gen.FillWithTestData(&punishments.GetPlayerWarnsResponse{}).(*punishments.GetPlayerWarnsResponse), nil
}

func (p *PunishmentsServer) GetPlayerAdminJail(ctx context.Context, request *punishments.GetPlayerAdminJailRequest) (*punishments.GetPlayerAdminJailResponse, error) {
	return p.gen.FillWithTestData(&punishments.GetPlayerAdminJailResponse{}).(*punishments.GetPlayerAdminJailResponse), nil
}

func (p *PunishmentsServer) GetBan(ctx context.Context, request *punishments.GetBanRequest) (*punishments.BanMessage, error) {
	return p.gen.FillWithTestData(&punishments.BanMessage{}).(*punishments.BanMessage), nil
}

func (p *PunishmentsServer) GetWarn(ctx context.Context, request *punishments.GetWarnRequest) (*punishments.WarnMessage, error) {
	return p.gen.FillWithTestData(&punishments.WarnMessage{}).(*punishments.WarnMessage), nil
}

func (p *PunishmentsServer) GetBlock(ctx context.Context, request *punishments.GetBlockRequest) (*punishments.BlockMessage, error) {
	return p.gen.FillWithTestData(&punishments.BlockMessage{}).(*punishments.BlockMessage), nil
}

func (p *PunishmentsServer) IsPlayerBanned(ctx context.Context, request *punishments.IsPlayerBannedRequest) (*punishments.IsPlayerBannedResponse, error) {
	return p.gen.FillWithTestData(&punishments.IsPlayerBannedResponse{}).(*punishments.IsPlayerBannedResponse), nil
}

func (p *PunishmentsServer) IsCharacterBlocked(ctx context.Context, request *punishments.IsCharacterBlockedRequest) (*punishments.IsCharacterBlockedResponse, error) {
	return p.gen.FillWithTestData(&punishments.IsCharacterBlockedResponse{}).(*punishments.IsCharacterBlockedResponse), nil
}

func (p *PunishmentsServer) IsCharacterJailed(ctx context.Context, request *punishments.IsCharacterJailedRequest) (*punishments.IsCharacterJailedResponse, error) {
	return p.gen.FillWithTestData(&punishments.IsCharacterJailedResponse{}).(*punishments.IsCharacterJailedResponse), nil
}

func (p *PunishmentsServer) WatchBans(request *punishments.WatchBansRequest, server punishments.MruVPunishmentsService_WatchBansServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchBlocks(request *punishments.WatchBlocksRequest, server punishments.MruVPunishmentsService_WatchBlocksServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchWarns(request *punishments.WatchWarnsRequest, server punishments.MruVPunishmentsService_WatchWarnsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchAdminJails(request *punishments.WatchAdminJailsRequest, server punishments.MruVPunishmentsService_WatchAdminJailsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchUnBans(request *punishments.WatchUnBansRequest, server punishments.MruVPunishmentsService_WatchUnBansServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchUnBlocks(request *punishments.WatchUnBlocksRequest, server punishments.MruVPunishmentsService_WatchUnBlocksServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchUnWarns(request *punishments.WatchUnWarnsRequest, server punishments.MruVPunishmentsService_WatchUnWarnsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchUnAdminJails(request *punishments.WatchUnAdminJailsRequest, server punishments.MruVPunishmentsService_WatchUnAdminJailsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchPlayerPunishments(request *punishments.WatchPlayerPunishmentsRequest, server punishments.MruVPunishmentsService_WatchPlayerPunishmentsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchPlayerAcquittals(request *punishments.WatchPlayerAcquittalsRequest, server punishments.MruVPunishmentsService_WatchPlayerAcquittalsServer) error {
	panic("implement me")
}

func (p *PunishmentsServer) WatchPunishments(request *punishments.WatchPunishmentsRequest, server punishments.MruVPunishmentsService_WatchPunishmentsServer) error {
	panic("implement me")
}
