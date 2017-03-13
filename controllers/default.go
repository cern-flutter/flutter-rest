package controllers

import (
	"fmt"
	models "flutter-rest/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetVersion() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Get() {


	db, err := orm.GetDB("default")
	if err != nil {
    fmt.Println("get default DataBase")
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT major, minor, patch, message from t_schema_version")
	if err != nil{
		fmt.Println(err)
	}
	sv := models.SchemaVersion{}
	for rows.Next(){
		if err = rows.Scan(&sv.Major, &sv.Minor, &sv.Patch, &sv.Message); err!=nil{
			fmt.Println(err)
		}
	 fmt.Println(sv)
	}
	rows.Close()


	c.Data["Version"]= sv
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
    main.Data["Website"] = "My Website"
    main.Data["Email"] = "your.email.address@example.com"
    main.Data["EmailName"] = "Your Name"
    main.TplName = "default/hello-sitepoint.tpl"
}
