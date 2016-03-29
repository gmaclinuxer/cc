package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type App struct {
	Id              int    `orm:"column(id);auto"`
	Type            int8   `orm:"column(type)"`
	ApplicationName string `orm:"column(application_name);size(255)"`
	LifeCycle       string `orm:"column(life_cycle);size(255)"`
	Level           int8   `orm:"column(level)"`
	OwnerId         int    `orm:"column(owner_id)"`
	Default         bool   `orm:"column(default);null"`
}

func (t *App) TableName() string {
	return "app"
}

func init() {
	orm.RegisterModel(new(App))
}

// AddApp insert a new App into database and returns
// last inserted Id on success.
func AddApp(m *App) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddDefApp(userId int) {
	if ExistByName("资源池") {
		return
	}
	defApp := new(App)
	defApp.Type = 0
	defApp.Level = 3
	defApp.ApplicationName = "资源池"
	defApp.LifeCycle = "公测"
	defApp.OwnerId = userId
	defApp.Default = true

	o := orm.NewOrm()
	if id, err := o.Insert(defApp); err == nil {
		s := new(Set)
		s.ApplicationID = int(id)
		s.SetName = "空闲机池"
		s.EnviType = 3
		s.ServiceStatus = 1
		s.Default = true
		s.Owner = userId
		if setId, err := AddSet(s); err == nil {
			m := new(Module)
			m.ApplicationId = int(id)
			m.SetId = int(setId)
			m.ModuleName = "空闲机"
			m.Owner = userId
			AddModule(m)
		}
	}
	return
}

func GetDefAppByUserId(userId int) (info map[string]interface{}, err error) {
	var app App
	var set Set
	var mod Module

	info = make(map[string]interface{})
	o := orm.NewOrm()

	if err = o.QueryTable("app").Filter("OwnerId", userId).Filter("Default", true).One(&app); err != nil {
		return
	}
	
	if err = o.QueryTable("set").Filter("ApplicationID", app.Id).One(&set); err != nil {
		return
	}
	
	if err = o.QueryTable("module").Filter("ApplicationId", app.Id).Filter("SetId", set.SetID).One(&mod); err != nil {
		return
	}
	
	fmt.Println("iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii app=", app, "set=", set, "mod=", mod)
	fmt.Println("app.id", app.Id)
	info["AppId"] = app.Id
	info["AppName"] = app.ApplicationName
	info["SetId"] = set.SetID
	info["SetName"] = set.SetName
	info["ModuleId"] = mod.Id
	info["ModuleName"] = mod.ModuleName
	return
}

