 Package nodefflag extends the go flag package to allow for a "no default"
 variation of standard flag variables.  In order to accomplish this,
 we have to use pointers to pointers.  If the pp references a nil pointer,
 the flag was not set.  If the pp references a non-nil pointer, the flag
 was set, and **ptr contains the value.  The pp itself returned will never
 be nil, and it is expected that the ND*Var methods will never receive
 a nil **.

 Example:
```go
package main

import (
	"flag"
	"os"
	"fmt"

	ndf "github.com/myENA/nodefflag"
)

func main() {
	var (
		bv *bool
		sv *string
	)

	flags := ndf.NewNDFlagSet(os.Args[0], flag.ExitOnError)

	flags.NDBoolVar(&bv, "bool flag", "this is a bool flag")
	flags.NDStringVar(&sv, "string flag", "this is a string flag")

	flags.Parse(os.Args[1:])

	// if -bv is not passed, bv == nil.  if -bv=true or -bv is passed,
	// *bv == true .  if -bv=false is passed, *bv == false.

	// if -sv is not passed, sv == nil.  if -sv="something" is passed,
	// *sv == "something".  if -sv="" is passed, *sv == "".

}
```
