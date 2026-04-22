package services

import (
	"context"

	"github.com/swastiijain24/psp/internals/httpclient"
)

type AccountService interface {
	DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error)
	SetMpin(ctx context.Context, vpaId string, mpinEn string) error
	ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string, newMpinEn string) error
	GetBalance(ctx context.Context, vpaId string, mpinEn string) (int64, error)
}

type AccountSvc struct {
	bankServiceClient httpclient.HttpClient
	vpaService VpaService
}

func NewAccountService(bankServiceClient httpclient.HttpClient, vpaService VpaService) AccountService {
	return &AccountSvc{
		bankServiceClient: bankServiceClient,
		vpaService: vpaService,
	}
}

func (s *AccountSvc) DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error){
	return s.bankServiceClient.DiscoverAccounts(ctx, phone, bankCode)
}

func (s *AccountSvc) SetMpin(ctx context.Context, vpaId string, mpinEn string) error {

	accountId, bankCode, err:=s.vpaService.ResolveVpa(ctx, vpaId)
	if err != nil {
		return err 
	}

	return s.bankServiceClient.SetMpin(ctx, accountId, bankCode, mpinEn)
}

func (s *AccountSvc) ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string, newMpinEn string) error {

	accountId, bankCode, err:=s.vpaService.ResolveVpa(ctx, vpaId)
	if err != nil {
		return err 
	}

	return s.bankServiceClient.ChangeMpin(ctx, accountId, bankCode, oldMpinEn, newMpinEn)
}

func (s *AccountSvc) GetBalance(ctx context.Context, vpaId string, mpinEn string) (int64, error) {
	accountId, bankCode, err:=s.vpaService.ResolveVpa(ctx, vpaId)
	if err != nil {
		return 0, err 
	}
	return s.bankServiceClient.GetBalance(ctx, accountId, bankCode, mpinEn)
}
