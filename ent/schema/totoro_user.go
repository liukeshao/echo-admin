package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	ge "github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/ent/hook"
	"strings"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(IdGenerator).Comment("用户ID"),
		field.String("name").Optional().Comment("用户账号"),
		field.String("display_name").Optional().Comment("显示名称"),
		field.String("email").Optional().Comment("用户邮箱"),
		field.String("phone").Optional().Comment("手机号码"),
		field.String("gender").Optional().Default("0").Comment("用户性别（0保密 1男 2女）"),
		field.String("avatar").Optional().Comment("头像路径"),
		field.String("password").Sensitive().Optional().Comment("密码"),
		field.Int("non_locked").Optional().Default(1).Comment("锁定标志（0锁定 1正常）"),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "totoro_user"}}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ge.UserMutation) (ent.Value, error) {
					if v, exists := m.Email(); exists {
						m.SetEmail(strings.ToLower(v))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("email", "delete_time").Unique(),
	}
}
