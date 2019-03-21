package main

import (
	"fmt"
	string_util "github.com/billychou/go-startup/stringutil"
	"math/rand"

	"github.com/billychou/go-startup/stringutil"
)

const boilingF = 212.0

func main() {
	fmt.Printf("Hello,world.\n")
	//  fmt.Printf(stringutil.Reverse("kskskafjkajaf"))
	fmt.Printf(string_util.Reverse("kskskafjkajaf"))
	fmt.Printf(stringutil.Reverse("kskskafjkajaf"))
	fmt.Println("Helloworld!")
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Println("换行")	
}
