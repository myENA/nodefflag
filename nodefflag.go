// Package nodefflag extends the go flag package to allow for a "no default"
// variation of standard flag variables.  In order to accomplish this,
// we have to use pointers to pointers.  If the pp references a nil pointer,
// the flag was not set.  If the pp references a non-nil pointer, the flag
// was set, and **ptr contains the value.  The pp itself returned will never
// be nil, and it is expected that the ND*Var methods will never receive
// a nil **.
//
// Example:
//
//    package main
//
//    import (
//    	"flag"
//    	"os"
//
//    	ndf "github.com/myENA/nodefflag"
//    )
//
//    func main() {
//    	var (
//    		bv *bool
//    		sv *string
//    	)
//
//    	flags := ndf.NewNDFlagSet(os.Args[0], flag.ExitOnError)
//
//    	flags.NDBoolVar(&bv, "bool flag", true, "this is a bool flag")
//    	flags.NDStringVar(&sv, "string flag", "Example", "this is a string flag")
//
//    	flags.Parse(os.Args[1:])
//
//    	// if -bv is not passed, bv == nil.  if -bv=true or -bv is passed,
//    	// *bv == true .  if -bv=false is passed, *bv == false.
//
//    	// if -sv is not passed, sv == nil.  if -sv="something" is passed,
//    	// *sv == "something".  if -sv="" is passed, *sv == "".
//
//    }
package nodefflag

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// implement the Value interface for flags
type ndsf struct {
	sv      **string
	example string
}

func (s *ndsf) String() string {
	return s.example
}

func (s *ndsf) Set(val string) error {
	*s.sv = &val
	return nil
}

func (s *ndsf) Get() interface{} {
	return *s.sv
}

type ndbf struct {
	bv      **bool
	example string
}

func (b *ndbf) String() string {
	return b.example
}

func (b *ndbf) Set(val string) error {
	pb, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*b.bv = &pb
	return nil
}

func (b *ndbf) Get() interface{} {
	return *b.bv
}

func (b *ndbf) IsBoolFlag() bool {
	return true
}

type ndif struct {
	iv      **int
	example string
}

func (i *ndif) String() string {
	return i.example
}

func (i *ndif) Set(val string) error {
	pi, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*i.iv = &pi
	return nil
}

func (i *ndif) Get() interface{} {
	return *i.iv
}

type ndi64f struct {
	iv      **int64
	example string
}

func (i *ndi64f) String() string {
	return i.example
}

func (i *ndi64f) Set(val string) error {
	pi, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	*i.iv = &pi
	return nil
}

func (i *ndi64f) Get() interface{} {
	return *i.iv
}

type nduif struct {
	uiv     **uint
	example string
}

func (ui *nduif) String() string {
	return ui.example
}

func (ui *nduif) Set(val string) error {
	pui, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	pui2 := uint(pui)
	*ui.uiv = &pui2
	return nil
}

func (ui nduif) Get() interface{} {
	return *ui.uiv
}

type ndui64f struct {
	uiv     **uint64
	example string
}

func (ui *ndui64f) String() string {
	return ui.example
}

func (ui *ndui64f) Set(val string) error {
	pui, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	*ui.uiv = &pui
	return nil
}

func (ui *ndui64f) Get() interface{} { return *ui.uiv }

type ndff struct {
	fv      **float64
	example string
}

func (f *ndff) String() string {
	return f.example
}

func (f *ndff) Set(val string) error {
	pf, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	*f.fv = &pf
	return nil
}

func (f *ndff) Get() interface{} {
	return *f.fv
}

type nddf struct {
	dv      **time.Duration
	example string
}

func (d *nddf) String() string {
	return d.example
}

func (d nddf) Set(val string) error {
	pd, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	*d.dv = &pd
	return nil
}

func (d nddf) Get() interface{} {
	return *d.dv
}

// NDFlagSet - extends the flag package to add "no default" variants,
// where no defaults are specified.
type NDFlagSet struct {
	*flag.FlagSet
	output io.Writer
	name   string
}

// NewNDFlagSet - factory method, initializes the underlying FlagSet
func NewNDFlagSet(name string, errorHandling flag.ErrorHandling) *NDFlagSet {
	fs := flag.NewFlagSet(name, errorHandling)
	ndf := &NDFlagSet{
		FlagSet: fs,
		name:    name,
	}
	ndf.FlagSet.Usage = ndf.ndfUsage
	return ndf
}

// NDString - returns double string pointer, will reference nil
// string pointer if flag was not set, will reference non-nil otherwise.
// This allows you to differentiate between the zero val ("") and not set.
func (ndf *NDFlagSet) NDString(name, example, usage string) **string {
	var sv *string
	ndf.NDStringVar(&sv, name, example, usage)
	return &sv
}

// NDStringVar - Similar to NDString, but you supply the double
// string pointer.
func (ndf *NDFlagSet) NDStringVar(sv **string, name, example, usage string) {
	s := &ndsf{sv: sv, example: example}
	ndf.Var(s, name, usage)
}

// NDBool - returns double bool pointer, will reference
// nil bool pointer if flag was not set, will reference non-nil otherwise.
// This allows you to differentiate between the zero val (false) and not set.
func (ndf *NDFlagSet) NDBool(name string, example bool, usage string) **bool {
	var bv *bool
	ndf.NDBoolVar(&bv, name, example, usage)
	return &bv
}

