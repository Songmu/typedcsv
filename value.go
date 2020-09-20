package typedcsv

type Value interface {
	RawValue() string
	Value() string
	Quoted() bool
	Valid() bool
}

type valueBackend struct {
	rawValue string
	value    string
	quoted   bool
}

var _ Value = (*valueBackend)(nil)

func (vb *valueBackend) RawValue() string {
	return vb.rawValue
}

func (vb *valueBackend) Value() string {
	return vb.value
}

func (vb *valueBackend) Quoted() bool {
	return vb.quoted
}

func (vb *valueBackend) Valid() bool {
	return true
}
