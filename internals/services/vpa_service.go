package services

import (
	"context"

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