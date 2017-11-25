package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AirAd_20171109_164549 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AirAd_20171109_164549{}
	m.Created = "20171109_164549"

	migration.Register("AirAd_20171109_164549", m)
}

// Run the migrations
func (m *AirAd_20171109_164549) Up() {
	m.SQL("ALTER TABLE air_ad ADD createdAt int, ADD nh3 varchar(10), " +
		"ADD co varchar(10), ADD pm25 varchar(10), ADD pm10 varchar(10), ADD o3 varchar(10)," +
			"ADD so2 varchar(10)")

}

// Reverse the migrations
func (m *AirAd_20171109_164549) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
