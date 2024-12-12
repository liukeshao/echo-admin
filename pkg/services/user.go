package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/ent/user"
)

type UserService struct {
	orm *ent.Client
}

func NewUserService(orm *ent.Client) *UserService {
	return &UserService{orm: orm}
}

func (s *UserService) FindById(ctx context.Context, id uint64) (*ent.User, error) {
	u, err := s.orm.User.Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := s.orm.User.Query().
		Where(user.Email(strings.ToLower(email))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user %s: %w", email, err)
	}
	return u, nil
}

func (s *UserService) Create(ctx context.Context, email, password string) (*ent.User, error) {
	u, err := s.orm.User.Create().
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}
