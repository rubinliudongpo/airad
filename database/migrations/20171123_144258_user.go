package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171123_144258 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171123_144258{}
	m.Created = "20171123_144258"

	migration.Register("User_20171123_144258", m)
}

// Run the migrations
func (m *User_20171123_144258) Up() {
	m.SQL("alter TABLE user add unique(username, id)")
}

// Reverse the migrations
func (m *User_20171123_144258) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
