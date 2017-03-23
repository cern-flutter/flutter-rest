package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"gitlab.cern.ch/flutter/go-proxy"
)

type NestPreparer interface {
	NestPrepare()
}

func getProxyPath() string {
	if path := os.Getenv("X509_USER_PROXY"); path != "" {
		return path
	}

	return fmt.Sprintf("/tmp/x509up_u%d", os.Getuid())
}

// baseRouter implements global settings for all other routers.
type baseController struct {
	beego.Controller
}

// Prepare implements Prepare method for baseRouter.
func (this *baseController) Prepare() {
	var err error
	var proxyPath string
	var gmt *time.Location

	beego.Debug("baseController")
	if gmt, err = time.LoadLocation("GMT"); err != nil {
		beego.Error(err)
	}

	current_time := time.Now()
	beego.Debug(beego.BConfig.RunMode)
	if beego.BConfig.RunMode == "test" {
		proxyPath = beego.AppConfig.String("proxy")
		current_time = time.Date(2016, 05, 18, 12, 37, 30, 0, gmt)
	}
	if proxyPath == "" {
		proxyPath = getProxyPath()
	}
	var p proxy.X509Proxy
	if err = p.DecodeFromFile(proxyPath); err != nil {
		beego.Error(err)
	}

	verifyOptions := proxy.VerifyOptions{
		VomsDir:     beego.AppConfig.String("vomsDir"),
		CurrentTime: current_time,
	}
	capath := beego.AppConfig.String("capath")
	crls, _ := strconv.ParseBool(beego.AppConfig.String("crls"))

	if verifyOptions.Roots, err = proxy.LoadCAPath(capath, crls); err != nil {
		beego.Debug("Failed to load the CA Path: ", err)
		return
	}

	if err = p.Verify(verifyOptions); err != nil {
		beego.Debug("Your proxy doesn't seem valid ", err)
	} else {
		beego.Debug("Your proxy seems valid")
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
