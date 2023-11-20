package models

import (
	"github.com/goravel/framework/database/orm"
)

type User2 struct {
	orm.Model
	Id		 int
	orm.SoftDeletes
}

func (r *User2) TableName() string {
	return "users2"
}

func (r *User2) Connection() string {
	return "mysql2"
}
