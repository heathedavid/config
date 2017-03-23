package config

import (
	"testing"
	"time"
)

func TestStringFlag(t *testing.T) {
	s := StringFlag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != "" {
		t.Error("Flag should initially be empty")
	}

	if s.String() != "" {
		t.Error("Flag should initially be empty")
	}

	s.Set("fred")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != "fred" {
		t.Error("Flag should be set to fred")
	}

	if s.String() != "fred" {
		t.Error("Flag should be set to fred")
	}
}

func TestIntFlag(t *testing.T) {
	s := IntFlag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != 0 {
		t.Error("Flag should initially be 0")
	}

	if s.String() != "0" {
		t.Error("Flag should initially be 0")
	}

	s.Set("-1")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != -1 {
		t.Error("Flag should be set to -1")
	}

	if s.String() != "-1" {
		t.Error("Flag should be set to -1")
	}
}

func TestInt64Flag(t *testing.T) {
	s := Int64Flag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != int64(0) {
		t.Error("Flag should initially be 0")
	}

	if s.String() != "0" {
		t.Error("Flag should initially be 0")
	}

	s.Set("-123456789")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != int64(-123456789) {
		t.Error("Flag should be set to -123456789")
	}

	if s.String() != "-123456789" {
		t.Error("Flag should be set to -123456789")
	}
}

func TestUint64Flag(t *testing.T) {
	s := Uint64Flag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != uint64(0) {
		t.Error("Flag should initially be 0")
	}

	if s.String() != "0" {
		t.Error("Flag should initially be 0")
	}

	s.Set("123456789")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != uint64(123456789) {
		t.Error("Flag should be set to 123456789")
	}

	if s.String() != "123456789" {
		t.Error("Flag should be set to 123456789")
	}
}

func TestBoolFlag(t *testing.T) {
	s := BoolFlag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() == true {
		t.Error("Flag should initially be false")
	}

	if s.String() != "false" {
		t.Error("Flag should initially be false")
	}

	s.Set("true")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() == false {
		t.Error("]Flag should be set to true")
	}

	if s.String() != "true" {
		t.Error("Flag should be set to true")
	}
}

func TestFloat64Flag(t *testing.T) {
	s := Float64Flag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != float64(0) {
		t.Error("Flag should initially be 0")
	}

	if s.String() != "0" {
		t.Error("Flag should initially be 0")
	}

	s.Set("12345.6789")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != float64(12345.6789) {
		t.Error("Flag should be set to 12345.6789")
	}

	if s.String() != "12345.6789" {
		t.Error("Flag should be set to 12345.6789")
	}
}

func TestDurationFlag(t *testing.T) {
	s := DurationFlag{}

	if s.IsSet() {
		t.Error("Flag should not be set on initialization")
	}

	if s.Get() != time.Duration(0) {
		t.Error("Flag should initially be 0")
	}

	if s.String() != "0s" {
		t.Error("Flag should initially be 0")
	}

	s.Set("10s")

	if !s.IsSet() {
		t.Error("Flag should be set")
	}

	if s.Get() != time.Duration(time.Second*10) {
		t.Error("Flag should be set to 10s")
	}

	if s.String() != "10s" {
		t.Error("Flag should be set to 10s")
	}
}
