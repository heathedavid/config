/*
Collection of flags which contain notion of whether flag has been set.
*/
package config

import (
	"flag"
	"strconv"
	"time"
)

// Combines flag.Getter (which includes flag.Value) and helper to determine if
// flag was set. All flags below will implement this.
type ConfigFlag interface {
	flag.Getter
	IsSet() bool
}

// String flag
type StringFlag struct {
	set   bool
	value string
}

func (f *StringFlag) Set(x string) error {
	f.value = x
	f.set = true
	return nil
}

func (f *StringFlag) String() string {
	return f.value
}

func (f *StringFlag) Get() interface{} {
	return f.value
}

func (f *StringFlag) IsSet() bool {
	return f.set
}

// Int flag
type IntFlag struct {
	set   bool
	value int
}

func (f *IntFlag) Set(x string) error {
	value, err := strconv.ParseInt(x, 0, 64)
	f.value = int(value)
	f.set = true
	return err
}

func (f *IntFlag) String() string {
	return strconv.Itoa(f.value)
}

func (f *IntFlag) Get() interface{} {
	return f.value
}

func (f *IntFlag) IsSet() bool {
	return f.set
}

// Int64 flag
type Int64Flag struct {
	set   bool
	value int64
}

func (f *Int64Flag) Set(x string) error {
	value, err := strconv.ParseInt(x, 0, 64)
	f.value = value
	f.set = true
	return err
}

func (f *Int64Flag) String() string {
	return strconv.FormatInt(int64(f.value), 10)
}

func (f *Int64Flag) Get() interface{} {
	return f.value
}

func (f *Int64Flag) IsSet() bool {
	return f.set
}

// Unsigned Uint flag
type UintFlag struct {
	set   bool
	value uint
}

func (f *UintFlag) Set(x string) error {
	value, err := strconv.ParseUint(x, 0, 64)
	f.value = uint(value)
	f.set = true
	return err
}

func (f *UintFlag) String() string {
	return strconv.FormatUint(uint64(f.value), 10)
}

func (f *UintFlag) Get() interface{} {
	return f.value
}

func (f *UintFlag) IsSet() bool {
	return f.set
}

// Unsigned Int64 flag
type Uint64Flag struct {
	set   bool
	value uint64
}

func (f *Uint64Flag) Set(x string) error {
	value, err := strconv.ParseUint(x, 0, 64)
	f.value = value
	f.set = true
	return err
}

func (f *Uint64Flag) String() string {
	return strconv.FormatUint(f.value, 10)
}

func (f *Uint64Flag) Get() interface{} {
	return f.value
}

func (f *Uint64Flag) IsSet() bool {
	return f.set
}

// Bool flag
type BoolFlag struct {
	set   bool
	value bool
}

func (f *BoolFlag) Set(x string) error {
	value, err := strconv.ParseBool(x)
	f.value = value
	f.set = true
	return err
}

func (f *BoolFlag) String() string {
	return strconv.FormatBool(f.value)
}

func (f *BoolFlag) Get() interface{} {
	return f.value
}

func (f *BoolFlag) IsSet() bool {
	return f.set
}

// Float64 flag
type Float64Flag struct {
	set   bool
	value float64
}

func (f *Float64Flag) Set(x string) error {
	value, err := strconv.ParseFloat(x, 64)
	f.value = value
	f.set = true
	return err
}

func (f *Float64Flag) String() string {
	return strconv.FormatFloat(f.value, 'g', -1, 64)
}

func (f *Float64Flag) Get() interface{} {
	return f.value
}

func (f *Float64Flag) IsSet() bool {
	return f.set
}

// Duration flag
type DurationFlag struct {
	set   bool
	value time.Duration
}

func (f *DurationFlag) Set(x string) error {
	value, err := time.ParseDuration(x)
	f.value = value
	f.set = true
	return err
}

func (f *DurationFlag) String() string {
	return f.value.String()
}

func (f *DurationFlag) Get() interface{} {
	return f.value
}

func (f *DurationFlag) IsSet() bool {
	return f.set
}
