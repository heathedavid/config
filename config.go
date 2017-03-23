/*
config is used to initialize configuration structures from the struct field tags
or environmental variables.

env_name - if defined will be the name of the environmental var (else field name)
env_def  - if defined is used for a string representation of the default value
 					 if not defined and the environment variable is not defined the
					 zero value for the field will be used.
env_desc - A description of the environmental var
 */
package config

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type setters struct {
	s		[]SetValue
}

// Pass a point to the annotated structures you want to initialize
func Load(structs ...interface{}) error {
	if (flag.Parsed()) {
		return fmt.Errorf("Load must be called before a call to flag.Pars")
	}

	m := make(map[string]ConfigFlag)

	numStructs := len(structs)

	fieldSetters := make([]setters, numStructs)

	for i := 0; i < numStructs; i++ {
		s := structs[i]

		t := reflect.ValueOf(s).Elem()
		typeOfT := t.Type()

		fieldSetters[i].s = make([]SetValue, t.NumField())

		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)

			tag := typeOfT.Field(i).Tag

			var ok bool
			var name string
			if name, ok = tag.Lookup("env_name"); !ok {
				name = typeOfT.Field(i).Name
			}

			defVal, isDefVal := tag.Lookup("env_def")

			var desc string
			if desc, ok = tag.Lookup("env_desc"); !ok {
				desc = name
			}

			var err error
			fieldSetters[i].s[j], err = parseDefault(m, f, name, defVal, desc, isDefVal)

			if err != nil {
				return nil
			}
		}
	}

	flag.Parse()
	for i := 0; i < numStructs; i++ {
		for j := 0; j < len(fieldSetters[i].s); j++ {
			fieldSetters[i].s[j].Set()
		}
	}
	return nil
}

var stringType = reflect.TypeOf((*string)(nil)).Elem()
var intType = reflect.TypeOf((*int)(nil)).Elem()
var int64Type = reflect.TypeOf((*int64)(nil)).Elem()
var uint64Type = reflect.TypeOf((*uint64)(nil)).Elem()
var float64Type = reflect.TypeOf((*float64)(nil)).Elem()
var boolType = reflect.TypeOf((*bool)(nil)).Elem()
var durationType = reflect.TypeOf((*time.Duration)(nil)).Elem()

func parseDefault(m map[string]ConfigFlag, t reflect.Value, name, defVal,
	desc string, isDefVal bool) (SetValue, error) {

	var err error
	switch t.Type() {
	case stringType:
		var def string
		if isDefVal {
			def = defVal
		}
		if cf, ok := m[name]; ok {
			val := StringValue{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := StringFlag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := StringValue{isDefVal, def, &cf, t}
			return &val, nil
		}

	case intType:
		var def int
		if isDefVal {
			if def, err = strconv.Atoi(defVal); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := IntValue{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := IntFlag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := IntValue{isDefVal, def, &cf, t}
			return &val, nil
		}

	case int64Type:
		var def int64
		if isDefVal {
			if def, err = strconv.ParseInt(defVal, 0, 64); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := Int64Value{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := Int64Flag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := Int64Value{isDefVal, def, &cf, t}
			return &val, nil
		}

	case uint64Type:
		var def uint64
		if isDefVal {
			if def, err = strconv.ParseUint(defVal, 0, 64); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := Uint64Value{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := Uint64Flag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := Uint64Value{isDefVal, def, &cf, t}
			return &val, nil
		}

	case float64Type:
		var def float64
		if isDefVal {
			if def, err = strconv.ParseFloat(defVal, 64); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := Float64Value{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := Float64Flag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := Float64Value{isDefVal, def, &cf, t}
			return &val, nil
		}

	case boolType:
		var def bool
		if isDefVal {
			if def, err = strconv.ParseBool(defVal); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := BoolValue{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := BoolFlag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := BoolValue{isDefVal, def, &cf, t}
			return &val, nil
		}

	case durationType:
		var def time.Duration
		if isDefVal {
			if def, err = time.ParseDuration(defVal); err != nil {
				return nil, err
			}
		}
		if cf, ok := m[name]; ok {
			val := DurationValue{isDefVal, def, cf, t}
			return &val, nil
		} else {
			cf := DurationFlag{}
			m[name] = &cf
			flag.Var(&cf, name, desc)
			val := DurationValue{isDefVal, def, &cf, t}
			return &val, nil
		}
	}

	return nil, fmt.Errorf("Unknow Type: %s", t.Type())
}
