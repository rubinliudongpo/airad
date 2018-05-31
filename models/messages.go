package models

import (
	//"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id       int64 `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment"`
	UserId   int64 `json:"user_id" orm:"column(user_id);size(11)"`
	ToUserId int64 `json:"user_id" orm:"column(user_id);size(11)"`
	Type     int `json:"type, omitempty" orm:"column(type);size(4)"`
	SubType  int `json:"sub_type, omitempty" orm:"column(sub_type);size(4)"`
	Title    string `json:"title, omitempty" orm:"column(title);varbinary"`
	Url      string `json:"title, omitempty" orm:"column(url);varbinary"`
	Viewed   int `json:"viewed" orm:"column(viewed);size(1)"`// 1: viewed, 0:not-viewed
	CreatedAt  int64 `json:"created_at, omitempty" orm:"column(created_at);size(11)"`
}

func (this *Message) TableName() string {
	return TableName("message")
}
func init() {
	orm.RegisterModel(new(Message))
}

func Messages() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Message))
}

func AddMessages(m * Message)(id int64, err error) {
	o := orm.NewOrm()

	message := Message{
		UserId: m.UserId,
		ToUserId: m.ToUserId,
		Type: m.Type,
		SubType: m.SubType,
		CreatedAt: time.Now().UTC().Unix(),
		Title: m.Title,
		Viewed: 0,
		Url: m.Url,
	}

	//var id int64
	id, err = o.Insert(&message)
	if err == nil {
		return id, err
	}

	return 0, err
}

func ListMessages(query map[string]string, page int, offset int) (msg []interface{}, totalCount int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
	cond := orm.NewCondition()

	if query["toUserId"] != "" {
		cond = cond.And("ToUserId", query["toUserId"])
	}
	if query["View"] != "" {
		cond = cond.And("View", query["View"])
	}
	if query["Type"] != "" {
		cond = cond.And("Type", query["type"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageOffset")
	}
	start := (page - 1) * offset


	if totalCount, err := qs.Limit(offset, start).All(&msg); err == nil {
		return msg, totalCount, nil
	}
	return nil, 0, err
}

//统计数量
func CountMessages(query map[string]string) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("message"))
	cond := orm.NewCondition()
	if query["toUserId"] != "" {
		cond = cond.And("ToUserId", query["toUserId"])
	}
	if query["View"] != "" {
		cond = cond.And("View", query["view"])
	}
	if query["Type"] != "" {
		cond = cond.And("Type", query["type"])
	}
	if num, err := qs.SetCond(cond).Count(); err == nil {
		return num, nil
	}
	return 0, err
}

func ChangeMessageStatus(id int64, viewed int) (err error) {
	o := orm.NewOrm()
	v := Message{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		if _, err = o.Update(Message{Viewed:viewed}); err == nil {
			return nil
		}
	}
	return err
}

func ChangeMessageStatusAll(toUserId int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE "+ TableName("message") + " SET Viewed=1 WHERE ToUserId=? AND Viewed=0", toUserId).Exec()
	return err
}

func DeleteMessages(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + TableName("message") + " WHERE id IN(" + ids + ")").Exec()
	return err
}