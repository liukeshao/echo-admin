// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/liukeshao/echo-admin/ent/org"
	"github.com/liukeshao/echo-admin/ent/role"
	"github.com/liukeshao/echo-admin/ent/rolebinding"
	"github.com/liukeshao/echo-admin/ent/schema"
	"github.com/liukeshao/echo-admin/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orgMixin := schema.Org{}.Mixin()
	orgMixinHooks0 := orgMixin[0].Hooks()
	org.Hooks[0] = orgMixinHooks0[0]
	orgMixinInters0 := orgMixin[0].Interceptors()
	org.Interceptors[0] = orgMixinInters0[0]
	orgMixinFields0 := orgMixin[0].Fields()
	_ = orgMixinFields0
	orgMixinFields1 := orgMixin[1].Fields()
	_ = orgMixinFields1
	orgFields := schema.Org{}.Fields()
	_ = orgFields
	// orgDescDeleteTime is the schema descriptor for delete_time field.
	orgDescDeleteTime := orgMixinFields0[0].Descriptor()
	// org.DefaultDeleteTime holds the default value on creation for the delete_time field.
	org.DefaultDeleteTime = orgDescDeleteTime.Default.(int64)
	// orgDescCreateTime is the schema descriptor for create_time field.
	orgDescCreateTime := orgMixinFields1[0].Descriptor()
	// org.DefaultCreateTime holds the default value on creation for the create_time field.
	org.DefaultCreateTime = orgDescCreateTime.Default.(func() time.Time)
	// orgDescUpdateTime is the schema descriptor for update_time field.
	orgDescUpdateTime := orgMixinFields1[1].Descriptor()
	// org.DefaultUpdateTime holds the default value on creation for the update_time field.
	org.DefaultUpdateTime = orgDescUpdateTime.Default.(func() time.Time)
	// org.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	org.UpdateDefaultUpdateTime = orgDescUpdateTime.UpdateDefault.(func() time.Time)
	// orgDescName is the schema descriptor for name field.
	orgDescName := orgFields[1].Descriptor()
	// org.NameValidator is a validator for the "name" field. It is called by the builders before save.
	org.NameValidator = orgDescName.Validators[0].(func(string) error)
	// orgDescLogo is the schema descriptor for logo field.
	orgDescLogo := orgFields[4].Descriptor()
	// org.LogoValidator is a validator for the "logo" field. It is called by the builders before save.
	org.LogoValidator = orgDescLogo.Validators[0].(func(string) error)
	// orgDescStatus is the schema descriptor for status field.
	orgDescStatus := orgFields[5].Descriptor()
	// org.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	org.StatusValidator = orgDescStatus.Validators[0].(func(string) error)
	// orgDescType is the schema descriptor for type field.
	orgDescType := orgFields[6].Descriptor()
	// org.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	org.TypeValidator = orgDescType.Validators[0].(func(string) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinHooks0 := roleMixin[0].Hooks()
	role.Hooks[0] = roleMixinHooks0[0]
	roleMixinInters0 := roleMixin[0].Interceptors()
	role.Interceptors[0] = roleMixinInters0[0]
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescDeleteTime is the schema descriptor for delete_time field.
	roleDescDeleteTime := roleMixinFields0[0].Descriptor()
	// role.DefaultDeleteTime holds the default value on creation for the delete_time field.
	role.DefaultDeleteTime = roleDescDeleteTime.Default.(int64)
	// roleDescCreateTime is the schema descriptor for create_time field.
	roleDescCreateTime := roleMixinFields1[0].Descriptor()
	// role.DefaultCreateTime holds the default value on creation for the create_time field.
	role.DefaultCreateTime = roleDescCreateTime.Default.(func() time.Time)
	// roleDescUpdateTime is the schema descriptor for update_time field.
	roleDescUpdateTime := roleMixinFields1[1].Descriptor()
	// role.DefaultUpdateTime holds the default value on creation for the update_time field.
	role.DefaultUpdateTime = roleDescUpdateTime.Default.(func() time.Time)
	// role.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	role.UpdateDefaultUpdateTime = roleDescUpdateTime.UpdateDefault.(func() time.Time)
	rolebindingMixin := schema.RoleBinding{}.Mixin()
	rolebindingMixinHooks0 := rolebindingMixin[0].Hooks()
	rolebinding.Hooks[0] = rolebindingMixinHooks0[0]
	rolebindingMixinInters0 := rolebindingMixin[0].Interceptors()
	rolebinding.Interceptors[0] = rolebindingMixinInters0[0]
	rolebindingMixinFields0 := rolebindingMixin[0].Fields()
	_ = rolebindingMixinFields0
	rolebindingMixinFields1 := rolebindingMixin[1].Fields()
	_ = rolebindingMixinFields1
	rolebindingFields := schema.RoleBinding{}.Fields()
	_ = rolebindingFields
	// rolebindingDescDeleteTime is the schema descriptor for delete_time field.
	rolebindingDescDeleteTime := rolebindingMixinFields0[0].Descriptor()
	// rolebinding.DefaultDeleteTime holds the default value on creation for the delete_time field.
	rolebinding.DefaultDeleteTime = rolebindingDescDeleteTime.Default.(int64)
	// rolebindingDescCreateTime is the schema descriptor for create_time field.
	rolebindingDescCreateTime := rolebindingMixinFields1[0].Descriptor()
	// rolebinding.DefaultCreateTime holds the default value on creation for the create_time field.
	rolebinding.DefaultCreateTime = rolebindingDescCreateTime.Default.(func() time.Time)
	// rolebindingDescUpdateTime is the schema descriptor for update_time field.
	rolebindingDescUpdateTime := rolebindingMixinFields1[1].Descriptor()
	// rolebinding.DefaultUpdateTime holds the default value on creation for the update_time field.
	rolebinding.DefaultUpdateTime = rolebindingDescUpdateTime.Default.(func() time.Time)
	// rolebinding.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	rolebinding.UpdateDefaultUpdateTime = rolebindingDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	user.Hooks[1] = userHooks[0]
	userMixinInters0 := userMixin[0].Interceptors()
	user.Interceptors[0] = userMixinInters0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDeleteTime is the schema descriptor for delete_time field.
	userDescDeleteTime := userMixinFields0[0].Descriptor()
	// user.DefaultDeleteTime holds the default value on creation for the delete_time field.
	user.DefaultDeleteTime = userDescDeleteTime.Default.(int64)
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields1[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields1[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescGender is the schema descriptor for gender field.
	userDescGender := userFields[5].Descriptor()
	// user.DefaultGender holds the default value on creation for the gender field.
	user.DefaultGender = userDescGender.Default.(string)
	// userDescNonLocked is the schema descriptor for non_locked field.
	userDescNonLocked := userFields[8].Descriptor()
	// user.DefaultNonLocked holds the default value on creation for the non_locked field.
	user.DefaultNonLocked = userDescNonLocked.Default.(int)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uint64)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)