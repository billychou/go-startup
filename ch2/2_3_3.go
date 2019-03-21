/*
 * Filename: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2/2_3_3.go
 * Path: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2
 * Created Date: Saturday, March 16th 2019, 10:36:15 am
 * Author: qppzsc/*
 * Filename: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2/2_3_3.go
 * Path: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2
 * Created Date: Saturday, March 16th 2019, 10:36:15 am
 * Author: qppzsc
 *
 * Copyright (c) 2019 Your Company
 */

package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func newInt() *int {
	return new(int)
}

func delta(old, new int) int {
	return new - old
}
