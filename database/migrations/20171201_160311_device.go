package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Device_20171201_160311 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Device_20171201_160311{}
	m.Created = "20171201_160311"

	migration.Register("Device_20171201_160311", m)
}

// Run the migrations
func (m *Device_20171201_160311) Up() {
	m.SQL("create table device(id int(11) unsigned unique NOT NULL AUTO_INCREMENT," +
		"device_name varchar(32) unique NOT NULL DEFAULT 'device name', status tinyint(1) default 0, " +
		"address varchar(50) NULL DEFAULT 'address', created_at int(11) unsigned NOT NULL DEFAULT '0'," +
		"updated_at int(11) unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`))" +
		"ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4")

}

// Reverse the migrations
func (m *Device_20171201_160311) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
