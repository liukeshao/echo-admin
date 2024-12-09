package services

import (
	"context"
	"github.com/liukeshao/echo-admin/pkg/errors"
	"strings"

	"github.com/joomcode/errorx"
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
	switch err.(type) {
	case *ent.NotFoundError:
		return nil, errors.RecordNotFound.New(email)
	case nil:
	default:
		return nil, errorx.Decorate(err, email)
	}
	return u, nil
}

func (s *UserService) Create(ctx context.Context, email, password string) (*ent.User, error) {
	u, err := s.orm.User.Create().
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, errorx.Decorate(err, "create user database failed")
	}
	return u, nil
}