// GetAppById retrieves App by Id. Returns error if
// Id doesn't exist
func GetAppById(id int) (v *App, err error) {
	o := orm.NewOrm()
	v = &App{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func ExistByName(name string) bool {
	o := orm.NewOrm()
	return o.QueryTable("app").Filter("ApplicationName", name).Exist()
}

func GetAppTopoById(id int) (ml []interface{}, err error) {
	//	[{"id":"5524","text":"aaa","spriteCssClass":"c-icon icon-group","type":"set","expanded":false,"number":32,"items":[{"id":"7025","spriteCssClass":"c-icon icon-modal","text":"1","operator":"1842605324","bakoperator":"1842605324","type":"module","number":32}]}]
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	query["application_id"] = strconv.Itoa(id)
	app, _ := GetAppById(id)
	if app.Level == 3 {
		query["default"] = "0"
	} else {
		query["default"] = "1"
	}

	if sets, err := GetAllSet(query, fields, sortby, order, offset, limit); err == nil {
		for _, set := range sets {
			s := make(map[string]interface{})
			s["id"] = set.(Set).SetID
			s["text"] = set.(Set).SetName
			if app.Level == 3 {
				s["spriteCssClass"] = "c-icon icon-group"
				s["expanded"] = false
			} else {
				s["spriteCssClass"] = "c-icon icon-group hide"
				s["expanded"] = true
			}
			s["type"] = "set"

			s["number"] = 32
			var items []interface{}
			var fields []string
			var sortby []string
			var order []string
			var query map[string]string = make(map[string]string)
			var limit int64 = 0
			var offset int64 = 0

			query["application_id"] = strconv.Itoa(id)
			query["set_id"] = strconv.Itoa(set.(Set).SetID)
			if mods, err := GetAllModule(query, fields, sortby, order, offset, limit); err == nil {
				for _, mod := range mods {
					if mod.(Module).ModuleName != "空闲机" {
						m := make(map[string]interface{})
						m["id"] = mod.(Module).Id
						m["text"] = mod.(Module).ModuleName
						m["spriteCssClass"] = "c-icon icon-modal"
						m["type"] = "module"
						m["operator"] = mod.(Module).Operator
						m["bakoperator"] = mod.(Module).BakOperator
						m["number"] = 32
						items = append(items, m)
					}
				}
				s["items"] = items
			}
			ml = append(ml, s)
		}
	}

	return
}

func GetEmptyById(id int) (info map[string]interface{}, options map[int]string, err error) {
	info = make(map[string]interface{})
	options = make(map[int]string)
	var topo []interface{}
	var setItems []interface{}
	var emptyItems []interface{}
	var app *App
	var sets []interface{}
	var mods []interface{}

	setItems = append(setItems,1)[:0]
	// 1. 获取业务信息
	if app, err = GetAppById(id); err != nil {
		return
	}

	// 2.获取业务下集群信息，2级结构需要特殊处理
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	query["application_id"] = strconv.Itoa(id)

	if sets, err = GetAllSet(query, fields, sortby, order, offset, limit); err != nil {
		return
	}
	query["application_id"] = strconv.Itoa(id)

	if mods, err = GetAllModule(query, fields, sortby, order, offset, limit); err != nil {
		return
	}

	for _, set := range sets {
		s := set.(Set)
		setItem := make(map[string]interface{})
		setItem["id"] = s.SetID
		setItem["appId"] = app.Id
		setItem["Name"] = s.SetName
		setItem["spriteCssClass"] = "c-icon icon-group fa-hide set"
		setItem["type"] = "set"
		setItem["expanded"] = false
		setItem["number"], _ = GetHostCount(s.SetID, "SetID")
		var modItems []interface{}
		for _, mod := range mods {
			m := mod.(Module)
			if m.SetId == s.SetID {
				modItem := make(map[string]interface{})
				modItem["id"] = m.Id
				modItem["appId"] = app.Id
				modItem["Name"] = m.ModuleName
				modItem["spriteCssClass"] = "c-icon icon-modal module"
				modItem["type"] = "module"
				modItem["number"], _ = GetHostCount(m.Id, "ModuleID")
				modItems = append(modItems, modItem)
				if m.ModuleName == "空闲机" {
					emptyItems = append(emptyItems, modItem)
				} else {
					if app.Level == 2 {
						setItems = append(setItems, modItem)
						options[m.Id] = m.ModuleName
					} else {
						options[m.Id] = fmt.Sprintf("%s-%s", s.SetName, m.ModuleName)
					}
				}
			}
		}
		if !s.Default {
			setItem["items"] = modItems
			setItems = append(setItems, setItem)
		}
	}

	topoItem := make(map[string]interface{})
	topoItem["id"] = app.Id
	topoItem["Name"] = app.ApplicationName
	topoItem["type"] = "application"
	topoItem["spriteCssClass"] = "c-icon icon-app application"
	topoItem["expanded"] = true
	topoItem["lvl"] = app.Level
	topoItem["number"], _ = GetHostCount(app.Id, "ApplicationID")
	topoItem["items"] = setItems
	topo = append(topo, topoItem)
	info["topo"] = topo
	info["empty"] = emptyItems

	return
}

// GetAllApp retrieves all App matches certain condition. Returns empty list if
// no records exist
func GetAllApp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(App))
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

	var l []App
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

// UpdateApp updates App by Id and returns error if
// the record to be updated doesn't exist
func UpdateAppById(m *App) (err error) {
	o := orm.NewOrm()
	v := App{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApp deletes App by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApp(id int) (err error) {
	// TODO: 检测要删除业务下是否有主机
	o := orm.NewOrm()
	v := App{Id: id}
	DeleteSetByAppId(id)
	DeleteModuleByAppId(id)
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&App{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
