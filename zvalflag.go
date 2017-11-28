package nodefflag

import (
	"strconv"
	"time"
)

// implement the Value interface for flags
type zvsf struct {
	sv      *string
	example string
}

func (s *zvsf) String() string {
	return s.example
}

func (s *zvsf) Set(val string) error {
	*s.sv = val
	return nil
}

func (s *zvsf) Get() interface{} {
	return *s.sv
}

type zvbf struct {
	bv      *bool
	example string
}

func (b *zvbf) String() string {
	return b.example
}

func (b *zvbf) Set(val string) error {
	pb, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*b.bv = pb
	return nil
}

func (b *zvbf) Get() interface{} {
	return *b.bv
}

func (b *zvbf) IsBoolFlag() bool {
	return true
}

type zvif struct {
	iv      *int
	example string
}

func (i *zvif) String() string {
	return i.example
}

func (i *zvif) Set(val string) error {
	pi, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*i.iv = pi
	return nil
}

func (i *zvif) Get() interface{} {
	return *i.iv
}

type zvi64f struct {
	iv      *int64
	example string
}

func (i *zvi64f) String() string {
	return i.example
}

func (i *zvi64f) Set(val string) error {
	pi, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	*i.iv = pi
	return nil
}

func (i *zvi64f) Get() interface{} {
	return *i.iv
}

type zvuif struct {
	uiv     *uint
	example string
}

func (ui *zvuif) String() string {
	return ui.example
}

func (ui *zvuif) Set(val string) error {
	pui, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	pui2 := uint(pui)
	*ui.uiv = pui2
	return nil
}

func (ui zvuif) Get() interface{} {
	return *ui.uiv
}

type zvui64f struct {
	uiv     *uint64
	example string
}

func (ui *zvui64f) String() string {
	return ui.example
}

func (ui *zvui64f) Set(val string) error {
	pui, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	*ui.uiv = pui
	return nil
}

func (ui *zvui64f) Get() interface{} { return *ui.uiv }

type zvff struct {
	fv      *float64
	example string
}

func (f *zvff) String() string {
	return f.example
}

func (f *zvff) Set(val string) error {
	pf, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	*f.fv = pf
	return nil
}

func (f *zvff) Get() interface{} {
	return *f.fv
}

type zvdff struct {
	dv      *time.Duration
	example string
}

func (d *zvdff) String() string {
	return d.example
}

func (d *zvdff) Set(val string) error {
	pd, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	*d.dv = pd
	return nil
}

func (d *zvdff) Get() interface{} {
	return *d.dv
}

// ZVString - returns double string pointer, will reference nil
// string pointer if flag was not set, will reference non-nil otherwise.
// This allows you to differentiate between the zero val ("") and not set.
func (ndf *NDFlagSet) ZVString(name, example, usage string) *string {
	var sv string
	ndf.ZVStringVar(&sv, name, example, usage)
	return &sv
}

// AVStringVar - Similar to AVString, but you supply the double
// string pointer.
func (ndf *NDFlagSet) ZVStringVar(sv *string, name, example, usage string) {
	s := &zvsf{sv: sv, example: example}
	ndf.Var(s, name, usage)
}

// ZVBool - returns bool pointer, will be
// false if flag was not set, will be the parsed value otherwise.
func (ndf *NDFlagSet) ZVBool(name string, example bool, usage string) *bool {
	var bv bool
	ndf.ZVBoolVar(&bv, name, example, usage)
	return &bv
}

// ZVBoolVar - similar to ZVBool, but you supply the double
// bool pointer.
func (ndf *NDFlagSet) ZVBoolVar(bv *bool, name string, example bool, usage string) {
	b := &zvbf{bv: bv, example: strconv.FormatBool(example)}
	ndf.Var(b, name, usage)
}

// ZVInt - returns an int pointers, will be
// zero if flag was not set, will be the parsed value otherwise.
func (ndf *NDFlagSet) ZVInt(name string, example int, usage string) *int {
	var iv int
	ndf.ZVIntVar(&iv, name, example, usage)
	return &iv
}

// ZVIntVar - similar to ZVInt, but you supply the pointer.
func (ndf *NDFlagSet) ZVIntVar(iv *int, name string, example int, usage string) {
	i := &zvif{iv: iv, example: strconv.FormatInt(int64(example), 10)}
	ndf.Var(i, name, usage)
}

// ZVInt64 - ZVInt but type is int64
func (ndf *NDFlagSet) ZVInt64(name string, example int64, usage string) *int64 {
	var iv int64
	ndf.ZVInt64Var(&iv, name, example, usage)
	return &iv
}

// ZVInt64Var - ZVIntVar but for int64
func (ndf *NDFlagSet) ZVInt64Var(iv *int64, name string, example int64, usage string) {
	i := &zvi64f{iv: iv, example: strconv.FormatInt(example, 10)}
	ndf.Var(i, name, usage)
}

// ZVUint - returns pointer to a uint.
func (ndf *NDFlagSet) ZVUint(name string, example uint, usage string) *uint {
	var uiv uint
	ndf.ZVUintVar(&uiv, name, example, usage)
	return &uiv
}

// ZVUintVar - same as ZVUint, but you supply the double p.
func (ndf *NDFlagSet) ZVUintVar(uiv *uint, name string, example uint, usage string) {
	ui := &zvuif{uiv: uiv, example: strconv.FormatUint(uint64(example), 10)}
	ndf.Var(ui, name, usage)
}

// ZVUint64 - uint64 version of ZVUint
func (ndf *NDFlagSet) ZVUint64(name string, example uint64, usage string) *uint64 {
	var uiv uint64
	ndf.ZVUint64Var(&uiv, name, example, usage)
	return &uiv
}

// ZVUint64Var - uint64 version of ZVUintVar
func (ndf *NDFlagSet) ZVUint64Var(uiv *uint64, name string, example uint64, usage string) {
	ui := &zvui64f{uiv: uiv, example: strconv.FormatUint(example, 10)}
	ndf.Var(ui, name, usage)
}

// ZVFloat64 - returns pointer to a float64.  Works the same
// as all the other numeric types.
func (ndf *NDFlagSet) ZVFloat64(name string, example float64, usage string) *float64 {
	var fv float64
	ndf.ZVFloat64Var(&fv, name, example, usage)
	return &fv
}

// ZVFloat64Var - you supply the pointer, but same as ZVFloat64
func (ndf *NDFlagSet) ZVFloat64Var(fv *float64, name string, example float64, usage string) {
	f := &zvff{fv: fv, example: strconv.FormatFloat(example, 'g', -1, 64)}
	ndf.Var(f, name, usage)
}

// ZVDuration - duration flag.  returns pointer
func (ndf *NDFlagSet) ZVDuration(name string, example time.Duration, usage string) *time.Duration {
	var dv time.Duration
	ndf.ZVDurationVar(&dv, name, example, usage)
	return &dv
}

// ZVDurationVar - BYO duration pp version of ZVDuration
func (ndf *NDFlagSet) ZVDurationVar(dv *time.Duration, name string, example time.Duration, usage string) {
	d := &zvdff{dv: dv, example: example.String()}
	ndf.Var(d, name, usage)
}
