package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"remitano-share-video/models"
	_ "remitano-share-video/routers"
	"strings"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)


// Test register user happy
func TestRegisterHappy(t *testing.T) {
	body := fmt.Sprintf("{\"username\":\"%v\",\"password\":\"%v\"}", uuid.New().String(), uuid.New().String())
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
	username := uuid.New().String()
	password := uuid.New().String()
	user := models.User{
		Username: username,
		Password: password,
	}
	if err := user.Register(); err == nil {
		body := fmt.Sprintf("{\"username\":\"%v\",\"password\":\"%v\"}", username, password)
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
}