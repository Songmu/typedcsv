package typedcsv

type Value interface {
	String() string
	Quoted() bool
}

type valueBackend struct {
	value  string
	quoted bool
}

var _ Value = (*valueBackend)(nil)

func (vb *valueBackend) String() string {
	return vb.value
}

func (vb *valueBackend) Quoted() bool {
	return vb.quoted
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
