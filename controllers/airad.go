package controllers

import (
	//"github.com/astaxie/beego"
	"github.com/rubinliudongpo/airad/models"
	"strings"
	"errors"
	"encoding/json"
	"strconv"
	"github.com/rubinliudongpo/airad/utils"
	"fmt"
)

// AirAdController operations for AirAd
type AirAdController struct {
	BaseController
}

// URLMapping ...
func (c *AirAdController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create AirAd
// @Param	body		body 	models.AirAd	true		"body for AirAd content"
// @Success 201 {object} models.AirAd
// @Failure 403 body is empty
// @router / [post]
func (c *AirAdController) Post() {
	var v models.AirAd

	token := c.Ctx.Input.Header("token")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(token)
	if !valido {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := utils.CheckNewAirAdPost(v.DeviceId, v.Co, v.Humidity, v.Temperature,
			v.Pm25, v.Pm10, v.Nh3, v.O3, v.Suggest, v.AqiQuality); errorMessage != "ok"{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,errorMessage, ""}
			c.ServeJSON()
			return
		}
		if !models.CheckDeviceId(v.DeviceId){
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"设备Id不存在", ""}
			c.ServeJSON()
			return
		}

		if !models.CheckDeviceIdAndToken(v.DeviceId, token){
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"用户ID和Token不匹配", ""}
			c.ServeJSON()
			return
		}


		if airAdId, err := models.AddAirAd(&v); err == nil {
			if device, err := models.GetDeviceById(v.DeviceId); err == nil {
				models.UpdateDeviceAirAdCount(device)
				c.Ctx.Output.SetStatus(201)
				var returnData = &CreateObjectData{int(airAdId)}
				c.Data["json"] = &Response{0, 0, "ok", returnData}
			} else {
				c.Ctx.ResponseWriter.WriteHeader(403)
				c.Data["json"] = Response{403, 403,"设备Id不存在", ""}
				c.ServeJSON()
				return
			}

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get AirAd by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.AirAd
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AirAdController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAirAdById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get AirAd
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AirAd
// @Failure 403
// @router / [get]
func (c *AirAdController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 20
	var offset int
	var deviceId int

	deviceId, err := strconv.Atoi(c.Ctx.Input.Header("device_id"))
	token := c.Ctx.Input.Header("token")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	found, _ := models.GetUserByToken(token)
	if  !found {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = &Response{401, 401, "未找到相关的用户", ""}
		c.ServeJSON()
		return
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, totalCount, err := models.GetAllAirAds(query, fields, sortby, order, offset, limit, deviceId)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var returnData = &GetAirAdData{totalCount, l}
		c.Data["json"] = &Response{0, 0, "ok", returnData}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the AirAd
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AirAd	true		"body for AirAd content"
// @Success 200 {object} models.AirAd
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AirAdController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.AirAd{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAirAdById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the AirAd
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AirAdController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAirAd(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
