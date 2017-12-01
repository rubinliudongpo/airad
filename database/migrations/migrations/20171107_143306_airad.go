package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AirAd_20171107_143306 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AirAd_20171107_143306{}
	m.Created = "20171107_143306"

	migration.Register("AirAd_20171107_143306", m)
}

// Run the migrations
func (m *AirAd_20171107_143306) Up() {
	m.SQL("CREATE TABLE air_ad (id int not null primary key auto_increment)")
}

// Reverse the migrations
func (m *AirAd_20171107_143306) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
