package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Mqtt struct {
	Id int `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment"`
	DeviceId int `json:"user_id" orm:"column(device_id);size(11)"`
	MqttOptionId int `json:"mqtt_option_id" orm:"column(mqtt_option_id);size(11)"`
	// QoS is the QoS of the fixed header.
	QoS byte `json:"qos, omitempty" orm:"column(qos);varbinary"`
	// Retain is the Retain of the fixed header.
	Retain bool `json:"retain, omitempty" orm:"column(retain);varbinary"`
	// TopicName is the Topic Name of the variable header.
	TopicName string `json:"topic_name, omitempty" orm:"column(topic_name);varbinary"`
	// Message is the Application Message of the payload.
	//Message []byte
	// TopicFilter is the Topic Filter of the Subscription.
	TopicFilter string `json:"topic_filter, omitempty" orm:"column(topic_filter);varbinary"`
}

func init() {
	orm.RegisterModel(new(Mqtt))
}

func Mqtts() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Mqtt))
}

// AddMqtt insert a new Mqtt into database and returns
// last inserted Id on success.
func AddMqtt(m *Mqtt) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMqttById retrieves Mqtt by Id. Returns error if
// Id doesn't exist
func GetMqttById(id int) (v *Mqtt, err error) {
	o := orm.NewOrm()
	v = &Mqtt{Id: id}
	if err = o.QueryTable(new(Mqtt)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMqtt retrieves all Mqtt matches certain condition. Returns empty list if
// no records exist
func GetAllMqtt(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Mqtt))
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

	var l []Mqtt
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

// UpdateMqtt updates Mqtt by Id and returns error if
// the record to be updated doesn't exist
func UpdateMqttById(m *Mqtt) (err error) {
	o := orm.NewOrm()
	v := Mqtt{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMqtt deletes Mqtt by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMqtt(id int) (err error) {
	o := orm.NewOrm()
	v := Mqtt{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Mqtt{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
