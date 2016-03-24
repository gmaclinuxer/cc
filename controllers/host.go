package controllers

import (
	"path"
	"fmt"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
	
	"github.com/shwinpiocess/cc/models"
)

// oprations for Host
type HostController struct {
	BaseController
}


func (this *HostController) ImportPrivateHostByExcel() {
	out := make(map[string]interface{})
	fmt.Println("jing...............................")
	f, h, err := this.GetFile("importPrivateHost")
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxjjjjjjjjjjjjjjjjjjjjjj f", f)
	if f == nil {
		out["success"] = false
		out["message"] = "请先提供后缀名为xlsx的Excel文件再进行上传操作！"
		out["name"] = "importToCC"
		goto render
	}
	defer f.Close()
	fmt.Println("uuuuuuuuuuuuuuuuuuuuuulllllllllllllllllllll")
	if err != nil {
		fmt.Println("111111111111111111111111111")
		out["success"] = false
		out["message"] = "主机导入失败！上传文件不合法！"
		out["name"] = "importToCC"
		goto render
	} else {
		if suffix := path.Ext(h.Filename); suffix != ".xlsx" {
			out["success"] = false
			out["message"] = "请提供后缀名为xlsx的Excel文件！"
			out["name"] = "importToCC"
			goto render
		}
		excelFileName := "static/upload/" + h.Filename
		this.SaveToFile("importPrivateHost", excelFileName)
		
		if xlFile, err := xlsx.OpenFile(excelFileName); err == nil {
			fmt.Println("222222222222222222222222222")
			var hosts []*models.Host
			var ips string
			var sns string
			
			for _, sheet := range xlFile.Sheets {
				for index, row := range sheet.Rows {
					fmt.Println("len(row.cells", len(row.Cells))
					if index > 0 && len(row.Cells) >= 6 {
						fmt.Println("3333333333333333333333333")
						Sn, err1 := row.Cells[0].Int64()
						Hostname, err2 := row.Cells[1].String()
						InnerIp, err3 := row.Cells[2].String()
						OuterIp, err4 := row.Cells[3].String()
						Operator, err5 := row.Cells[4].String()
						OsName, err6 := row.Cells[5].String()
						if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
							fmt.Println("444444444444444444444")
							out["success"] = false
							out["message"] = "主机导入失败！上传文件内容格式不正确！"
							out["name"] = "importToCC"
							goto render
						} else {
							fmt.Println("5555555555555555555555555")
							if models.GetHostByInnerIp(InnerIp) {
								fmt.Println("err=", err)
								fmt.Println("6666666666666666666666")
								ips = ips + "<li>" + InnerIp + "</li>"
							}
							if models.GetHostBySn(Sn) {
								fmt.Println("7777777777777777777777777777777777")
								sns = sns + fmt.Sprintf("`SN`%d已存在</br>", Sn)
							}
							host := new(models.Host)
							host.SN = Sn
							host.HostName = Hostname
							host.InnerIP = InnerIp
							host.OuterIP = OuterIp
							host.Operator = Operator
							host.OSName = OsName
							host.Source = 3
							hosts = append(hosts, host)
						}
					}
				}
			}
			
			if ips != "" {
				fmt.Println("888888888888888888888888")
				out["success"] = false
				out["message"] = `有内网IP在私有云中已经存在,请先修改这些IP的平台再做导入,具体如下：<ul class="">` + ips + "</ul>"
				out["name"] = "importToCC"
				goto render
			}
			
			if sns != "" {
				fmt.Println("999999999999999999999999999999999999")
				out["success"] = false
				out["message"] = sns
				out["name"] = "importToCC"
				goto render
			}
			
			if err := models.AddHost(hosts); err == nil {
				fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa err=", err)
				out["success"] = true
				out["message"] = "导入成功！"
				out["name"] = "importToCC"
			} else {
				fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbberr=", err)
				out["success"] = false
				out["message"] = "主机导入数据库出现问题，请联系管理员！"
				out["name"] = "importToCC"
			}
		} else {
			fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccc")
			out["success"] = false
			out["message"] = "主机导入失败！上传文件格式不正确！"
			out["name"] = "importToCC"
		}
    }
	
	render:
	fmt.Println("ddddddddddddddddddddddddddddddddddddddddd")
	this.Data["result"] = out
	this.TplName = "host/upload.html"
}

// @Title Get
// @Description get Host by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Host
// @Failure 403 :id is empty
// @router /:id [get]
func (this *HostController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetHostById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJSON()
}

// @Title Get All
// @Description get Host
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Host
// @Failure 403
// @router / [get]
func (this *HostController) GetHost4QuickImport() {
	isDistributed, _ := this.GetBool("IsDistributed")
	source := this.GetString("Source")
	applicationId := this.GetString("ApplicationID")
	
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0
	
	query["is_distributed"] = strconv.FormatBool(isDistributed)
	query["source"] = source
	
	if isDistributed {
		query["application_id"] = applicationId
	}
	
	fmt.Println("query=", query)

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

	l, err := models.GetAllHost(query, fields, sortby, order, offset, limit)
	
	out := make(map[string]interface{})
	if err != nil {
		out["success"] = false
		this.jsonResult(out)
	} else {
		out["success"] = true
		out["data"] = l
		out["total"] = len(l)
		this.jsonResult(out)
	}
}

// @Title Update
// @Description update the Host
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Host	true		"body for Host content"
// @Success 200 {object} models.Host
// @Failure 403 :id is not int
// @router /:id [put]
func (this *HostController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Host{HostID: id}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateHostById(&v); err == nil {
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
// @Description delete the Host
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *HostController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteHost(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
