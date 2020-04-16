package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("hello world\n")
	//println("#$%^&*()_\n%%%%%%%%%%%%%%%%%%%%\n")
	//
	//var digits int = 9876543210
	//digits += 111
	//println("digits var: %s", digits)
	//fmt.Print("#{digits}")
	//fmt.Printf("#{digits}")
	//
	//cycle
	//var arr = [...]int{1,2,3}
	//for i:=0; i<len(arr); i++{
	//	println(arr[i])
	//}
	//
	//println("arr[1]: ",arr[1])
	//
	//
	//var slice []int
	//var slice2 = []int{11,22}
	//
	//
	//println(slice, slice2)
	//println("slice2 cap: ", cap(slice2), " len: ", len(slice2))
	//println("slice", slice)
	//slice2 = append(slice2, 33)
	//println("slice2 cap: ", cap(slice2), " len: ", len(slice2))
	//
	//slice = append(slice, slice2...)
	//println("slice", slice)
	//fmt.Println("sl", slice)
	//fmt.Println("s2", slice2)

	cpSlice := []string{"one", "two", "three"}
	//trgSlice := make([]string, 2)
	//
	//copy(trgSlice, cpSlice)
	//fmt.Println(trgSlice)
	myArray := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("myArray", myArray)
	fmt.Println("cpSlice", cpSlice)

	//iteration throw allay
	//myArrayLen := len(myArray)
	//println ("----------------------")
	//i := 0
	//for i < myArrayLen {
	//	println(myArray[i])
	//	i++
	//}
	//println ("----------------------")
	//for i:= range myArray{
	//	println(myArray[i])
	//}
	//println ("----------------------")
	//for i := range cpSlice {
	//	println(i, cpSlice[i])
	//}
	//println ("----------------------")
	//	//for _, val:= range cpSlice{
	//	//	//println(i,val)
	//	//	println(val)
	//	//}
	/* MAPS section*/
	//var m = make([string]string map)

	// ERROR -----var m1 map[string]string //cant assign value for map element with this declaration
	// need declaration below like
	//var m1 map[string]string = map[string]string{}
	//m1 := map[string]string{}
	//m1["WOW !!!"] = "effect"
	////map  of map
	//mm1 := map[string] map[string]string{}
	//mm1["SOME"] = m1
	//fmt.Println(mm1)

	//var m2 = make(map[string]int)
	//m2 := make(map[string]int)
	//m2["ONE"] = 1
	//fmt.Println("m2", m2)
	//
	//exist := false
	//_, exist = m2["Apsent"]
	//if !exist{
	//	fmt.Println("Apsent")
	//}
	//fmt.Println("----------------------")
	//mm := map[int]int{
	//	1:111,
	//	2:222,
	//}
	//if _, ok := mm[1]; ok{
	//	fmt.Println(mm[1])
	//}
	//// map iteration
	//for k,v :=  range mm{
	//	fmt.Println(k, v)
	//}

	//for i := range mm1 {
	//	mm1[i] = map[string]string{}
	//}

	//m4 :=map[string]string{
	//	"Пароль": "-Ответ",
	//	"Палка": "Селедка",
	//}
	//fmt.Println("m1", m1)
	//fmt.Println("m4", m4)

	// SWITCHES   ------------------------
	fmt.Println("----------------------")
	mm22 := map[int]string{
		0: "False", //0:"Abro-Cadabro",//
		1: "True",
	}

	switch mm22[0] {
	case "NULL":
		println("year - NULL")
	case "False":
		println("year - FASLE")
	case "SOME":
		println("it's some...")
	default:
		println("Nothing (((")
	}

	// strings.Contains(it, "lse")
	switch {
	case strings.Contains(mm22[0], "als"):
		println("contains 'als'")
		//break
		//fallthrough
	case strings.Contains(mm22[0], "aaa"):
		println("aaa")
	case strings.Contains(mm22[0], "lse"):
		println("contains 'lse'")
		//fallthrough
	default:
		println("no subs found")
	}

myAnchor:
	for true {
		println("one")
		break myAnchor
		println("two")
		println("three")
	}
	println("END")

	fmt.Println(mm22)
}
