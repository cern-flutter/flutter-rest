package main

import (
	"fmt"

	_ "gitlab.cern.ch/flutter/flutter-rest/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		beego.AppConfig.String("pgUser"), beego.AppConfig.String("pgPass"), beego.AppConfig.String("pgDB"), beego.AppConfig.String("pgSslMode")))

}

func main() {

	orm.RunCommand()
	beego.Run()
}
