package main
import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNum uint32 = 30000
	fmt.Println(intNum)

	var floatNum float32 = 30000.09
	fmt.Println(floatNum)

	var floatNum32 float32 = 2.2
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int32 = 208
	var intNum2 int32 = 100
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var MyString string  = "Hello" + "World"
	fmt.Println(MyString)
	fmt.Println(utf8.RuneCountInString(MyString))

	var myBolean bool = false
	fmt.Println(myBolean)

	var intNum3 int
	fmt.Println(intNum3)

	MyVar := "text"
	fmt.Println(MyVar)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)

	var4,var5 := 1,2
	fmt.Println(var4,var5)

	const MyConst string = "bimbimbambam"
	fmt.Println(MyConst)
}