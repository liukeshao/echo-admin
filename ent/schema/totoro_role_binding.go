package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type RoleBinding struct {
	ent.Schema
}

func (RoleBinding) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Comment("角色名称"),
		field.String("display_name").Optional().Comment("角色显示名称"),
		field.String("role_name").Comment("角色名称"),
		field.String("user_id").Comment("用户名称"),
		field.Int64("deleted_time").Optional().Comment("删除标志（0代表存在）"),
	}
}

func (RoleBinding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}
func (RoleBinding) Edges() []ent.Edge {
	return nil
}
func (RoleBinding) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "totoro_role_binding"}}
}
