package main

import (
	"fmt"
)

func before(i string) string {
	return i
}

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

	//cpSlice := []string{"one", "two", "three"}
	//trgSlice := make([]string, 2)
	//
	//copy(trgSlice, cpSlice)
	//fmt.Println(trgSlice)
	//myArray := [...]int{0, 1, 2, 3, 4, 5}
	//fmt.Println("myArray", myArray)
	//fmt.Println("cpSlice", cpSlice)

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
	//	fmt.Println("----------------------")
	//	mm22 := map[int]string{
	//		0: "False", //0:"Abro-Cadabro",//
	//		1: "True",
	//	}
	//
	//	switch mm22[0] {
	//	case "NULL":
	//		println("year - NULL")
	//	case "False":
	//		println("year - FASLE")
	//	case "SOME":
	//		println("it's some...")
	//	default:
	//		println("Nothing (((")
	//	}
	//
	//	// strings.Contains(it, "lse")
	//	switch {
	//	case strings.Contains(mm22[0], "als"):
	//		println("contains 'als'")
	//		//break
	//		//fallthrough
	//	case strings.Contains(mm22[0], "aaa"):
	//		println("aaa")
	//	case strings.Contains(mm22[0], "lse"):
	//		println("contains 'lse'")
	//		//fallthrough
	//	default:
	//		println("no subs found")
	//	}
	//
	//myAnchor:
	//	for true {
	//		println("one")
	//		break myAnchor
	//		println("two")
	//		println("three")
	//	}
	//	println("END")
	//
	//	fmt.Println(mm22)

	// *** FUNCTIONS ***

	//fmt.Println(before("1111"))
	//fmt.Println(after("2222"))
	//fmt.Println(multi())
	//fmt.Println(rargs(1,2,3))
	//// defer panic recover
	//z := 1e4
	//fmt.Println("z: ", z)
	//fmt.Println("z/2: ", z/2)
	//fmt.Println("---------------------------")
	//str_1 := "first"
	//str_2 := "second"
	//
	//defer
	//	fmt.Println(str_1)
	//
	//panic("Before first & second")
	//
	//msg := recover()
	//fmt.Println(msg)
	//
	//fmt.Println(str_2)

	//	matreshka_1()

	// *** POINTERS ***
	// pString  := new(string)
	// //ps2  = new(string)
	//var psi *int = new(int)
	//*psi = 5
	//
	//fmt.Println("pString: ", *pString)
	//fmt.Println("*psi: ", *psi)
	//fmt.Println("psi: ", psi)
	//
	//psi2 := new(int)
	//psi2 = psi
	//*psi2 = 4
	//
	//fmt.Println("*psi: ", *psi)
	//fmt.Println("psi: ", psi)
	//
	//x := 1.5
	//square(&x)
	//fmt.Println(x)
	// *** STRUCTURES ***
	type FileMon struct {
		fname   string // file name
		fpath   string // file path
		fsize   int    // file size
		foffset int    // offset from the start
	}

	fm := FileMon{fname: "1.xtx", fpath: "c:/111/222", fsize: 100, foffset: 50}
	fmt.Println(fm)

	fm2 := FileMon2{"", "", 100, 50}
	fmt.Println(fm2)

	fm3 := FileMon2{"one", "two", 400, 500}
	prns(&fm3)
	fmt.Println(fm3.getName())

	a1 := new(Android)
	a2 := new(Android2)
	a1.Talk()
	a1.Person.Talk()
	a2.Person.Talk() //a2.Talk() - dont work
	//call interface
	essTolk(a1)
	// essTolk(a2) - dont work
}

// structure functions
type FileMon2 struct {
	fname, fpath   string // file name, file path
	fsize, foffset int    // file size, offset from the start
}

func prns(pfm *FileMon2) {
	pfm.fname = "nname.txt"
}

func (fm2 *FileMon2) getName() string {
	return fm2.fname
}

// person struct
type Person struct {
	Name string
}
type Animal struct {
	Name string
}

func (p *Animal) Talk() {
	fmt.Println("WOw wow wow")
}
func (p *Person) Talk() {
	fmt.Println("The Name: ", p.Name)
}

type Android struct {
	Person
	Model string
}
type Android2 struct {
	Person Person
	Model  string
}

//interfaces
type Essence interface {
	Talk()
}

func essTolk(ess Essence) {
	fmt.Println("talk from interface")
	ess.Talk()
}

type EssCommunity struct {
	ecom []Essence
}

func (ec *EssCommunity) Talk() {
	for _, s := range ec.ecom {
		s.Talk()
	}
}

//pointer examples
func square(x *float64) {
	*x = *x * *x
}

// function after main
func after(i string) (ret string) {
	ret = i
	return
}

// multiple results
func multi() (one string, two string) {
	one = "ONE"
	two = "TWO"
	return
}

// random args
func rargs(args ...int) (total int) {
	total = 0
	for _, v := range args {
		total += v
	}
	return
}

// Замыкания не понял http://golang-book.ru/chapter-07-functions.html

// panic recover - print stack
func matreshka_1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Exception in matreshka_1")
			panic(e)
		}
	}()
	matreshka_2()
}
func matreshka_2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Exception in matreshka_2")
			panic(e)
		}
	}()
	matreshka_3()
}
func matreshka_3() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Exception in matreshka_3: ", e)
			panic(e)
		}
	}()
	panic("root error")
}
