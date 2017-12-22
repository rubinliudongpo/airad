package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"reflect"
	"fmt"
	"errors"
	"time"
)

type Device struct {
	Id int `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment"`
	UserId int `json:"user_id" orm:"column(user_id);size(11)"`
	DeviceName string `json:"device_name" orm:"column(device_name);unique;size(32)"`
	Address string `json:"address" orm:"column(address);size(50)"`
	Status int `json:"status" orm:"column(status);size(1)"`// 0: enabled, 1:disabled
	CreatedAt int64 `json:"created_at, omitempty" orm:"column(created_at);size(11)"`
	UpdatedAt int64 `json:"updated_at, omitempty" orm:"column(updated_at);size(11)"`
	Latitude string `json:"latitude, omitempty" orm:"column(latitude);size(12)"`
	Longitude string `json:"longitude, omitempty" orm:"column(longitude);size(12)"`
	AirAdCount int64 `json:"airad_count, omitempty" orm:"column(airad_count);size(64)"`
	//User *User `json:"user_id" orm:"rel(fk)"`
	//AirAd []*AirAd `orm:"reverse(many)"` // 设置一对多的反向关系
}

type DeviceRequestStruct struct {
	UserId int `json:"userId"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
    Fields string `json:"fields"`
}

func init() {
	orm.RegisterModel(new(Device))
}

func Devices() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Device))
}

// AddDevice insert a new Device into database and returns
// last inserted Id on success.
func AddDevice(m *Device) (id int64, err error) {
	o := orm.NewOrm()

	CreatedAt := time.Now().UTC().Unix()
	UpdatedAt := CreatedAt

	device := Device{
		DeviceName:m.DeviceName,
		UserId:m.UserId,
		Address:m.Address,
		Status:m.Status,
		CreatedAt:CreatedAt,
		UpdatedAt:UpdatedAt,
		Latitude:m.Latitude,
		Longitude:m.Longitude,
		//User *User `json:"user_id" orm:"rel(fk)"`
		//AirAd []*AirAd `orm:"reverse(many)"`
	}

	//var id int64
	id, err = o.Insert(&device)
	if err == nil{
		return id, err
	}

	return 0, err
}

// 检测DeviceName是否存在
func CheckDeviceName(deviceName string) bool {
	exist := Devices().Filter("DeviceName", deviceName).Exist()
	return exist
}

// GetDeviceById retrieves Device by Id. Returns error if
// Id doesn't exist
func GetDeviceById(id int) (v *Device, err error) {
	o := orm.NewOrm()
	v = &Device{Id: id}
	if err = o.QueryTable(new(Device)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetDeviceByUser retrieves Device by User. Returns error if
// Id doesn't exist
func GetDevicesByUserId(userId int, fields []string, limit int, offset int) (devices []*Device, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Device)).Filter("user_id", userId).Limit(limit, offset).All(&devices, fields...); err == nil {
		return devices, nil
	}
	return nil, err
}

// GetAllDevices retrieves all Device matches certain condition. Returns empty list if
// no records exist
func GetAllDevices(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int, userId int) (ml []interface{}, totalCount int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Device))
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Device
	qs = qs.OrderBy(sortFields...).RelatedSel()
	totalCount, err = qs.Filter("UserId", userId).Count()
	if _, err = qs.Filter("UserId", userId).Limit(limit, offset).All(&l, fields...); err == nil {
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
		return ml, totalCount, nil
	}
	return nil, 0, err
}

// UpdateDevice updates Device by Id and returns error if
// the record to be updated doesn't exist
func UpdateDeviceById(m *Device) (err error) {
	o := orm.NewOrm()
	v := Device{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateDevice updates Device by AirAdCount and returns error if
// the record to be updated doesn't exist
func UpdateDeviceAirAdCount(m *Device) (err error) {
	o := orm.NewOrm()
	v := Device{Id: m.Id}
	m.AirAdCount += 1
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}


// DeleteDevice deletes Device by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDevice(id int) (err error) {
	o := orm.NewOrm()
	v := Device{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Device{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}