// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/liukeshao/echo-admin/ent/rolebinding"
)

// RoleBinding is the model entity for the RoleBinding schema.
type RoleBinding struct {
	config `json:"-"`
	// ID of the ent.
	// ID
	ID string `json:"id,omitempty"`
	// delete time stamp
	DeleteTime int64 `json:"delete_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 角色名称
	Name string `json:"name,omitempty"`
	// 角色显示名称
	DisplayName string `json:"display_name,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty"`
	// 用户名称
	UserID string `json:"user_id,omitempty"`
	// 删除标志（0代表存在）
	DeletedTime  int64 `json:"deleted_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleBinding) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rolebinding.FieldDeleteTime, rolebinding.FieldDeletedTime:
			values[i] = new(sql.NullInt64)
		case rolebinding.FieldID, rolebinding.FieldName, rolebinding.FieldDisplayName, rolebinding.FieldRoleName, rolebinding.FieldUserID:
			values[i] = new(sql.NullString)
		case rolebinding.FieldCreateTime, rolebinding.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleBinding fields.
func (rb *RoleBinding) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rolebinding.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rb.ID = value.String
			}
		case rolebinding.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				rb.DeleteTime = value.Int64
			}
		case rolebinding.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rb.CreateTime = value.Time
			}
		case rolebinding.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				rb.UpdateTime = value.Time
			}
		case rolebinding.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rb.Name = value.String
			}
		case rolebinding.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				rb.DisplayName = value.String
			}
		case rolebinding.FieldRoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_name", values[i])
			} else if value.Valid {
				rb.RoleName = value.String
			}
		case rolebinding.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				rb.UserID = value.String
			}
		case rolebinding.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				rb.DeletedTime = value.Int64
			}
		default:
			rb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoleBinding.
// This includes values selected through modifiers, order, etc.
func (rb *RoleBinding) Value(name string) (ent.Value, error) {
	return rb.selectValues.Get(name)
}

// Update returns a builder for updating this RoleBinding.
// Note that you need to call RoleBinding.Unwrap() before calling this method if this RoleBinding
// was returned from a transaction, and the transaction was committed or rolled back.
func (rb *RoleBinding) Update() *RoleBindingUpdateOne {
	return NewRoleBindingClient(rb.config).UpdateOne(rb)
}

// Unwrap unwraps the RoleBinding entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rb *RoleBinding) Unwrap() *RoleBinding {
	_tx, ok := rb.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleBinding is not a transactional entity")
	}
	rb.config.driver = _tx.drv
	return rb
}

// String implements the fmt.Stringer.
func (rb *RoleBinding) String() string {
	var builder strings.Builder
	builder.WriteString("RoleBinding(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rb.ID))
	builder.WriteString("delete_time=")
	builder.WriteString(fmt.Sprintf("%v", rb.DeleteTime))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(rb.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(rb.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rb.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(rb.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("role_name=")
	builder.WriteString(rb.RoleName)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(rb.UserID)
	builder.WriteString(", ")
	builder.WriteString("deleted_time=")
	builder.WriteString(fmt.Sprintf("%v", rb.DeletedTime))
	builder.WriteByte(')')
	return builder.String()
}

// RoleBindings is a parsable slice of RoleBinding.
type RoleBindings []*RoleBinding
