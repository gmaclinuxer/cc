package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int
	UserName  string `orm:"column(username);size(20)"`
	Password  string `orm:"column(password);size(32)"`
	Salt      string `orm:"column(salt);size(10)"`
	Email     string
	LastLogin int64
	LastIp    string
	Status    int
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByName(userName string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{UserName: userName}
	if err = o.QueryTable("user").Filter("username", userName).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func UpdateUser(user *User, fields ...string) error {
	o := orm.NewOrm()
	_, err := o.Update(user, fields...)
	return err
}
