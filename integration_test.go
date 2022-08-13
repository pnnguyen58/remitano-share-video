package main

import (
	"github.com/beego/beego/v2/server/web/session"
	"net/http"
	"net/http/httptest"
	_ "remitano-share-video/routers"
	"strings"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

var globalSessions *session.Manager
func init() {
	mc := session.ManagerConfig{
		CookieName: "gosessionid",
		EnableSetCookie: true,
		Gclifetime: 3600,
		Maxlifetime: 3600,
		Secure: false,
		SessionNameInHTTPHeader: "sha1",
		CookieLifeTime: 3600,
	}
	globalSessions, _ = session.NewManager("memory", &mc)
	go globalSessions.GC()
}

// Test register user happy
func TestRegisterHappy(t *testing.T) {
	body := `{
		"username": 	"testhappy",
		"password": 	"qazxwscdfe"
	}`
	r, _ := http.NewRequest("POST", "/v1/users/register/test", strings.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	logs.Trace("testing", "TestRegisterHappy", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestLoginHappy(t *testing.T) {
	body := `{
		"username": 	"testhappy",
		"password": 	"qazxwscdfe"
	}`
	r, _ := http.NewRequest("POST", "/v1/users/login/test", strings.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	logs.Trace("testing", "TestLoginHappy", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}