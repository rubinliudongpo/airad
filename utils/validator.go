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
	valid.Required(Age, "Age").Message("年龄必填")
	valid.Range(Age, 1,100, "Age").Message("年龄在1到100岁")
	valid.Range(Gender, 0,2, "Gender").Message("性别不正确")
	valid.Required(Address, "Address").Message("地址必填")
	valid.MinSize(Address, 6,"Address").Message("地址不能少于6位")
	valid.MaxSize(Address, 50,"Address").Message("地址不能多于50位")
	valid.Required(Email, "Email").Message("邮箱必填")
	valid.Email(Email, "Email").Message("邮箱格式不正确")

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

func CheckNewDevicePost(userId int, deviceName string, address string,
	status int, latitude string, longitude string) (errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	valid.Required(userId, "UserId").Message("用户ID必填")
	valid.Required(deviceName, "DeviceName").Message("设备名必填")
	valid.AlphaNumeric(deviceName, "DeviceName").Message("设备名必须是数字或字符")
	valid.Required(address, "Address").Message("地址必填")
	valid.MinSize(address, 6,"Address").Message("地址不能少于6位")
	valid.MaxSize(address, 50,"Address").Message("地址不能多于50位")
	valid.Range(status, 0, 1 ,"Status").Message("状态必填")
	valid.Required(latitude, "Latitude").Message("纬度必填")
	valid.Required(longitude, "Longitude").Message("经度必填")
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

func CheckUserDevicePost(userId int, limit int, offset int) (errorMessage string){
	valid := validation.Validation{}
	valid.Required(userId, "UserId").Message("用户ID必填")
	valid.Min(userId, 1, "UserId").Message("用户ID必须是数字")
	//valid.Required(limit, "Limit").Message("Limit必填")
	valid.Range(limit, 0, 20,"Limit").Message("Limit必须是数字")
	//valid.Required(offset, "Offset").Message("Offset必填")
	valid.Range(offset, 0, 20,"Offset").Message("Offset必须是数字")
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

func CheckNewAirAdPost(deviceId int, co string, humidity string, temperature string,
	pm25 string, pm10 string, nh3 string, o3 string, suggest string, aqiQuality string)(errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	valid.Required(deviceId, "DeviceId").Message("用户ID必填")
	valid.Required(co, "Co").Message("设备名必填")
	valid.Required(humidity, "Humidity").Message("地址必填")
	valid.Required(temperature, "DeviceId").Message("用户ID必填")
	valid.Required(pm25, "Co").Message("设备名必填")
	valid.Required(pm10, "Humidity").Message("地址必填")
	valid.Required(o3, "DeviceId").Message("用户ID必填")
	valid.Required(suggest, "Co").Message("设备名必填")
	valid.Required(aqiQuality, "Humidity").Message("地址必填")
	valid.Required(nh3, "DeviceId").Message("用户ID必填")

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