package main

import (
	"fmt"
	string_util "github.com/billychou/go-startup/stringutil"
	"math/rand"
)

func main() {
	fmt.Printf("Hello,world.\n")
	//  fmt.Printf(stringutil.Reverse("kskskafjkajaf"))
	fmt.Printf(string_util.Reverse("kskskafjkajaf"))
	fmt.Println("my favorite number is", rand.Intn(10))
}
