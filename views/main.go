package main

import (
	_ "github.com/lib/pq"
	_ "flutter-rest/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)


func init(){
    orm.RegisterDriver("postgres", orm.DRPostgres)
	  orm.RegisterDataBase("default", "postgres", "user=marsuaga password=chikitina dbname=mysite_development sslmode=disable")

}

func main() {
		orm.RunCommand()
		beego.Run()
}
