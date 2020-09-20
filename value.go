package typedcsv

type Value interface {
	String() string
	Quoted() bool
	Valid() bool
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

func (vb *valueBackend) Valid() bool {
	return true
}
