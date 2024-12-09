package services

import (
	"context"
	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/types"
)

type OrgService struct {
	orm *ent.Client
}

func NewOrgService(orm *ent.Client) *OrgService {
	return &OrgService{orm: orm}
}

func (s *OrgService) Create(ctx context.Context, input *types.OrgCreateInput) (*ent.Org, error) {
	org, err := s.orm.Org.Create().
		SetName(input.Name).
		SetDisplayName(input.DisplayName).
		SetDescription(input.Description).
		SetLogo(input.Logo).
		SetType(input.Type).
		SetStatus(input.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return org, nil
}