// NDBoolVar - similar to NDBool, but you supply the double
// bool pointer.
func (ndf *NDFlagSet) NDBoolVar(bv **bool, name string, example bool, usage string) {
	b := &ndbf{bv: bv, example: strconv.FormatBool(example)}
	ndf.Var(b, name, usage)
}

// NDInt - returns an int double pointers, will reference
// nil int pointer if flag was not set, will reference non-nil otherwise.
// This allows you to differentiate between the zero val (0) and not set.
func (ndf *NDFlagSet) NDInt(name string, example int, usage string) **int {
	var iv *int
	ndf.NDIntVar(&iv, name, example, usage)
	return &iv
}

// NDIntVar - similar to NDInt, but you sply the double pointer.
func (ndf *NDFlagSet) NDIntVar(iv **int, name string, example int, usage string) {
	i := &ndif{iv: iv, example: strconv.FormatInt(int64(example), 10)}
	ndf.Var(i, name, usage)
}

// NDInt64 - NDInt but type is int64
func (ndf *NDFlagSet) NDInt64(name string, example int64, usage string) **int64 {
	var iv *int64
	ndf.NDInt64Var(&iv, name, example, usage)
	return &iv
}

// NDInt64Var - NDIntVar but for int64
func (ndf *NDFlagSet) NDInt64Var(iv **int64, name string, example int64, usage string) {
	i := &ndi64f{iv: iv, example: strconv.FormatInt(example, 10)}
	ndf.Var(i, name, usage)
}

// NDUint - returns double pointer to a uint.
func (ndf *NDFlagSet) NDUint(name string, example uint, usage string) **uint {
	var uiv *uint
	ndf.NDUintVar(&uiv, name, example, usage)
	return &uiv
}

// NDUintVar - same as NDUint, but you supply the double p.
func (ndf *NDFlagSet) NDUintVar(uiv **uint, name string, example uint, usage string) {
	ui := &nduif{uiv: uiv, example: strconv.FormatUint(uint64(example), 10)}
	ndf.Var(ui, name, usage)
}

// NDUint64 - uint64 version of NDUint
func (ndf *NDFlagSet) NDUint64(name string, example uint64, usage string) **uint64 {
	var uiv *uint64
	ndf.NDUint64Var(&uiv, name, example, usage)
	return &uiv
}

// NDUint64Var - uint64 version of NDUintVar
func (ndf *NDFlagSet) NDUint64Var(uiv **uint64, name string, example uint64, usage string) {
	ui := &ndui64f{uiv: uiv, example: strconv.FormatUint(example, 10)}
	ndf.Var(ui, name, usage)
}

// NDFloat64 - returns double pointer to a float64.  Works the same
// as all the other numeric types.
func (ndf *NDFlagSet) NDFloat64(name string, example float64, usage string) **float64 {
	var fv *float64
	ndf.NDFloat64Var(&fv, name, example, usage)
	return &fv
}

// NDFloat64Var - you supply the pointer, but same as NDFloat64
func (ndf *NDFlagSet) NDFloat64Var(fv **float64, name string, example float64, usage string) {
	f := &ndff{fv: fv, example: strconv.FormatFloat(example, 'g', -1, 64)}
	ndf.Var(f, name, usage)
}

// NDDuration - duration flag.  returns double pointer, if references
// nil the flag was not set, otherwise it was set.
func (ndf *NDFlagSet) NDDuration(name string, example time.Duration, usage string) **time.Duration {
	var dv *time.Duration
	ndf.NDDurationVar(&dv, name, example, usage)
	return &dv
}

// NDDurationVar - BYO duration pp version of NDDuration
func (ndf *NDFlagSet) NDDurationVar(dv **time.Duration, name string, example time.Duration, usage string) {
	d := &nddf{dv: dv, example: example.String()}
	ndf.Var(d, name, usage)
}

// Lifted from / adapted from std lib flag.PrintDefauls.
func (ndf *NDFlagSet) printDefaults() {
	ndf.VisitAll(func(fl *flag.Flag) {
		s := fmt.Sprintf("  -%s", fl.Name) // Two spaces before -; see next two comments.
		name, usage := flag.UnquoteUsage(fl)
		if len(name) > 0 {
			s += " " + name
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}

		s += usage

		if _, ok := fl.Value.(*ndsf); ok {
			// put quotes on the value
			s += fmt.Sprintf(" (example %q)", fl.DefValue)
		} else {
			s += fmt.Sprintf(" (example %v)", fl.DefValue)
		}

		fmt.Fprint(ndf.out(), s, "\n")
	})
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (ndf *NDFlagSet) SetOutput(output io.Writer) {
	ndf.output = output
	ndf.FlagSet.SetOutput(output)
}

func (ndf *NDFlagSet) out() io.Writer {
	if ndf.output == nil {
		return os.Stderr
	}
	return ndf.output
}

func (ndf *NDFlagSet) ndfUsage() {

	if ndf.name == "" {
		fmt.Fprintf(ndf.out(), "Usage:\n")
	} else {
		fmt.Fprintf(ndf.out(), "Usage of %s:\n", ndf.name)
	}
	ndf.printDefaults()
}
