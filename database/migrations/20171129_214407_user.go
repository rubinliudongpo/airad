package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171129_214407 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171129_214407{}
	m.Created = "20171129_214407"

	migration.Register("User_20171129_214407", m)
}

// Run the migrations
func (m *User_20171129_214407) Up() {
	m.SQL("alter table user modify column salt varchar(128)")
}

// Reverse the migrations
func (m *User_20171129_214407) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
