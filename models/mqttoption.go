package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type MqttOption struct {
	Id int `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment;int(11)"`
	// ClientID is the Client Identifier of the payload.
	ClientID string `json:"client_id, omitempty" orm:"column(client_id);varbinary"`
	// CONNACKTimeout is timeout in seconds for the Client
	// to wait for receiving the CONNACK Packet after sending
	// the CONNECT Packet.
	ConnAckTimeout time.Duration `json:"conn_ack_timeout, omitempty" orm:"column(conn_ack_timeout);size(64)"`
	// UserName is the User Name of the payload.
	UserName string `json:"user_name, omitempty" orm:"column(user_name);size(32)"`
	// Password is the Password of the payload.
	Password  string `json:"password, omitempty" orm:"column(password);size(128)"`
	// CleanSession is the Clean Session of the variable header.
	CleanSession bool `json:"clean_session, omitempty" orm:"column(clean_session)"`
	// KeepAlive is the Keep Alive of the variable header.
	KeepAlive uint16 `json:"keep_alive, omitempty" orm:"column(keep_alive)"`
	// WillTopic is the Will Topic of the payload.
	WillTopic string `json:"will_topic, omitempty" orm:"column(will_topic);varbinary"`
	// WillMessage is the Will Message of the payload.
	WillMessage string `json:"will_message, omitempty" orm:"column(will_message);varbinary"`
	// WillQoS is the Will QoS of the variable header.
	WillQoS bool `json:"will_qos, omitempty" orm:"column(will_qos)"`
	// WillRetain is the Will Retain of the variable header.
	WillRetain bool `json:"will_retain, omitempty" orm:"column(will_retain)"`
}

func init() {
	orm.RegisterModel(new(MqttOption))
}

func MqttOptions() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(MqttOption))
}

// AddMqttOption insert a new MqttOption into database and returns
// last inserted Id on success.
func AddMqttOption(m *MqttOption) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMqttOptionById retrieves MqttOption by Id. Returns error if
// Id doesn't exist
func GetMqttOptionById(id int) (v *MqttOption, err error) {
	o := orm.NewOrm()
	v = &MqttOption{Id: id}
	if err = o.QueryTable(new(MqttOption)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMqttOption retrieves all MqttOption matches certain condition. Returns empty list if
// no records exist
func GetAllMqttOption(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MqttOption))
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

	var l []MqttOption
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

// UpdateMqttOption updates MqttOption by Id and returns error if
// the record to be updated doesn't exist
func UpdateMqttOptionById(m *MqttOption) (err error) {
	o := orm.NewOrm()
	v := MqttOption{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMqttOption deletes MqttOption by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMqttOption(id int) (err error) {
	o := orm.NewOrm()
	v := MqttOption{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MqttOption{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
