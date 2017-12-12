package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Device_20171208_141841 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Device_20171208_141841{}
	m.Created = "20171208_141841"

	migration.Register("Device_20171208_141841", m)
}

// Run the migrations
func (m *Device_20171208_141841) Up() {
	m.SQL("alter table device add COLUMN `user_id` int(11) unsigned not null comment '用户ID'")
}

// Reverse the migrations
func (m *Device_20171208_141841) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
