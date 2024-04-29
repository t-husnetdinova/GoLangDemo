package hello_test

import "testing"
import "github.com/t-husnetdinova/GoLangDemo/hello"

func TestHello(t *testing.T) {
	if hello.HelloWorld() != "Hey Planet" {
		t.Fatal("wrong greeting :(")
	}
}