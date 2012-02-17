// errorcheck

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 1871.

package p

type a interface {
	foo(x int) (x int) // ERROR "redeclared|redefinition"
}

var b interface {
	bar(y int) (y int) // ERROR "redeclared|redefinition"
}

/*
Previously:

bug.go:1 x redclared in this block
    previous declaration at bug.go:1
*/
