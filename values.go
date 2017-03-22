/*
Values are used to store the information for the fields within the structure
we are trying to initialize - the flag package only allows a name to be used
once - Values are used to remove this limitation.

If a flag is not set and there is not default value the zero value will be used
for the field.
*/
package config

import (
  "reflect"
  "time"
)

// Each field within a struct which is to be initialized from a flag will have
// its own SetValue.
type SetValue interface {
  Set()
}

// String Value
type StringValue struct {
	isDefVal 	bool
	defVal		string
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *StringValue) Set() {
	if v.flag.IsSet() {
	   v.t.SetString(v.flag.Get().(string))
	}	else if v.isDefVal {
		v.t.SetString(v.defVal)
	}
}

// Int Value
type IntValue struct {
	isDefVal 	bool
	defVal		int
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *IntValue) Set() {
	if v.flag.IsSet() {
	   v.t.SetInt(v.flag.Get().(int64))
	}	else if v.isDefVal {
		v.t.SetInt(int64(v.defVal))
	}
}

// Int64 Value
type Int64Value struct {
	isDefVal 	bool
	defVal		int64
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *Int64Value) Set() {
	if v.flag.IsSet() {
	   v.t.SetInt(v.flag.Get().(int64))
	}	else if v.isDefVal {
		v.t.SetInt(v.defVal)
	}
}

// Uint64 Value
type Uint64Value struct {
	isDefVal 	bool
	defVal		uint64
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *Uint64Value) Set() {
	if v.flag.IsSet() {
	   v.t.SetUint(v.flag.Get().(uint64))
	}	else if v.isDefVal {
		v.t.SetUint(v.defVal)
	}
}

// Bool Value
type BoolValue struct {
	isDefVal 	bool
	defVal		bool
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *BoolValue) Set() {
	if v.flag.IsSet() {
	   v.t.SetBool(v.flag.Get().(bool))
	}	else if v.isDefVal {
		v.t.SetBool(v.defVal)
	}
}

// Float64 Value
type Float64Value struct {
	isDefVal 	bool
	defVal		float64
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *Float64Value) Set() {
	if v.flag.IsSet() {
	   v.t.SetFloat(v.flag.Get().(float64))
	}	else if v.isDefVal {
		v.t.SetFloat(v.defVal)
	}
}

// Duration Value
type DurationValue struct {
	isDefVal 	bool
	defVal		time.Duration
	flag			ConfigFlag
	t 				reflect.Value
}

func (v *DurationValue) Set() {
	if v.flag.IsSet() {
	   v.t.Set(reflect.ValueOf(v.flag.Get()))
	}	else if v.isDefVal {
		v.t.Set(reflect.ValueOf(v.defVal))
	}
}
