package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171129_220237 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171129_220237{}
	m.Created = "20171129_220237"

	migration.Register("User_20171129_220237", m)
}

// Run the migrations
func (m *User_20171129_220237) Up() {
	m.SQL("alter table user modify column password varchar(128)")

}

// Reverse the migrations
func (m *User_20171129_220237) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
