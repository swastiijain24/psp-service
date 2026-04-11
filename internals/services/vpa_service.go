package services

import "context"

type VpaService struct {
}

func NewVpaService() *VpaService {
	return &VpaService{}
}

func (s *VpaService) ResolveVpa(ctx context.Context, VPA string) string {
	//this will either resolve by itself or call the psp to resolve the add
	return ""
}