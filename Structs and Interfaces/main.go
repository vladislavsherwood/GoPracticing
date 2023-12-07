package main

import (
	"fmt"
)

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
	int
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type owner struct{
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh*e.kwh
}

type engine interface{
	milesLeft() uint8
}

func canYouMakeIt(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("Ye, i fckn can")
	}else{
		fmt.Println("Thats not gonna happen")
	}
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 10}
	myEngine.mpg = 27
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)

	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{100,100}
	var myEngine3 electricEngine = electricEngine{20,10}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("\nTotal miles left for Engine1:%v\n",myEngine.milesLeft())
	fmt.Printf("\nTotal miles left for Engine3:%v\n",myEngine3.milesLeft())
	


	canYouMakeIt(myEngine, 150)
	canYouMakeIt(myEngine3, 200)
	
}