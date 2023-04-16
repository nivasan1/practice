package utils

type Ord[C comparable] interface {
	Less(j C) bool
	Val() C
}

type OrdInt struct {
	i int
}

func NewOrdInt(i int) OrdInt {
	return OrdInt{i}
}

func (o OrdInt) Less(j int) bool {
	return o.i < j
}

func (o OrdInt) Val() int {
	return o.i
}

func Log2(exp int) int {
	if exp == 1 {
		return 0
	}
	// other wise
	return 1 + Log2(exp>>1)
}

func IsPow2(exp int) bool {
	if exp == 1 {
		return true
	}
	if exp%2 != 0 {
		return false
	}
	return IsPow2(exp >> 1)
}
