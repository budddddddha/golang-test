package main

import (
    "testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestExampleSuccess(t *testing.T) {
	err := true
	result := "moge"
    if err != false {
        t.Fatalf("failed test %#v", err)
    }
    if result != "hige" {
        t.Fatal("failed test")
    }
}