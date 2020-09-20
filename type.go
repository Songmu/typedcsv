package typedcsv

import "time"

type Type int

const (
	TypeAny Type = iota
	TypeString
	TypeNumber
	TypeBool
	TypeNull
	TypeTime
)

func String(str string) Value {
	return &valueString{value: str}
}

func Float(f float64) Value {
	return &valueNumber{float: f}
}

func Int(i int64) Value {
	return &valueNumber{
		isInt: true,
		int:   i,
	}
}

func Bool(bool bool) Value {
	return &valueBool{bool: bool}
}

var null *valueNull

func Null() Value {
	return null
}

func Time(ti time.Time) Value {
	return &valueTime{time: ti}
}
