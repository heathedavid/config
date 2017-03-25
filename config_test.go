package config

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestStringParse1(t *testing.T) {
	ss := struct {
		Field string
	}{
		"A",
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var1", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var1"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != "A" {
			t.Error("Initial value should be A")
		}
		sv.Set()
		if ss.Field != "A" {
			t.Error("Should be initial value A")
		}
	}
}

func TestStringParse2(t *testing.T) {
	ss := struct {
		Field string
	}{
		"A",
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var2", "B", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var2"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != "A" {
			t.Error("Initial value should be A")
		}
		sv.Set()
		if ss.Field != "B" {
			t.Error("Should be default value B")
		}
	}
}

func TestIntParse1(t *testing.T) {
	ss := struct {
		Field int
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var3", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var3"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != 1 {
			t.Error("Should be initial value 1")
		}
	}
}

func TestIntParse2(t *testing.T) {
	ss := struct {
		Field int
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var4", "2", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var4"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != 2 {
			t.Error("Should be default value 2")
		}
	}
}

func TestInt64Parse1(t *testing.T) {
	ss := struct {
		Field int64
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var5", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var5"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != 1 {
			t.Error("Should be initial value 1")
		}
	}
}

func TestInt64Parse2(t *testing.T) {
	ss := struct {
		Field int64
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var6", "-23456789", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var6"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != -23456789 {
			t.Error("Should be default value -23456789")
		}
	}
}

func TestUint64Parse1(t *testing.T) {
	ss := struct {
		Field uint64
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var7", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var7"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != 1 {
			t.Error("Should be initial value 1")
		}
	}
}

func TestUint64Parse2(t *testing.T) {
	ss := struct {
		Field uint64
	}{
		1,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var8", "23456789", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var8"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1 {
			t.Error("Initial value should be 1")
		}
		sv.Set()
		if ss.Field != 23456789 {
			t.Error("Should be default value 23456789")
		}
	}
}

func TestFloat64Parse1(t *testing.T) {
	ss := struct {
		Field float64
	}{
		1.0,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var9", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var9"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1.0 {
			t.Error("Initial value should be 1.0")
		}
		sv.Set()
		if ss.Field != 1.0 {
			t.Error("Should be initial value 1.0")
		}
	}
}

func TestFloat64Parse2(t *testing.T) {
	ss := struct {
		Field float64
	}{
		1.0,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var10", "2.345", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var10"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != 1.0 {
			t.Error("Initial value should be 1.0")
		}
		sv.Set()
		if ss.Field != 2.345 {
			t.Error("Should be default value 2.345")
		}
	}
}

func TestBoolParse1(t *testing.T) {
	ss := struct {
		Field bool
	}{
		true,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var11", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var11"]; !ok {
			t.Error("Flag not created")
		}

		if !ss.Field {
			t.Error("Initial value should be true")
		}
		sv.Set()
		if !ss.Field {
			t.Error("Should be initial value true")
		}
	}
}

func TestBoolParse2(t *testing.T) {
	ss := struct {
		Field bool
	}{
		true,
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var12", "false", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var12"]; !ok {
			t.Error("Flag not created")
		}

		if !ss.Field {
			t.Error("Initial value should be true")
		}
		sv.Set()
		if ss.Field {
			t.Error("Should be default value false")
		}
	}
}

func TestDurationParse1(t *testing.T) {
	ss := struct {
		Field time.Duration
	}{
		time.Duration(1),
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var13", "", "test env var", false)

	if err != nil {
		t.Error("No default to parse")
	} else {
		if _, ok := m["var13"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != time.Duration(1) {
			t.Error("Initial value should be 1ns")
		}
		sv.Set()
		if ss.Field != time.Duration(1) {
			t.Error("Should be initial value 1ns")
		}
	}
}

func TestDurationParse2(t *testing.T) {
	ss := struct {
		Field time.Duration
	}{
		time.Duration(1),
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	sv, err := parseDefault(m, f, "var14", "2ns", "test env var", true)

	if err != nil {
		t.Error("Default parse issue")
	} else {
		if _, ok := m["var14"]; !ok {
			t.Error("Flag not created")
		}

		if ss.Field != time.Duration(1) {
			t.Error("Initial value should be 1ns")
		}
		sv.Set()
		if ss.Field != time.Duration(2) {
			t.Error("Should be default value 2ns")
		}
	}
}

type MyType struct {
	Field int
}

func TestMyTypeParse2(t *testing.T) {
	ss := struct {
		Field MyType
	}{
		MyType{3},
	}

	m := make(map[string]ConfigFlag)

	f := reflect.ValueOf(&ss).Elem().Field(0)

	_, err := parseDefault(m, f, "var15", "", "test env var", false)

	if err == nil {
		t.Error("Should have fialed as myType is unknown")
	}
}
