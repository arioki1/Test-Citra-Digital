package main

import (
	"fmt"
	"reflect"
	"strings"
)

func hitBil(x string) {
	var split  [] string
	countH := 0;
	countM := 0;
	splitString := strings.Split(x,"")
	for i:=0; i<len(splitString) ; i++ {
		if chehckItem(split,splitString[i]){
			split= append(split,splitString[i])
			if "a" == splitString[i] || "i" == splitString[i] || "u" == splitString[i] ||
				"e" == splitString[i] || "o" == splitString[i]{
				countH++
			}else if " " != splitString[i] {
					countM++
			}
		}
	}

	fmt.Println("Huruf Hidup : ",countH)
	fmt.Println("Huruf Mati : ", countM)
}

func chehckItem(arrayTipe interface{}, item interface{}) bool {
	arr := 	reflect.ValueOf(arrayTipe)
	for i:=0; i<arr.Len();i++  {
		if arr.Index(i).Interface() == item {
			if item != "a" && item != "o" {
				return true
			}else {
				return false
			}
		}
	}
	return true
}

func main()  {
	hitBil("omama")
}