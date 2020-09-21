package typedcsv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	Type() Type
	String() string
	Quoted() bool
}

type valueAny struct {
	value  string
	quoted bool
}

var _ Value = (*valueAny)(nil)

func (vb *valueAny) Type() Type {
	return TypeAny
}

func (vb *valueAny) String() string {
	return vb.value
}

func (vb *valueAny) Quoted() bool {
	return vb.quoted
}

func (vb *valueAny) guess(strict bool) (Value, error) {
	if vb.quoted {
		return String(vb.value), nil
	}
	switch strings.ToLower(vb.value) {
	case "", "null":
		return &valueNull{rawValue: &vb.value}, nil
	case "true", "false":
		return Bool(strings.ToLower(vb.value) == "true"), nil
	}

	i, err := strconv.ParseInt(vb.value, 10, 64)
	if err == nil {
		return &valueNumber{isInt: true, int: i, rawValue: vb.value}, nil
	}
	f, err := strconv.ParseFloat(vb.value, 64)
	if err == nil {
		return &valueNumber{float: f, rawValue: vb.value}, nil
	}

	ti, err := time.Parse(time.RFC3339Nano, vb.value)
	if err == nil {
		return &valueTime{time: ti, rawValue: vb.value}, nil
	}

	if !strict {
		return vb, nil
	}
	return nil, fmt.Errorf("can't guess the type for unquoted value: %s", vb.value)
}

type valueHeader struct {
	value string
	comma rune
}

var _ Value = (*valueHeader)(nil)

func (vh *valueHeader) Type() Type {
	return TypeHeader
}

func (vh *valueHeader) String() string {
	return vh.value
}

func (vh *valueHeader) Quoted() bool {
	return fieldNeedsQuotes(vh.value, vh.comma)
}

type valueString struct {
	value string
}

var _ Value = (*valueString)(nil)

func (vs *valueString) Type() Type {
	return TypeString
}

func (vs *valueString) String() string {
	return vs.value
}

func (vs *valueString) Quoted() bool {
	return true
}

type valueNull struct {
	rawValue *string
}

var _ Value = (*valueNull)(nil)

func (vn *valueNull) Type() Type {
	return TypeNull
}

func (vn *valueNull) String() string {
	if vn.rawValue != nil {
		return *vn.rawValue
	}
	return "null"
}

func (vn *valueNull) Quoted() bool {
	return false
}

type valueBool struct {
	bool bool
}

var _ Value = (*valueBool)(nil)

func (vb *valueBool) Type() Type {
	return TypeBool
}

func (vb *valueBool) String() string {
	return fmt.Sprintf("%v", vb.bool)
}

func (vb *valueBool) Quoted() bool {
	return false
}

type valueNumber struct {
	isInt    bool
	int      int64
	float    float64
	rawValue string
}

var _ Value = (*valueNumber)(nil)

func (vn *valueNumber) Type() Type {
	return TypeNumber
}

func (vn *valueNumber) String() string {
	if vn.rawValue != "" {
		return vn.rawValue
	}
	if vn.isInt {
		return fmt.Sprintf("%d", vn.int)
	}
	return fmt.Sprintf("%f", vn.float)
}

func (vn *valueNumber) Quoted() bool {
	return false
}

type valueTime struct {
	time     time.Time
	rawValue string
}

var _ Value = (*valueTime)(nil)

func (vt *valueTime) Type() Type {
	return TypeTime
}

func (vt *valueTime) String() string {
	if vt.rawValue != "" {
		return vt.rawValue
	}
	return vt.time.Format(time.RFC3339Nano)
}

func (vt *valueTime) Quoted() bool {
	return false
}
