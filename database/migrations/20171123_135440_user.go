package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20171123_135440 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20171123_135440{}
	m.Created = "20171123_135440"

	migration.Register("User_20171123_135440", m)
}

// Run the migrations
func (m *User_20171123_135440) Up() {
	m.SQL("create table user(id int(11) unsigned NOT NULL AUTO_INCREMENT," +
		"username varchar(32) NOT NULL DEFAULT '', password char(32) NOT NULL DEFAULT ''," +
		"token char(32) NOT NULL, gender tinyint(1) default 0," +
		"age tinyint(2) default 20, address varchar(50) NULL DEFAULT ''," +
		"email varchar(50) NOT NULL DEFAULT '', last_login int(11) NOT NULL DEFAULT '0'," +
		"status tinyint(1) default 0, created_at int(11) unsigned NOT NULL DEFAULT '0'," +
		"updated_at int(11) unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`))" +
		"ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4")
}

// Reverse the migrations
func (m *User_20171123_135440) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
