package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171208_145743 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171208_145743{}
	m.Created = "20171208_145743"

	migration.Register("User_20171208_145743", m)
}

// Run the migrations
func (m *User_20171208_145743) Up() {
	m.SQL("ALTER TABLE `user` ADD COLUMN `device_count` int(11) unsigned not null  DEFAULT 0")

}

// Reverse the migrations
func (m *User_20171208_145743) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
