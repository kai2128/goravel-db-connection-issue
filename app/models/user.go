package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Id		 int
	orm.SoftDeletes
}

func (r *User) TableName() string {
	return "users"
}

func (r *User) Connection() string {
	return "mysql"
}
