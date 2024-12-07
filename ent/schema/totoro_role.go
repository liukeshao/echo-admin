package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键"),
		field.String("name").Comment("角色名称"),
		field.String("display_name").Optional().Comment("角色显示名称"),
		field.String("remark").Optional().Comment("角色描述"),
		field.Int64("deleted_time").Comment("删除标志（0代表存在）"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}

func (Role) Edges() []ent.Edge {
	return nil
}
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "totoro_role"}}
}
