package core

import (
	"strings"
	"errors"
)

type Bool uint8

const (
	BOOL_UNKNOWN Bool = iota
	BOOL__TRUE    // true
	BOOL__FALSE   // false
)

func (v Bool) True() bool {
	return v == BOOL__TRUE
}

var InvalidBool = errors.New("invalid Bool")

func ParseBoolFromBoolValueString(s string) (Bool, error) {
	switch s {
	case "":
		return BOOL_UNKNOWN, nil
	case "false":
		return BOOL__FALSE, nil
	case "true":
		return BOOL__TRUE, nil
	}
	return BOOL_UNKNOWN, InvalidBool
}

func (v Bool) String() string {
	switch v {
	case BOOL_UNKNOWN:
		return ""
	case BOOL__FALSE:
		return "false"
	case BOOL__TRUE:
		return "true"
	}
	return "UNKNOWN"
}

func (v Bool) MarshalJSON() ([]byte, error) {
	str := v.String()
	if str == "" {
		return []byte("null"), nil
	}
	if str == "UNKNOWN" {
		return nil, InvalidBool
	}
	return []byte(str), nil
}

func (v *Bool) UnmarshalJSON(data []byte) (err error) {
	s := strings.ToLower(strings.Trim(string(data), `"`))
	if s == "null" {
		*v = BOOL_UNKNOWN
		return
	}
	*v, err = ParseBoolFromBoolValueString(s)
	return
}
