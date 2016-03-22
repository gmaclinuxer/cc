package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Set struct {
	SetID            int       `orm:"column(id);auto"`
	ApplicationID int      `orm:"column(application_id)"`
	Capacity      int       `orm:"column(capacity)"`
	ChnName       string    `orm:"column(chn_name);size(255);null"`
	CreateTime    time.Time `orm:"column(create_time);type(timestamp);null"`
	Default       int      `orm:"column(default);null"`
	Description   string    `orm:"column(description);size(255);null"`
	EnviType      int      `orm:"column(envi_type)"`
	LastTime      time.Time `orm:"column(last_time);type(timestamp);null"`
	OpenStatus    string      `orm:"column(open_status);size(255)"`
	ParentID      int       `orm:"column(parent_id);null"`
	ServiceStatus int      `orm:"column(service_status)"`
	SetName       string    `orm:"column(set_name);size(255)"`
	Owner         int      `orm:"column(owner)"`
}

func (t *Set) TableName() string {
	return "set"
}

func init() {
	orm.RegisterModel(new(Set))
}

// AddSet insert a new Set into database and returns
// last inserted Id on success.
func AddSet(m *Set) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSetById retrieves Set by Id. Returns error if
// Id doesn't exist
func GetSetById(id int) (v *Set, err error) {
	o := orm.NewOrm()
	v = &Set{SetID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSet retrieves all Set matches certain condition. Returns empty list if
// no records exist
func GetAllSet(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Set))
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

	var l []Set
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

// UpdateSet updates Set by Id and returns error if
// the record to be updated doesn't exist
func UpdateSetById(m *Set) (err error) {
	o := orm.NewOrm()
	v := Set{SetID: m.SetID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSet deletes Set by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSet(id int) (err error) {
	o := orm.NewOrm()
	v := Set{SetID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Set{SetID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
