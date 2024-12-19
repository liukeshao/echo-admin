package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/liukeshao/echo-admin/pkg/uid"
	"time"
)

type BaseFiledMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseFiledMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uid.ULID).Comment("ID"),
		field.Time("create_time").Immutable().Default(time.Now),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}
