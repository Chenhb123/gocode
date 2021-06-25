package main

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestGRPCCompatible(t *testing.T) {
	app := newApp()

	e := httptest.New(t, app)
	e.POST("/helloworld.Greeter/SayHello").WithJSON(map[string]string{"name": "makis"}).Expect().
		Status(httptest.StatusOK).JSON().Equal(map[string]string{"message": "Hello makis"})
}
