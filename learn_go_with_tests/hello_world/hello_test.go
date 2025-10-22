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
		assertEqual(t,Hello("Dhwanil",ENGLISH),"Hello Dhwanil");
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assertEqual(t,Hello("",ENGLISH),"Hello World")
	})

	t.Run("Testing language support with Gujarati",func(t *testing.T){
		assertEqual(t,Hello("Dhwanil",GUJARATI),"નમસ્તે Dhwanil");
	})
	t.Run("Testing if english is default language in case of invalid language provided",func(t *testing.T) {
		assertEqual(t,Hello("Dhwanil","zh"),"Hello Dhwanil")
	})
}