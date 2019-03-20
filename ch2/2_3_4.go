/*
 * Filename: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2/2_3_4.go
 * Path: /Users/qppzsc/go-work/src/github.com/billychou/go-startup/ch2
 * Created Date: Wednesday, March 20th 2019, 10:26:36 pm
 * Author: qppzsc
 *
 * Copyright (c) 2019 Your Company
 */
package main

import "math"

func main() {
	for t := 0; t < cycles*2*math.Pi; t += res {
		x := math.Sin(t)
		y := math.Sin(t*freq + phase)
		img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
	}
}
