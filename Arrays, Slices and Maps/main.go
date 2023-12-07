package main

import (
	"fmt"
	"time"
)

func main(){
	var intArr [3]int32

	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr2[0])

	intArr3 := [...]int32{1,2,3}
	fmt.Println(intArr3[0:3])

	var intSlice []int32 = []int32{4,6,8}
	fmt.Println(intSlice)
	fmt.Printf("\nThe length of the Slice is %v and a capacity is %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("\nThe length of the Slice is %v and a capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{10,11}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var MyMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(MyMap)

	var MyMap2 = map[string]uint8{"Adam":23,"Vlad":22}
	fmt.Println(MyMap2)
	fmt.Println(MyMap2["Vlad"])

	var age, ok = MyMap2["Adam"]
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	for name := range MyMap2{
		fmt.Printf("\nName: %v, Age: %v, Name in list: %v", name, age, ok)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v Value: %v\n", i, v)
	}

	var i int = 0
	for {
		if i>10{
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	var n = 1000000
	var TestSlice = []int{}
	var TestSlice2 = make([]int, 0, n)

	fmt.Printf("Total time is %v\n", timeLoop(TestSlice, n))
	fmt.Printf("Total time is %v\n", timeLoop(TestSlice2, n))
}
func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}