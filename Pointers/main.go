package main

import (
	"fmt"
)

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Println("Test text for pointer: %v", *p)
	fmt.Println("Just a value: %v", i)
	p = &i
	*p = 10
	fmt.Println("Test text for pointer: %v", *p)
	fmt.Println("Just a value: %v", i)


	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of thing1 is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is %v", result)
	fmt.Printf("\nThe value of the thing is %v", thing1)
}

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of thing2 is %p\n", thing2)	
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}
