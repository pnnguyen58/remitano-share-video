package main

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"remitano-share-video/models"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	assertions := assert.New(t)
	username := uuid.New().String()
	var tests = []struct {
		user models.User
		expected string
	}{
		{models.User{
			Username: username,
			Password: "1qazxsw23edc",
		}, ""},
		{models.User{
			Username: username,
			Password: "1qazxsw23edc",
		}, "ERROR: duplicate key value violates unique constraint \"users_un\" (SQLSTATE 23505)"},
		{models.User{
			Username: uuid.New().String(),
			Password: "1qazxsw23edc",
		}, ""},
	}
	for _, test := range tests {
		err := test.user.Register()
		if err != nil {
			assertions.Equal(test.expected, err.Error())
		} else {
			assertions.Equal(test.expected, "")
		}
	}
}

func TestLoginUser(t *testing.T) {
	assertions := assert.New(t)
	username := uuid.New().String()
	password := uuid.New().String()
	user := models.User{
		Username: username,
		Password: password,
	}
	var tests = []struct {
		user models.User
		expected string
	}{
		{models.User{
			Username: username,
			Password: password,
		}, ""},
		{models.User{
			Username: username + "123",
			Password: password,
		}, "record not found"},
		{models.User{
			Username: username,
			Password: password + "123",
		}, "crypto/bcrypt: hashedPassword is not the hash of the given password"},
	}
	if err := user.Register(); err == nil {
		err = user.Login()
		for _, test := range tests {
			err = test.user.Login()
			if err != nil {
				assertions.Equal(test.expected, err.Error())
			} else {
				assertions.Equal(test.expected, "")
			}
		}
	}
}

func TestGetUser(t *testing.T) {
	assertions := assert.New(t)
	username := uuid.New().String()
	password := uuid.New().String()
	user := models.User{
		Username: username,
		Password: password,
	}
	if err := user.Register(); err == nil {
		err = user.Get()
		if err != nil {
			assertions.Equal("user not found", err.Error())
		} else {
			assertions.Equal("", "")
		}
	}
}
/*func TestLogin(t *testing.T) {
	assertions := assert.New(t)
	_, err := models.
	assertions.Equal(nil, err)
}

func TestLogin(t *testing.T) {
	assertions := assert.New(t)
	user := models.User{
		Username: "abc",
		Password: "1qazxsw23edc",
	}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	bCtx := beeCtx.Context{
		Request: req,
	}
	accepted, _, err := helper.BeforeTask(&bCtx)
	assertions.Equal(nil, err)
	assertions.Equal(true, accepted)
}

func TestBeforeTaskDisconnected(t *testing.T) {
	assertions := assert.New(t)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Nanosecond)
	defer cancel()
	req = req.WithContext(ctx)
	bCtx := beeCtx.Context{
		Request: req,
	}
	accepted, _, err := helper.BeforeTask(&bCtx)
	assertions.NotEqual(nil, err)
	assertions.Equal(false, accepted)
}
*/