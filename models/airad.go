package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	"time"
)

type AirAd struct {
	Id int `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment"`
	DeviceId int `json:"device_id" orm:"column(device_id);size(11)"`
	CreatedAt int64 `json:"created_at, omitempty" orm:"column(created_at);size(11)"`
	Nh3 string `json:"nh3, omitempty" orm:"column(nh3);size(4)"`
	Co string `json:"co, omitempty" orm:"column(co);size(4)"`
	O3 string `json:"o3, omitempty" orm:"column(o3);size(4)"`
	Pm25 string `json:"pm25, omitempty" orm:"column(pm25);size(4)"`
	Pm10 string `json:"pm10, omitempty" orm:"column(pm10);size(4)"`
	So2 string `json:"so2, omitempty" orm:"column(so2);size(4)"`
	Temperature string `json:"temperature, omitempty" orm:"column(temperature);size(4)"`
	Humidity string `json:"humidity, omitempty" orm:"column(humidity);size(4)"`
	AqiQuality string `json:"aqi_quality, omitempty" orm:"column(aqi_quality);size(4)"`
	Suggest string `json:"suggest, omitempty" orm:"column(suggest);size(4)"`
	//Device *Device `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(AirAd))
}

func AirAds() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(AirAd))
}

// AddAirAd insert a new AirAd into database and returns
// last inserted Id on success.
func AddAirAd(m *AirAd) (id int64, err error) {
	o := orm.NewOrm()

	CreatedAt := time.Now().UTC().Unix()

	airAd := AirAd{
		DeviceId:m.DeviceId,
		Nh3:m.Nh3,
		Pm10:m.Pm10,
		Pm25:m.Pm25,
		Co:m.Co,
		O3:m.O3,
		So2:m.So2,
		Temperature:m.Temperature,
		Humidity:m.Humidity,
		AqiQuality:m.AqiQuality,
		Suggest:m.Suggest,
		CreatedAt:CreatedAt,
	}

	id, err = o.Insert(&airAd)
	if err == nil{
		return id, err
	}

	return 0, err
}

// GetAirAdById retrieves AirAd by Id. Returns error if
// Id doesn't exist
func GetAirAdById(id int) (v *AirAd, err error) {
	o := orm.NewOrm()
	v = &AirAd{Id: id}
	if err = o.QueryTable(new(AirAd)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 检测DeviceId是否存在
func CheckDeviceId(deviceId int) bool {
	exist := Devices().Filter("Id", deviceId).Exist()
	return exist
}

// GetAllAirAds retrieves all AirAd matches certain condition. Returns empty list if
// no records exist
func GetAllAirAds(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int, userId int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AirAd))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AirAd
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAirAd updates AirAd by Id and returns error if
// the record to be updated doesn't exist
func UpdateAirAdById(m *AirAd) (err error) {
	o := orm.NewOrm()
	v := AirAd{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAirAd deletes AirAd by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAirAd(id int) (err error) {
	o := orm.NewOrm()
	v := AirAd{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AirAd{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}