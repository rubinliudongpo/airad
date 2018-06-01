package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"fmt"
	"github.com/rubinliudongpo/airad/utils"
	"time"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego"
)

func (u *User) TableName() string {
	return TableName("user")
}

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id int `json:"id" orm:"column(id);pk;unique;auto_increment;int(11)"`
	Username string `json:"username" orm:"column(username);unique;size(32)"`
	Password string `json:"password" orm:"column(password);size(128)"`
	Avatar string `json:"avatar, omitempty" orm:"column(avatar);varbinary"`
	Salt string `json:"salt" orm:"column(salt);size(128)"`
	Token string `json:"token" orm:"column(token);size(256)"`
	Gender int `json:"gender" orm:"column(gender);size(1)"`  // 0:Male, 1: Female, 2: undefined
	Age int `json:"age" orm:"column(age):size(3)"`
	Address string `json:"address" orm:"column(address);size(50)"`
	Email string `json:"email" orm:"column(email);size(50)"`
	LastLogin int64 `json:"last_login" orm:"column(last_login);size(11)"`
	Status int `json:"status" orm:"column(status);size(1)"`// 0: enabled, 1:disabled
	CreatedAt int64 `json:"created_at" orm:"column(created_at);size(11)"`
	UpdatedAt int64 `json:"updated_at" orm:"column(updated_at);size(11)"`
	DeviceCount int `json:"device_count" orm:"column(device_count);size(64);default(0)"`
	//Device []*Device `orm:"reverse(many)"` // 设置一对多的反向关系
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测用户是否存在
func CheckUserId(userId int) bool {
	exist := Users().Filter("Id", userId).Exist()
	return exist
}

// 检测用户是否存在
func CheckUserName(username string) bool {
	exist := Users().Filter("Username", username).Exist()
	return exist
}

// 检测用户是否存在
func CheckUserIdAndToken(userId int, token string) bool {
	exist := Users().Filter("Id", userId).Filter("Token", token).Exist()
	return exist
}


// 检测用户是否存在
func CheckEmail(email string) bool {
	exist := Users().Filter("Email", email).Exist()
	return exist
}

// CheckPass compare input password.
func (u *User) CheckPassword(password string) (ok bool, err error) {
	hash, err := utils.GeneratePassHash(password, u.Salt)
	if err != nil {
		return false, err
	}

	return u.Password == hash, nil
}

// 根据用户ID获取用户
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}


// 根据用户名字获取用户
func GetUserByUserName(username string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Username: username}
	if err = o.QueryTable(new(User)).Filter("Username", username).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

func GetUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(username string, password string) (bool, *User) {
	o := orm.NewOrm()
	user, err := GetUserByUserName(username)
	if err != nil {
		return false, nil
	}
	passwordHash, err := utils.GeneratePassHash(password, user.Salt)
	if err != nil {
		return false, nil
	}
	err = o.QueryTable(user).Filter("Username", username).Filter("Password", passwordHash).One(user)
	return err != orm.ErrNoRows, user

}

func GetUserByUsername(username string) (err error, user *User) {
	o := orm.NewOrm()
	user = &User{Username: username}
	if err := o.QueryTable(user).Filter("Username", username).One(&user); err == nil {
		return nil, user
	}
	return err, nil
}

func AddUser(m *User) (*User, error) {
	o := orm.NewOrm()
	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}

	passwordHash, err := utils.GeneratePassHash(m.Password, salt)
	if err != nil {
		return nil, err
	}
	CreatedAt := time.Now().UTC().Unix()
	UpdatedAt := CreatedAt
	LastLogin := CreatedAt

	et := utils.EasyToken{
		Username: m.Username,
		Uid: 0,
		Expires:  time.Now().Unix() + 2 * 3600,
	}
	token, err := et.GetToken()
	user := User{
		Username:m.Username,
		Password: passwordHash,
		Salt:salt,
		Token:token,
		Gender:m.Gender,
		Age:m.Age,
		Address:m.Address,
		Email:m.Email,
		LastLogin:LastLogin,
		Status:m.Status,
		CreatedAt:CreatedAt,
		UpdatedAt:UpdatedAt,
	}
	_, err = o.Insert(&user)
	if err == nil{
		return &user, err
	}

	return nil, err
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

// UpdateDevice updates User by DeviceCount and returns error if
// the record to be updated doesn't exist
func UpdateUserDeviceCount(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.DeviceCount += 1
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// updates User's Token and returns error if
// the record to be updated doesn't exist
func UpdateUserToken(m *User, token string) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.Token = token
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// updates User's LastLogin and returns error if
// the record to be updated doesn't exist
func UpdateUserLastLogin(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	lastLogin := time.Now().UTC().Unix()
	m.LastLogin = lastLogin
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AirAd{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func GetUsername(id int) string {
	var err error
	var username string

	err = utils.GetCache("GetUsername.id."+fmt.Sprintf("%d", id), &username)
	if err != nil {
		cacheExpire, _ := beego.AppConfig.Int("cache_expire")
		var user User
		o := orm.NewOrm()
		o.QueryTable(TableName("user")).Filter("Id", id).One(&user, "username")
		username = user.Username
		utils.SetCache("GetRealname.id."+fmt.Sprintf("%d", id), username, cacheExpire)
	}
	return username
}

//func HashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes), err
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}



//func generateToken() (tokenString string, err error) {
//	/* Create the token */
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	/* Create a map to store our claims
//	claims := token.Claims.(jwt.MapClaims)
//
//	/* Set token claims */
//	claims["admin"] = true
//	claims["name"] = "Ado Kukic"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	/* Sign the token with our secret */
//	tokenString, _ := token.SignedString(mySigningKey)
//}