package config

import (
	"reflect"
	"testing"
	"time"
)

func TestStringValue1(t *testing.T) {
	ss := struct {
		Field string
	}{
		"A",
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := StringFlag{}

	v := StringValue{false, "", &flag, f}

	v.Set()

	if ss.Field != "A" {
		t.Error("No default, no flag, should be initial value A")
	}
}

func TestStringValue2(t *testing.T) {
	ss := struct {
		Field string
	}{
		"A",
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := StringFlag{}

	v := StringValue{true, "fred", &flag, f}

	v.Set()

	if ss.Field != "fred" {
		t.Error("default, no flag, should be default value fred")
	}
}

func TestStringValue3(t *testing.T) {
	ss := struct {
		Field string
	}{
		"A",
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := StringFlag{}

	v := StringValue{false, "", &flag, f}

	flag.Set("bert")

	v.Set()

	if ss.Field != "bert" {
		t.Error("No default, flag bert, should be flag value bert")
	}
}

func TestInt64Value1(t *testing.T) {
	ss := struct {
		Field int64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Int64Flag{}

	v := Int64Value{false, 0, &flag, f}

	v.Set()

	if ss.Field != 1 {
		t.Error("No default, no flag, should be initial value 1")
	}
}

func TestInt64Value2(t *testing.T) {
	ss := struct {
		Field int64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Int64Flag{}

	v := Int64Value{true, -23456789, &flag, f}

	v.Set()

	if ss.Field != -23456789 {
		t.Error("default, no flag, should be default value -23456789")
	}
}

func TestInt64Value3(t *testing.T) {
	ss := struct {
		Field int64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Int64Flag{}

	v := Int64Value{false, 0, &flag, f}

	flag.Set("-3456789")

	v.Set()

	if ss.Field != -3456789 {
		t.Error("No default, flag -3456789, should be flag value -3456789")
	}
}

func TestUint64Value1(t *testing.T) {
	ss := struct {
		Field uint64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Uint64Flag{}

	v := Uint64Value{false, 0, &flag, f}

	v.Set()

	if ss.Field != 1 {
		t.Error("No default, no flag, should be initial value 1")
	}
}

func TestUint64Value2(t *testing.T) {
	ss := struct {
		Field uint64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Uint64Flag{}

	v := Uint64Value{true, 23456789, &flag, f}

	v.Set()

	if ss.Field != 23456789 {
		t.Error("default, no flag, should be default value 23456789")
	}
}

func TestUint64Value3(t *testing.T) {
	ss := struct {
		Field uint64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Uint64Flag{}

	v := Uint64Value{false, 0, &flag, f}

	flag.Set("3456789")

	v.Set()

	if ss.Field != 3456789 {
		t.Error("No default, flag 3456789, should be initial value 3456789")
	}
}

func TestBoolValue1(t *testing.T) {
	ss := struct {
		Field bool
	}{
		true,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := BoolFlag{}

	v := BoolValue{false, false, &flag, f}

	v.Set()

	if !ss.Field {
		t.Error("No default, no flag, should be initial value true")
	}
}

func TestBoolValue2(t *testing.T) {
	ss := struct {
		Field bool
	}{
		false,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := BoolFlag{}

	v := BoolValue{true, true, &flag, f}

	v.Set()

	if !ss.Field {
		t.Error("default, no flag, should be default value true")
	}
}

func TestBoolValue3(t *testing.T) {
	ss := struct {
		Field bool
	}{
		false,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := BoolFlag{}

	v := BoolValue{false, false, &flag, f}

	flag.Set("true")

	v.Set()

	if !ss.Field {
		t.Error("No default, flag true, should be flag true")
	}
}

func TestFloat64Value1(t *testing.T) {
	ss := struct {
		Field float64
	}{
		1.0,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Float64Flag{}

	v := Float64Value{false, 0, &flag, f}

	v.Set()

	if ss.Field != 1.0 {
		t.Error("No default, no flag, should be initial value 1.0")
	}
}

func TestFloat64Value2(t *testing.T) {
	ss := struct {
		Field float64
	}{
		1.0,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Float64Flag{}

	v := Float64Value{true, -2.345, &flag, f}

	v.Set()

	if ss.Field != -2.345 {
		t.Error("default, no flag, should be default value -2.345")
	}
}

func TestFloat64Value3(t *testing.T) {
	ss := struct {
		Field float64
	}{
		1,
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := Float64Flag{}

	v := Float64Value{false, 0, &flag, f}

	flag.Set("-3.456")

	v.Set()

	if ss.Field != -3.456 {
		t.Error("No default, flag -3.456, should be initial value -3.456")
	}
}

func TestDurationValue1(t *testing.T) {
	ss := struct {
		Field time.Duration
	}{
		time.Duration(1),
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := DurationFlag{}

	v := DurationValue{false, time.Duration(0), &flag, f}

	v.Set()

	if ss.Field != time.Duration(1) {
		t.Error("No default, no flag, should be initial value 1ns")
	}
}

func TestDurationValue2(t *testing.T) {
	ss := struct {
		Field time.Duration
	}{
		time.Duration(1),
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := DurationFlag{}

	v := DurationValue{true, time.Duration(2), &flag, f}

	v.Set()

	if ss.Field != time.Duration(2) {
		t.Error("default, no flag, should be default value 2ns")
	}
}

func TestDurationValue3(t *testing.T) {
	ss := struct {
		Field time.Duration
	}{
		time.Duration(1),
	}

	f := reflect.ValueOf(&ss).Elem().Field(0)

	flag := DurationFlag{}

	v := DurationValue{false, time.Duration(0), &flag, f}

	flag.Set("3ns")

	v.Set()

	if ss.Field != time.Duration(3) {
		t.Error("No default, flag 3ns, should be flag 3ns")
	}
}
