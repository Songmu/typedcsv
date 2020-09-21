package typedcsv

import "time"

type Type int

const (
	TypeAny Type = iota
	TypeHeader
	TypeString
	TypeNumber
	TypeBool
	TypeNull
	TypeTime
)

func (t Type) String() string {
	switch t {
	case TypeAny:
		return "any"
	case TypeHeader:
		return "header"
	case TypeString:
		return "string"
	case TypeNumber:
		return "number"
	case TypeBool:
		return "bool"
	case TypeNull:
		return "null"
	case TypeTime:
		return "time"
	}
	return "unknown"
}

func (t Type) compatible(t2 Type) bool {
	return t == TypeNull || t2 == TypeNull || t == t2
}

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
