package main

import "testing"

func TestHello(t *testing.T){
	expected:= "Hello Dhwanil"
	if got:=Hello("Dhwanil");got!=expected{
		t.Errorf("Got %q but expected %q",got,expected)
	}
}