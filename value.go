package typedcsv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	String() string
	Quoted() bool
}

type valueAny struct {
	value  string
	quoted bool
}

var _ Value = (*valueAny)(nil)

func (vb *valueAny) String() string {
	return vb.value
}

func (vb *valueAny) Quoted() bool {
	return vb.quoted
}

func (vb *valueAny) guess(loose bool) (Value, error) {
	if vb.quoted {
		return String(vb.value), nil
	}
	switch strings.ToLower(vb.value) {
	case "null":
		return Null(), nil
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

	if loose {
		return vb, nil
	}
	return nil, fmt.Errorf("can't guess the type for unquoted value: %s", vb.value)
}

func (vb *valueAny) mustGuess(loose bool) Value {
	v, err := vb.guess(loose)
	if err != nil {
		panic(err)
	}
	return v
}

type valueString struct {
	value string
}

var _ Value = (*valueString)(nil)

func (vb *valueString) String() string {
	return vb.value
}

func (vb *valueString) Quoted() bool {
	return true
}

func String(str string) Value {
	return &valueString{value: str}
}

type valueNull struct{}

var _ Value = (*valueNull)(nil)

func (vn *valueNull) String() string {
	return "null"
}

func (vn *valueNull) Quoted() bool {
	return false
}

var null *valueNull

func Null() Value {
	return null
}

type valueBool struct {
	bool bool
}

var _ Value = (*valueBool)(nil)

func (vn *valueBool) String() string {
	return fmt.Sprintf("%v", vn.bool)
}

func (vn *valueBool) Quoted() bool {
	return false
}

func Bool(bool bool) Value {
	return &valueBool{bool: bool}
}

type valueNumber struct {
	isInt    bool
	int      int64
	float    float64
	rawValue string
}

var _ Value = (*valueNumber)(nil)

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

func Int(i int64) Value {
	return &valueNumber{
		isInt: true,
		int:   i,
	}
}

func Float(f float64) Value {
	return &valueNumber{float: f}
}

type valueTime struct {
	time     time.Time
	rawValue string
}

var _ Value = (*valueTime)(nil)

func (vt *valueTime) String() string {
	if vt.rawValue != "" {
		return vt.rawValue
	}
	return vt.time.Format(time.RFC3339Nano)
}

func (vt *valueTime) Quoted() bool {
	return false
}

func Time(ti time.Time) Value {
	return &valueTime{time: ti}
}
