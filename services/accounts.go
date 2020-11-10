package services

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/accounts"
)

type AccountsServer struct {
}

func NewAccountsServer() *AccountsServer {
	return &AccountsServer{}
}

func (a *AccountsServer) RegisterAccount(ctx context.Context, request *accounts.RegisterAccountRequest) (*accounts.RegisterAccountResponse, error) {
	panic("implement me")
}

func (a *AccountsServer) LogIn(ctx context.Context, request *accounts.LogInRequest) (*accounts.LogInResponse, error) {
	panic("implement me")
}

func (a *AccountsServer) IsAccountExist(ctx context.Context, request *accounts.IsAccountExistRequest) (*accounts.IsAccountExistResponse, error) {
	panic("implement me")
}

func (a *AccountsServer) GetAccount(ctx context.Context, request *accounts.GetAccountRequest) (*accounts.GetAccountResponse, error) {
	panic("implement me")
}

func (a *AccountsServer) GetAccountCharacters(ctx context.Context, request *accounts.GetAccountCharactersRequest) (*accounts.GetAccountCharactersResponse, error) {
	panic("implement me")
}
