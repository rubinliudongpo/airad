package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171205_133550 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171205_133550{}
	m.Created = "20171205_133550"

	migration.Register("User_20171205_133550", m)
}

// Run the migrations
func (m *User_20171205_133550) Up() {
	m.SQL("alter table user modify column token varchar(256)")

}

// Reverse the migrations
func (m *User_20171205_133550) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
