package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

// 声明常量
const boilingF = 212.0

var n = flag.Bool("n", false, "omit trailing newline")
var name = flag.String("name", "", "name")
var sep = flag.String("s", " ", "separator")

func main() {
	// 声明变量
	var f = boilingF
	var c = (f - 32) * 5 / 9
	// Printf
	// Doc: https://go-zh.org/pkg/fmt/
	fmt.Printf("boilingF point = %gF or %gC\n", f, c)
	fmt.Printf("func: point=%g\n", fToC(f))
	fmt.Println("func: point=%g\n", fToC(f))

	var s string = "欢迎"
	fmt.Println(s)
	execFmt()
	execRand()
	execRand()
	execPointer()
	execInt()
	execFloat64()
	
	fmt.Printf("test for flag\n")
	flag.Parse()
	fmt.Printf("your name:%s\n", *name)
	fmt.Printf("your sep:%s", *sep)	
	fmt.Print(strings.Join(flag.Args(), *sep))
}

// fmt.

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9;
}

func execFmt()  {
	name := "billychou"
	fmt.Printf("Welcome, %s\n", name)	
}

func execRand() {
	value := rand.Float64() * 3.0	
	fmt.Printf("value:%f\n", value)
}

func execPointer() {
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)
	var x1, y1 int
	fmt.Println(&x1)
	fmt.Println(&y1)
	var p1 = f()
	var p2 = f()
    fmt.Println(p1)	
    fmt.Println(p2)	

	var x2 *int
	fmt.Println(x2)
    v := 1	
	v1 := incr(&v)
	fmt.Println("incr:", v1)
}


func execFlag() {

}

func execInt() {
	var x,y int
	fmt.Println(x, y)
}

func execFloat64() {
	var x,y float64
	fmt.Println(x, y)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++	
	return *p
}
