package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Cluster struct {
	ent.Schema
}

func (Cluster) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(IdGenerator).Comment("主键"),
		field.Uint64("org_id").Comment("组织id"),
		field.String("name").MaxLen(255).Comment("名称"),
		field.String("display_name").Optional().Comment("显示名称"),
		field.String("description").Optional().Comment("描述"),
		field.Text("config").Sensitive().Optional().Comment("kube config"),
	}
}

func (Cluster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}

func (Cluster) Edges() []ent.Edge {
	return nil
}
func (Cluster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "totoro_cluster"}}
}
