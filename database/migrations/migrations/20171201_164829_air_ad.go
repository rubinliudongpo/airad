package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AirAd_20171201_164829 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AirAd_20171201_164829{}
	m.Created = "20171201_164829"

	migration.Register("AirAd_20171201_164829", m)
}

// Run the migrations
func (m *AirAd_20171201_164829) Up() {
	m.SQL("ALTER TABLE air_ad ADD CONSTRAINT fk_device_id FOREIGN KEY(device_id) REFERENCES device(id)")

}

// Reverse the migrations
func (m *AirAd_20171201_164829) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
