package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main(){
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com","costco.com","wholefoods.com"}
	for i:= range websites{
		go checkChickenPrices(websites[i],chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(websites string, chickenChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var ChickenPrice = rand.Float32()*20
		if ChickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- websites
			break
		}
	}
}

func sendMessage(chickenChannel chan string){
	fmt.Printf("\nFound a good deal for chicken at %s", <- chickenChannel)
}