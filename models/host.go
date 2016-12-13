package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Host struct {
	HostID          int       `orm:"column(id);auto"`
	CreateTime      time.Time `orm:"column(create_time);type(timestamp)"`
	HostName        string    `orm:"column(host_name);size(255);null"`
	InnerIP         string    `orm:"column(inner_ip);size(32)"`
	BgpIP           string    `orm:"column(bgp_ip);size(255);null"`
	OuterIP         string    `orm:"column(outer_ip);size(32);null"`
	IloIP           string    `orm:"column(ilo_ip);size(255);null"`
	Source          int8      `orm:"column(source)"`
	ModuleID        int       `orm:"column(module_id);null"`
	ModuleName      string    `orm:"column(module_name);size(255);null"`
	SetID           int       `orm:"column(set_id);null"`
	SetName         string    `orm:"column(set_name);size(255);null"`
	ApplicationID   int       `orm:"column(application_id);null"`
	ApplicationName string    `orm:"column(application_name);size(255);null"`
	Owner           string    `orm:"column(owner);size(255);null"`
	Checked         string    `orm:"column(checked);size(255);null"`
	IsDistributed   bool      `orm:"column(is_distributed)"`
}

func (t *Host) TableName() string {
	return "host"
}

func init() {
	orm.RegisterModel(new(Host))
}

// AddHost insert a new Host into database and returns
// last inserted Id on success.
func AddHost(hosts []*Host) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	_, err = o.InsertMulti(len(hosts), hosts)

	if err == nil {
		o.Commit()
	} else {
		o.Rollback()
	}
	return
}

// GetHostById retrieves Host by Id. Returns error if
// Id doesn't exist
func GetHostById(id int) (v *Host, err error) {
	o := orm.NewOrm()
	v = &Host{HostID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetHostById retrieves Host by Id. Returns error if
// Id doesn't exist
func GetHostByInnerIp(inner_ip string) bool {
	o := orm.NewOrm()
	return o.QueryTable("host").Filter("inner_ip", inner_ip).Exist()
}

// GetAllHost retrieves all Host matches certain condition. Returns empty list if
// no records exist
func GetAllHost(query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Host))
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

	var l []Host
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateHost updates Host by Id and returns error if
// the record to be updated doesn't exist
func UpdateHostById(m *Host) (err error) {
	o := orm.NewOrm()
	v := Host{HostID: m.HostID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 分配主机
func UpdateHostToApp(ids []int, appID int) (num int64, err error) {
	var app *App
	var set Set
	var mod Module
	o := orm.NewOrm()
	if app, err = GetAppById(appID); err != nil {
		return
	}

	if err = o.QueryTable("set").Filter("ApplicationID", appID).Filter("SetName", "空闲机池").One(&set); err != nil {
		return
	}

	if err = o.QueryTable("module").Filter("ApplicationId", appID).Filter("SetId", set.SetID).Filter("ModuleName", "空闲机").One(&mod); err != nil {
		return
	}

	num, err = o.QueryTable("host").Filter("HostID__in", ids).Update(orm.Params{
		"ApplicationID":   appID,
		"ApplicationName": app.ApplicationName,
		"SetID":           set.SetID,
		//		"SetName": set.SetName,
		"ModuleID":      mod.Id,
		"IsDistributed": true,
	})
	return
}

// 上交主机
func ResHostModule(ids []int, appID int) (num int64, err error) {
	var set Set
	var mod Module
	o := orm.NewOrm()
	if err = o.QueryTable("set").Filter("ApplicationID", appID).Filter("SetName", "空闲机池").One(&set); err != nil {
		return
	}

	if err = o.QueryTable("module").Filter("ApplicationId", appID).Filter("SetId", set.SetID).One(&mod); err != nil {
		return
	}

	num, err = o.QueryTable("host").Filter("HostID__in", ids).Update(orm.Params{
		"ApplicationID":   appID,
		"ApplicationName": "资源池",
		"SetID":           set.SetID,
		//		"SetName": set.SetName,
		"ModuleID":      mod.Id,
		"IsDistributed": false,
	})
	return
}

// DeleteHost deletes Host by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHost(id int) (err error) {
	o := orm.NewOrm()
	v := Host{HostID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Host{HostID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func DeleteHosts(id []int) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("host").Filter("HostID__in", id).Delete()
	return
}

func GetHostCount(id int, field string) (cnt int64, err error) {
	o := orm.NewOrm()
	if field == "ApplicationID" {
		cnt, err = o.QueryTable("host").Exclude("ModuleName", "空闲机").Filter(field, id).Count()
	} else {
		cnt, err = o.QueryTable("host").Filter(field, id).Count()
	}
	return
}

// 转移主机
func ModHostModule(appID int, moduleID int, hostIds []int) (num int64, err error) {
	o := orm.NewOrm()

	var app *App
	var set *Set
	var m *Module

	if m, err = GetModuleById(moduleID); err != nil {
		return
	}

	if appID != m.ApplicationId {
		err = errors.New("不能跨业务进行主机转移")
	}

	if app, err = GetAppById(appID); err != nil {
		return
	}

	if set, err = GetSetById(m.SetId); err != nil {
		return
	}

	num, err = o.QueryTable("host").Filter("HostID__in", hostIds).Update(orm.Params{
		"ApplicationID":   appID,
		"ApplicationName": app.ApplicationName,
		"SetID":           set.SetID,
		"SetName":         set.SetName,
		"ModuleID":        moduleID,
		"ModuleName":      m.ModuleName,
		"IsDistributed":   true,
	})
	return
}

// 移至空闲机/故障机
func DelHostModule(appID int, moduleName string, hostIds []int) (num int64, err error) {
	o := orm.NewOrm()
	var app *App
	var set Set
	var m Module

	if app, err = GetAppById(appID); err != nil {
		return
	}

	if set, err = GetDesetidByAppId(appID); err != nil {
		return
	}

	if err = o.QueryTable("module").Filter("ApplicationId", appID).Filter("SetId", set.SetID).Filter("ModuleName", moduleName).One(&m); err != nil {
		return
	}

	num, err = o.QueryTable("host").Filter("HostID__in", hostIds).Update(orm.Params{
		"ApplicationID":   appID,
		"ApplicationName": app.ApplicationName,
		"SetID":           set.SetID,
		"SetName":         set.SetName,
		"ModuleID":        m.Id,
		"ModuleName":      m.ModuleName,
	})
	return
}
