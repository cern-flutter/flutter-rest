package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/astaxie/beego"
	context "github.com/astaxie/beego/context"
	. "github.com/smartystreets/goconvey/convey"
)

var vomsPath = "certs/vomsdir"
var caPath = "certs/ca"
var gmt *time.Location

func init() {
	var err error
	if gmt, err = time.LoadLocation("GMT"); err != nil {
		beego.Error(err)
	}
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	//apppath_test, _ := filepath.Abs(filepath.Dir(filepath.Join("app_test.conf", ".."+string(filepath.Separator))))
	//beego.InitBeegoBeforeTest(apppath)
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
	ctx := req.Context()
	//date := beego.Date(time.Date(2016, 05, 18, 12, 37, 30, 0, gmt), time.UnixDate)
	ctx = context.WithValue(ctx, "date", "hola")
	req.WithContext(ctx)
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	beego.Trace("testing", "TestProxyValid", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

}
