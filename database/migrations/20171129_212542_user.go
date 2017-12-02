package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171129_212542 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171129_212542{}
	m.Created = "20171129_212542"

	migration.Register("User_20171129_212542", m)
}

// Run the migrations
func (m *User_20171129_212542) Up() {
	m.SQL("ALTER TABLE `user` ADD COLUMN `salt` varchar(64) NOT NULL  DEFAULT ''")

}

// Reverse the migrations
func (m *User_20171129_212542) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
