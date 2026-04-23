package services

import (
	"context"
	repo "github.com/swastiijain24/psp/internals/repositories"
)

type ApiKeyService interface {
	GetAPIKeyByPspId(ctx context.Context, pspId string) (repo.PspRegistration, error)
	IsValid(ctx context.Context, pspId string) (bool , error)
}

type ApiKeySvc struct {
	repo repo.Querier
}

func NewApiKeyService(repo repo.Querier) ApiKeyService {
	return &ApiKeySvc{
		repo: repo,
	}
}

func (s* ApiKeySvc) GetAPIKeyByPspId(ctx context.Context, pspId string) (repo.PspRegistration, error){
	return s.repo.GetPspRegistration(ctx, pspId)
	
}

func (s* ApiKeySvc) IsValid(ctx context.Context, pspId string) (bool , error){
	return s.repo.IsActive(ctx, pspId)
	
} 
