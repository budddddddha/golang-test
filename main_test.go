package main

import (
	"net/http"
	"fmt"
	// "net/http/httptest"
	"testing"

	"gopkg.in/gavv/httpexpect.v2"
)

func TestFruits(t *testing.T) {
	// create httpexpect instance
	e := httpexpect.New(t, "http://localhost:8080")

	fmt.Println("test")

	// is it working?
	obj := e.GET("/hoge").
		Expect().
		Status(http.StatusOK).JSON().Object()
	fmt.Println(obj)
}