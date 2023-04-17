package utils

const MAX_INT = 1<<63 - 1
const MIN_INT = -(1 << 63)

type Ord[C comparable] interface {
	Less(j C) bool
	LessOrd(j Ord[C]) bool
	Val() C
	Max() Ord[C]
	Min() Ord[C]
}

var x Ord[int] = OrdInt{}

type OrdInt struct {
	i int
}

func NewOrdInt(i int) OrdInt {
	return OrdInt{i}
}

func (o OrdInt) Less(j int) bool {
	return o.i < j
}

func (o OrdInt) LessOrd(j Ord[int]) bool {
	return o.Less(j.Val())
}

func (o OrdInt) Val() int {
	return o.i
}

func (o OrdInt) Max() Ord[int] {
	return NewOrdInt(MAX_INT)
}

func (o OrdInt) Min() Ord[int] {
	return NewOrdInt(MIN_INT)
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

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
