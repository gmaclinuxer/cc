package controllers

import (
    "fmt"
    "strings"
    "github.com/astaxie/beego"

	"github.com/shwinpiocess/cc/models"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) jsonResult(out interface{}) {
        this.Data["json"] = out
        this.ServeJSON()
        this.StopRun()
}

func (this *ApiController) GetCCModuleTree() {
    // 获取业务拓扑
    payload := make(map[string]interface{})

    applicationId, _ := this.GetInt("appId", 0)
    topo, err := models.GetCCModuleTree(applicationId)
    if err != nil {
        payload["success"] = false
        this.jsonResult(payload)
    } else {
        payload["data"] = topo
        payload["success"] = true
        this.jsonResult(payload)
    }
}

func (this *ApiController) GetCCHosts() {
    payload := make(map[string]interface{})

    moduleId, _ := this.GetInt("moduleId", 0)

    var fields []string
    var sortby []string
    var order []string
    var query = make(map[string]interface{})
    var limit int64 = 0
    var offset int64 = 0

    query["module_id"] = moduleId

    fmt.Println("moduleId", moduleId)

    data, err := models.GetAllHost(query, fields, sortby, order, offset, limit)

    fmt.Println(data)
    var items []interface{}
    for _, each := range data {
        item := make(map[string]interface{})
        item["source"] = each.(models.Host).Source
        item["alived"] = 1
        item["ipDesc"] = each.(models.Host).HostName
        item["ip"] = each.(models.Host).InnerIP
        items = append(items, item)
    }
    fmt.Println(err)
    payload["data"] = items
    payload["success"] = true
    this.jsonResult(payload)
}


func (this *ApiController) GetHostsByIPs() {
    // 通过内网IP批量获取主机名
    ips := this.GetString("ips", "")

    ipList := strings.Split(ips, ",")

    fmt.Println("要获取主机名的IP地址为", ips)
    var fields []string
    var sortby []string
    var order []string
    var query = make(map[string]interface{})
    var limit int64 = 0
    var offset int64 = 0

    query["inner_ip__in"] = ipList

    data, err := models.GetAllHost(query, fields, sortby, order, offset, limit)

    fmt.Println(data)
    fmt.Println(err)
    payload := make(map[string]string)

    for _, each := range data {
        payload[each.(models.Host).InnerIP] = each.(models.Host).HostName
    }

    fmt.Println(payload)
        
    this.jsonResult(payload)
}
