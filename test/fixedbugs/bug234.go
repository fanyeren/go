// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main() {
	//TODO(rsc): uncomment when this syntax is valid for receive+check closed
	//	c := make(chan int, 1)
	//	c <- 100
	//	x, ok := <-c
	//	if x != 100 || !ok {
	//		println("x=", x, " ok=", ok, " want 100, true")
	//		panic("fail")
	//	}
	//	x, ok = <-c
	//	if x != 0 || ok {
	//		println("x=", x, " ok=", ok, " want 0, false")
	//		panic("fail")
	//	}
}
