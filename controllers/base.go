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
	if proxyPath == "" {
		proxyPath = getProxyPath()
	}
	var p proxy.X509Proxy
	if err := p.DecodeFromFile(proxyPath); err != nil {
		fmt.Println("Error Decode")
	}

	verifyOptions := proxy.VerifyOptions{
		VomsDir:     beego.AppConfig.String("vomsDir"),
		CurrentTime: time.Now(),
	}
	capath := beego.AppConfig.String("capath")
	crls, _ := strconv.ParseBool(beego.AppConfig.String("crls"))

	if verifyOptions.Roots, err = proxy.LoadCAPath(capath, crls); err != nil {
		fmt.Println("Failed to load the CA Path: ", err)
		return
	}

	if err = p.Verify(verifyOptions); err != nil {
		fmt.Println("Your proxy doesn't seem valid ", err)
	} else {
		fmt.Println("Your proxy seems valid")
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
