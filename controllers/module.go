package controllers

import (
	"encoding/json"
	"errors"
	"github.com/shwinpiocess/cc/models"
	"strconv"
	"strings"
)

// oprations for Module
type ModuleController struct {
	BaseController
}

func (this *ModuleController) NewModule() {
	out := make(map[string]interface{})
	
	m := new(models.Module)
	m.ApplicationId, _ = this.GetInt("ApplicationID")
	m.SetId, _ = this.GetInt("SetID")
	m.ModuleName = this.GetString("ModuleName")
	m.Operator, _ = this.GetInt("Operator")
	m.BakOperator, _ = this.GetInt("BakOperator")
	m.Owner = this.userId
	
	if _, err := models.AddModule(m); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = "重复的模块名！"
		this.jsonResult(out)
	}
}

// @Title Get
// @Description get Module by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Module
// @Failure 403 :id is empty
// @router /:id [get]
func (this *ModuleController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetModuleById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJSON()
}

// @Title Get All
// @Description get Module
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Module
// @Failure 403
// @router / [get]
func (this *ModuleController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllModule(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJSON()
}

// @Title Update
// @Description update the Module
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Module	true		"body for Module content"
// @Success 200 {object} models.Module
// @Failure 403 :id is not int
// @router /:id [put]
func (this *ModuleController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Module{Id: id}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateModuleById(&v); err == nil {
			this.Data["json"] = "OK"
		} else {
			this.Data["json"] = err.Error()
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the Module
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ModuleController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteModule(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
