package integers

import (
	"fmt"
	"testing"
)

func assertEqual(t testing.TB, got,want int){
	t.Helper() //marks this function as helper so that it can provide correct context in case a test case fails
	if got != want {
		t.Errorf("Assertion Failed:- Got:- %v but Required:- %v",got,want)
	}

}
func TestAdd(t *testing.T){
	t.Run("Adding two possitive numbers",func(t *testing.T) {
		assertEqual(t,Add(5,4),9)
	})
	t.Run("Adding multiple positive & negative numbers",func(t *testing.T) {
		assertEqual(t,Add(3,-9,5,-4,-4,3,1),-5)
	})
}

// Testable examples:-  When running test they are automatically executed and tested from output of console with value written in comments with "Output :...""
// For more knowledge visit https://go.dev/blog/examples
func ExampleAdd(){
	fmt.Println(Add(5,5)) // Output: 10
}

func ExampleAdd_second(){
	fmt.Println(Add(-5,-4,-3,-2,0,-1,-10,100000)) // Output:99975
}