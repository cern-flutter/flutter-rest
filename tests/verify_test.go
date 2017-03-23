package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	//	context "github.com/astaxie/beego/context"

	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	_ "gitlab.cern.ch/flutter/flutter-rest/routers"
)

var vomsPath = "certs/vomsdir"
var caPath = "certs/ca"

func init() {

	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.BConfig.RunMode = "test"
	beego.TestBeegoInit(apppath)

}

// TestProxyValid is a sample to run an endpoint test
func TestProxyValid(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	beego.Debug("hola")
	//ctx := context.NewContext()
	//beego.Debug(ctx)
	//ctx.Reset(w, req)
	//ctx.Input = context.NewInput()

	//ctx.Input.SetData("hola", "maria")
	//beego.Debug(ctx.Input.GetData("hola"))

	//date := beego.Date(time.Date(2016, 05, 18, 12, 37, 30, 0, gmt), time.UnixDate)
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	beego.Trace("testing", "TestProxyValid", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

}
