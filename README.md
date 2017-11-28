[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/myENA/nodefflag)](https://goreportcard.com/report/github.com/myENA/nodefflag)
[![GoDoc](https://godoc.org/github.com/myENA/nodefflag?status.svg)](https://godoc.org/github.com/myENA/nodefflag)
[![Build Status](https://travis-ci.org/myENA/nodefflag.svg?branch=master)](https://travis-ci.org/myENA/nodefflag)

Package nodefflag extends the go flag package to allow for a "no default"
variation of standard flag variables.  There are two types of nodef flag
variables we define: those where a double pointer is specified, so you can
differentiate between the zero value and not set, and another where the zero
value is the same as not set.  For the double pointer variant, if the pp 
references a nil pointer, the flag was not set.  If the pp references a
non-nil pointer, the flag was set, and **ptr contains the value.  The pp
itself returned will never be nil, and it is expected that the ND*Var methods 
will never receive a nil **.

 Example:
```go
package main

import (
	"flag"
	"os"

	ndf "github.com/myENA/nodefflag"
)

func main() {
	var (
		bv *bool
		sv *string
		zvstr string
	)

	flags := ndf.NewNDFlagSet(os.Args[0], flag.ExitOnError)

	flags.NDBoolVar(&bv, "bf", true, "this is a bool flag")
	flags.NDStringVar(&sv, "sv", "Example", "this is a string flag")
	flags.ZVStringVar(&zvstr, "zv", "example", "this is a string flag")

	flags.Parse(os.Args[1:])

	// if -bv is not passed, bv == nil.  if -bv=true or -bv is passed,
	// *bv == true .  if -bv=false is passed, *bv == false.

	// if -sv is not passed, sv == nil.  if -sv="something" is passed,
	// *sv == "something".  if -sv="" is passed, *sv == "".
	
	// if -zv is not passed, zv == "".  if -zv="something" is passed,
	// zv == "something.  if -zv="" is passed, zv == ""

}
```
