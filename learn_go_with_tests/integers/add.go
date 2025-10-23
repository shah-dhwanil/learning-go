package integers

// Adds integer and return its result
// Uses varidaict type which is similar to args in python
func Add(values ...int) (sum int){
	for _,val:= range values{
		sum+=val
	}
	return
}