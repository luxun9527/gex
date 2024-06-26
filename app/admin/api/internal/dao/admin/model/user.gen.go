// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        uint32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nickname  string `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`  // 昵称
	Username  string `gorm:"column:username;not null;comment:用户名" json:"username"` // 用户名
	Password  string `gorm:"column:password;not null" json:"password"`
	CreatedAt uint32 `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt uint32 `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt uint32 `gorm:"column:deleted_at;not null" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
