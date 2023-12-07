package main

import (
	"fmt"
	"errors"
)

func main(){
	var printValue = "Hello World"
	PrintMe(printValue)

	var numerator = 233
	var denomenator = 9
	var result, remainder, err = intDevision(numerator,denomenator)
	switch{
		case err!=nil:
			fmt.Printf(err.Error())
		case remainder == 0:
			fmt.Printf("The result of integer devision is %v\n",result)
		default:
		fmt.Printf("The result of devision is: %v with remainder %v\n", result, remainder)
		}
	switch remainder{
		case 0:
			fmt.Println("No remainder!")
		case 1,2:
			fmt.Println("Remainder is 1 or 2.... not a big deal")
		default:
			fmt.Println("BIMBIMBAMBAM WE HAVE TO CARE ABOUT FCKN REMAINDER IT US MORE THAN 2")

	}
}

func PrintMe(printValue string){
	fmt.Println(printValue)
}

func intDevision(numerator int, denomenator int) (int, int, error) {
	var err error
	if denomenator==0{
		err=errors.New("Cannot Devide by Zero")
		return 0,0, err

	}
	var result int = numerator/denomenator
	var remainder int = numerator%denomenator
	return result, remainder, err
}