// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username    string `gorm:"column:username;not null;comment:用户名" json:"username"`
	Password    string `gorm:"column:password;not null;comment:密码" json:"password"`
	PhoneNumber int64  `gorm:"column:phone_number;not null;comment:手机号" json:"phone_number"`
	Status      int32  `gorm:"column:status;not null;comment:用户状态，1正常2锁定" json:"status"`
	CreatedAt   int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
