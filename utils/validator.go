package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func CheckUsernamePassword(username string, password string) (errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	valid.Required(username, "Username").Message("用户名必填")
	valid.Required(password, "Password").Message("密码必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}

func CheckNewUserPost(Username string, Password string, Age int,
	Gender int, Address string, Email string) (errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	valid.Required(Username, "Username").Message("用户名必填")
	valid.AlphaNumeric(Username, "Username").Message("用户名必须是数字或字符")
	valid.Required(Password, "Password").Message("密码必填")
	valid.MinSize(Password, 6,"Password").Message("密码不能少于6位")
	valid.MaxSize(Password, 20,"Password").Message("密码不能多于20位")
	valid.Range(Age, 1,100, "Age").Message("年龄在1到100岁")
	valid.Range(Gender, 0,2, "Gender").Message("性别不正确")
	valid.Required(Address, "Address").Message("地址必填")
	valid.Email(Email, "Email").Message("邮箱必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}

