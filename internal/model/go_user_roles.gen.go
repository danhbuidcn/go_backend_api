// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGoUserRole = "go_user_roles"

// GoUserRole mapped from table <go_user_roles>
type GoUserRole struct {
	UserID int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;primaryKey;comment:'primary Key is ID'" json:"role_id"` // 'primary Key is ID'
}

// TableName GoUserRole's table name
func (*GoUserRole) TableName() string {
	return TableNameGoUserRole
}
