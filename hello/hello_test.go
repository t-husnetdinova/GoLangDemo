package hello_test

import "testing"
import "github.com/t-husnetdinova/GoLangDemo/hello"

// sample test function (it will fail, since this is the wrong greeting)
func TestHello(t *testing.T) {
	if hello.HelloWorld() != "Hello World!!! :)" {
		t.Fatal("wrong greeting :(")
	}
}