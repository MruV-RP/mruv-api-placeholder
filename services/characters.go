package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/characters"
	"github.com/MruV-RP/mruv-pb-go/common"
)

type CharactersServer struct {
}

func (c *CharactersServer) CreateCharacter(ctx context.Context, request *characters.CreateCharacterRequest) (*characters.CreateCharacterResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) GetCharacter(ctx context.Context, request *characters.GetCharacterRequest) (*characters.GetCharacterResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) UpdateCharacter(ctx context.Context, request *characters.UpdateCharacterRequest) (*characters.UpdateCharacterResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) DeleteCharacter(ctx context.Context, request *characters.DeleteCharacterRequest) (*characters.DeleteCharacterResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) PermanentCharacterKill(ctx context.Context, id *characters.CharacterID) (*characters.CharacterID, error) {
	panic("implement me")
}

func (c *CharactersServer) ChangeClothes(ctx context.Context, request *characters.ChangeClothesRequest) (*characters.ChangeClothesResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) DeathsStream(request *characters.DeathStreamRequest, server characters.MruVCharactersService_DeathsStreamServer) error {
	panic("implement me")
}

func (c *CharactersServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	panic("implement me")
}

func (c *CharactersServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	panic("implement me")
}

func NewCharactersServer() *CharactersServer {
	return &CharactersServer{}
}
