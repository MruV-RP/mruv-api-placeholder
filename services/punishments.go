package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/punishments"
)

type PunishmentsServer struct {
}

func (p *PunishmentsServer) Ban(ctx context.Context, request *punishments.BanRequest) (*punishments.BanResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) Block(ctx context.Context, request *punishments.BlockRequest) (*punishments.BlockResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) Warn(ctx context.Context, request *punishments.WarnRequest) (*punishments.WarnResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) AdminJail(ctx context.Context, request *punishments.AdminJailRequest) (*punishments.AdminJailResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) MuteGlobalChats(ctx context.Context, request *punishments.MuteGlobalChatsRequest) (*punishments.MuteGlobalChatsResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) UnBan(ctx context.Context, request *punishments.UnBanRequest) (*punishments.UnBanResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) UnBlock(ctx context.Context, request *punishments.UnBlockRequest) (*punishments.UnBlockResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) UnWarn(ctx context.Context, request *punishments.UnWarnRequest) (*punishments.UnWarnResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) UnAdminJail(ctx context.Context, request *punishments.UnAdminJailRequest) (*punishments.UnAdminJailResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) UnMuteGlobalChats(ctx context.Context, request *punishments.UnMuteGlobalChatsRequest) (*punishments.UnMuteGlobalChatsResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetPlayerBans(ctx context.Context, request *punishments.GetPlayerBansRequest) (*punishments.GetPlayerBansResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetPlayerWarns(ctx context.Context, request *punishments.GetPlayerWarnsRequest) (*punishments.GetPlayerWarnsResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetPlayerAdminJail(ctx context.Context, request *punishments.GetPlayerAdminJailRequest) (*punishments.GetPlayerAdminJailResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetBan(ctx context.Context, request *punishments.GetBanRequest) (*punishments.BanMessage, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetWarn(ctx context.Context, request *punishments.GetWarnRequest) (*punishments.WarnMessage, error) {
	panic("implement me")
}

func (p *PunishmentsServer) GetBlock(ctx context.Context, request *punishments.GetBlockRequest) (*punishments.BlockMessage, error) {
	panic("implement me")
}

func (p *PunishmentsServer) IsPlayerBanned(ctx context.Context, request *punishments.IsPlayerBannedRequest) (*punishments.IsPlayerBannedResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) IsCharacterBlocked(ctx context.Context, request *punishments.IsCharacterBlockedRequest) (*punishments.IsCharacterBlockedResponse, error) {
	panic("implement me")
}

func (p *PunishmentsServer) IsCharacterJailed(ctx context.Context, request *punishments.IsCharacterJailedRequest) (*punishments.IsCharacterJailedResponse, error) {
	panic("implement me")
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

func NewPunishmentsServer() *PunishmentsServer {
	return &PunishmentsServer{}
}
