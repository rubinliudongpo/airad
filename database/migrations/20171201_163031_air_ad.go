package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AirAd_20171201_163031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AirAd_20171201_163031{}
	m.Created = "20171201_163031"

	migration.Register("AirAd_20171201_163031", m)
}

// Run the migrations
func (m *AirAd_20171201_163031) Up() {
	m.SQL("ALTER TABLE air_ad ADD device_id int(11) unsigned NOT NULL")
}

// Reverse the migrations
func (m *AirAd_20171201_163031) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
