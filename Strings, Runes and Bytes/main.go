package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Println(indexed)

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString{
		fmt.Println(i,v)
	}

	var strSlice = []string{"b","i","m","b","i","m","b","a","m","b","a","m"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v",catStr)

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v",catStr2)
}