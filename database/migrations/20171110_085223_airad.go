package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AirAd_20171110_085223 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AirAd_20171110_085223{}
	m.Created = "20171110_085223"

	migration.Register("AirAd_20171110_085223", m)
}

// Run the migrations
func (m *AirAd_20171110_085223) Up() {
	m.SQL("ALTER TABLE air_ad ADD temperature varchar(10), ADD humidity varchar(10), " +
		"ADD latitude varchar(10), ADD longitude varchar(10), ADD aqi_quality varchar(10), " +
		"ADD suggest text(140)")
}

// Reverse the migrations
func (m *AirAd_20171110_085223) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
