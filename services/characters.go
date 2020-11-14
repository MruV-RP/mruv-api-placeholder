package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-pb-go/characters"
	"github.com/MruV-RP/mruv-pb-go/common"
)

type CharactersServer struct {
	gen generator.IGenerator
}

func NewCharactersServer(gen generator.IGenerator) *CharactersServer {
	return &CharactersServer{gen: gen}
}

func (c *CharactersServer) CreateCharacter(ctx context.Context, request *characters.CreateCharacterRequest) (*characters.CreateCharacterResponse, error) {
	return c.gen.FillWithTestData(&characters.CreateCharacterResponse{}).(*characters.CreateCharacterResponse), nil
}

func (c *CharactersServer) GetCharacter(ctx context.Context, request *characters.GetCharacterRequest) (*characters.GetCharacterResponse, error) {
	return c.gen.FillWithTestData(&characters.GetCharacterResponse{}).(*characters.GetCharacterResponse), nil
}

func (c *CharactersServer) UpdateCharacter(ctx context.Context, request *characters.UpdateCharacterRequest) (*characters.UpdateCharacterResponse, error) {
	return c.gen.FillWithTestData(&characters.UpdateCharacterResponse{}).(*characters.UpdateCharacterResponse), nil
}

func (c *CharactersServer) DeleteCharacter(ctx context.Context, request *characters.DeleteCharacterRequest) (*characters.DeleteCharacterResponse, error) {
	return c.gen.FillWithTestData(&characters.DeleteCharacterResponse{}).(*characters.DeleteCharacterResponse), nil
}

func (c *CharactersServer) PermanentCharacterKill(ctx context.Context, id *characters.CharacterID) (*characters.CharacterID, error) {
	return c.gen.FillWithTestData(&characters.CharacterID{}).(*characters.CharacterID), nil
}

func (c *CharactersServer) ChangeClothes(ctx context.Context, request *characters.ChangeClothesRequest) (*characters.ChangeClothesResponse, error) {
	return c.gen.FillWithTestData(&characters.ChangeClothesResponse{}).(*characters.ChangeClothesResponse), nil
}

func (c *CharactersServer) DeathsStream(request *characters.DeathStreamRequest, server characters.MruVCharactersService_DeathsStreamServer) error {
	panic("implement me")
}

func (c *CharactersServer) GetServiceStatus(ctx context.Context, request *common.ServiceStatusRequest) (*common.ServiceStatusResponse, error) {
	return c.gen.FillWithTestData(&common.ServiceStatusResponse{}).(*common.ServiceStatusResponse), nil
}

func (c *CharactersServer) GetServiceVersion(ctx context.Context, request *common.VersionRequest) (*common.VersionResponse, error) {
	return c.gen.FillWithTestData(&common.VersionResponse{}).(*common.VersionResponse), nil
}
