package services

import (
	"context"
	"github.com/samber/oops"
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

func (s *UserService) FindById(ctx context.Context, id string) (*ent.User, error) {
	eb := oops.Code("user.not.exist").In("user").With("id", id).Public("not found")
	u, err := s.orm.User.Query().
		Where(user.ID(id)).
		Only(ctx)
	return u, eb.Wrap(err)
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	eb := oops.Code("user.not.exist").In("user:find_by_email").With("email", email).Public("not found")
	u, err := s.orm.User.Query().
		Where(user.Email(strings.ToLower(email))).
		Only(ctx)
	return u, eb.Wrap(err)
}

func (s *UserService) Create(ctx context.Context, email, password string) (*ent.User, error) {
	eb := oops.Code("user.create.error").In("user:create").With("email", email).Public("user create error")
	u, err := s.orm.User.Create().
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	return u, eb.Wrap(err)
}

func (s *UserService) UpdatePassword(ctx context.Context, id string, encryptPassword string) (*ent.User, error) {
	eb := oops.Code("user.update_password.error").In("user:update_password").With("id", id).Public("user update password error")
	param := ent.User{
		ID:       id,
		Password: encryptPassword,
	}
	u, err := s.orm.User.UpdateOne(&param).Save(ctx)
	return u, eb.Wrap(err)
}
