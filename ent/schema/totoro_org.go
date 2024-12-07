package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Org struct {
	ent.Schema
}

func (Org) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键"),
		field.String("name").MaxLen(255).Comment("组织名称"),
		field.String("display_name").Optional().Comment("显示名称"),
		field.String("description").Optional().Comment("描述"),
		field.String("logo").MaxLen(512).Comment("logo"),
		field.String("status").MaxLen(64).Comment("状态"),
		field.String("type").MaxLen(64).Comment("类型"),
	}
}

func (Org) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}

func (Org) Edges() []ent.Edge {
	return nil
}
func (Org) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "totoro_org"}}
}
