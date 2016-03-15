package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type App struct {
	Id int
	Type int
	ApplicationName string
	LifeCycle string
	Level int
	OwnerId int
}

func (t *App) TableName() string {
	return "app"
}

func init() {
	orm.RegisterModel(new(App))
}

func AddApp(m *App) (id int64, err error) {
	fmt.Println("999999999999999999=", m)
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func IsAppExistByName(applicationName string) (isExist bool) {
	o := orm.NewOrm()
	return o.QueryTable("app").Filter("application_name", applicationName).Exist()
}

func GetAppCountByUserId(userId int) (cnt int64, err error) {
	o := orm.NewOrm()
	cnt, err = o.QueryTable("app").Filter("owner_id", userId).Count()
	return
}

func GetApps(userId int) (num int64, l []App, err error) {
	o := orm.NewOrm()

	if num, err = o.QueryTable(new(App)).Filter("owner_id", userId).OrderBy("id").All(&l); err == nil {
		fmt.Println("num=", num)
		return num, l, nil
	}

	return 0, nil, err
}

func DeleteApp(id int) (err error) {
	o := orm.NewOrm()
	v := App{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&App{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database: ", num)
		}
	}
	return
}
