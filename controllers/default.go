package controllers

import (
	models "flutter-rest/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	baseController
}

func (this *MainController) NestPrepare() {

}

func (c *MainController) Get() {

	db, err := orm.GetDB("default")
	if err != nil {
		fmt.Println("get default DataBase")
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT major, minor, patch, message from t_schema_version")
	if err != nil {
		fmt.Println(err)
	}
	sv := models.SchemaVersion{}
	for rows.Next() {
		if err = rows.Scan(&sv.Major, &sv.Minor, &sv.Patch, &sv.Message); err != nil {
			fmt.Println(err)
		}
		fmt.Println(sv)
	}
	rows.Close()
	c.Data["APIVersion"] = beego.AppConfig.String("appversion")
	c.Data["SchemaVersion"] = sv
	c.Data["Email"] = "fts-support@cern.ch"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.TplName = "default/hello-sitepoint.tpl"
}
