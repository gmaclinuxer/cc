package models

import (
	"strconv"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type App struct {
	Id              int    `orm:"column(id);auto"`
	Type            int8   `orm:"column(type)"`
	ApplicationName string `orm:"column(application_name);size(255)"`
	LifeCycle       string `orm:"column(life_cycle);size(255)"`
	Level           int8   `orm:"column(level)"`
	OwnerId         int   `orm:"column(owner_id)"`
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

func GetAppTopoById(id int) (ml []interface{}, err error) {
//	[{"id":"5524","text":"aaa","spriteCssClass":"c-icon icon-group","type":"set","expanded":false,"number":32,"items":[{"id":"7025","spriteCssClass":"c-icon icon-modal","text":"1","operator":"1842605324","bakoperator":"1842605324","type":"module","number":32}]}]
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0
	
	query["application_id"] = strconv.Itoa(id)
	query["default"] = "0"
	
	
	
	if sets, err := GetAllSet(query, fields, sortby, order, offset, limit); err == nil {
		for _, set := range sets {
			s := make(map[string]interface{})
			s["id"] = set.(Set).SetID
			s["text"] = set.(Set).SetName
			s["spriteCssClass"] = "c-icon icon-group"
			s["type"] = "set"
			s["expanded"] = false
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
				s["items"] = items
			}
			ml = append(ml, s)
		}
	}
	
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
	o := orm.NewOrm()
	v := App{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&App{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
