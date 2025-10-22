package main

import "testing"

func assertEqual(t testing.TB, got,want string){
	t.Helper() //marks this function as helper so that it can provide correct context in case a test case fails
	if got != want {
		t.Errorf("Assertion Failed:- Got:- %v but Required:- %v",got,want)
	}

}
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		assertEqual(t,Hello("Dhwanil"),"Hello Dhwanil");
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assertEqual(t,Hello(""),"Hello World")
	})
}