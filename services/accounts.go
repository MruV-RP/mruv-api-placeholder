package services

import (
	"context"
	"github.com/MruV-RP/mruv-api-placeholder/generator"
	"github.com/MruV-RP/mruv-api-placeholder/utils"
	"github.com/MruV-RP/mruv-pb-go/accounts"
)

type AccountsServer struct {
	gen generator.IGenerator
}

func NewAccountsServer(gen generator.IGenerator) *AccountsServer {
	return &AccountsServer{gen: gen}
}

func (a *AccountsServer) RegisterAccount(ctx context.Context, request *accounts.RegisterAccountRequest) (*accounts.RegisterAccountResponse, error) {
	if request.Login == "good" {
		return &accounts.RegisterAccountResponse{
			Success:   true,
			AccountId: 1,
		}, nil
	} else if request.Login == "bad" {
		return &accounts.RegisterAccountResponse{
			Success:   false,
			AccountId: 0,
		}, nil
	} else if request.Login == "error" {
		return nil, utils.TestError
	} else {
		return nil, utils.UsageStringError
	}
}

func (a *AccountsServer) LogIn(ctx context.Context, request *accounts.LogInRequest) (*accounts.LogInResponse, error) {
	if request.Login == "good" {
		return &accounts.LogInResponse{
			Success:   true,
			AccountId: 1,
		}, nil
	} else if request.Login == "bad" {
		return &accounts.LogInResponse{
			Success:   false,
			AccountId: 0,
		}, nil
	} else if request.Login == "error" {
		return nil, utils.TestError
	} else {
		return nil, utils.UsageStringError
	}
}

func (a *AccountsServer) IsAccountExist(ctx context.Context, request *accounts.IsAccountExistRequest) (*accounts.IsAccountExistResponse, error) {
	if request.Login == "good" {
		return &accounts.IsAccountExistResponse{
			Exists: true,
			Id:     1,
		}, nil
	} else if request.Login == "bad" {
		return &accounts.IsAccountExistResponse{
			Exists: false,
			Id:     0,
		}, nil
	} else if request.Login == "error" {
		return nil, utils.TestError
	} else {
		return nil, utils.UsageStringError
	}
}

func (a *AccountsServer) GetAccount(ctx context.Context, request *accounts.GetAccountRequest) (*accounts.GetAccountResponse, error) {
	if request.Login == "good" {
		return &accounts.GetAccountResponse{
			Login: "TestLogin",
			Email: "TestMail@gmail.com",
		}, nil
	} else if request.Login == "bad" {
		return nil, utils.NotFoundError
	} else if request.Login == "error" {
		return nil, utils.TestError
	} else {
		return nil, utils.UsageStringError
	}
}

func (a *AccountsServer) GetAccountCharacters(ctx context.Context, request *accounts.GetAccountCharactersRequest) (*accounts.GetAccountCharactersResponse, error) {
	if request.Login == "good" {
		return &accounts.GetAccountCharactersResponse{
			CharacterIds: []uint32{1},
		}, nil
	} else if request.Login == "bad" {
		return &accounts.GetAccountCharactersResponse{
			CharacterIds: []uint32{},
		}, nil
	} else if request.Login == "error" {
		return nil, utils.TestError
	} else {
		return nil, utils.UsageStringError
	}
}
