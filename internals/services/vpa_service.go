package services

import (
	"context"
	"fmt"

	repo "github.com/swastiijain24/psp/internals/repositories"
)

type VpaService struct {
	repo repo.Querier 
}

func NewVpaService(repo repo.Querier) *VpaService {
	return &VpaService{
		repo: repo,
	}
}

func (s *VpaService) ResolveVpa(ctx context.Context, VPA string) (string, string, error) {

	mapping, err := s.repo.GetVpaMapping(ctx, VPA)
	if err != nil {
		return "", "", err 
	}

	return mapping.AccountID, mapping.BankCode, nil 
}

func (s *VpaService) RegisterVpa(ctx context.Context, vpaId string, accountId string, bankCode string) error  {
	exists, err := s.repo.CheckVpaExists(ctx, vpaId)
	if err != nil {
		return err 
	}
	if exists == true {
		return fmt.Errorf("VPAId already registered")
	}

	_, err =s.repo.CreateVpaMapping(ctx, repo.CreateVpaMappingParams{
		VpaID: vpaId,
		AccountID: accountId,
		BankCode: bankCode,
	}) 
	if err != nil {
		return err 
	}

	return nil 
}