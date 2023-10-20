package assert

import (
	"github.com/a20r/falta"
	"golang.org/x/exp/constraints"
)

var (
	ErrAssertionFailed         = falta.Newf("assert: [%s] assertion failed")
	ErrValueIsNil              = falta.Newf("<%T> is nil")
	ErrContainerIsEmpty        = falta.Newf("<%T> container is empty")
	ErrValueIsNegative         = falta.Newf("%v < 0")
	ErrValueIsNotPositive      = falta.Newf("%v <= 0")
	ErrValueIsZero             = falta.Newf("value == 0")
	ErrSlicesAreDifferentSizes = falta.Newf("slices are different sizes: %v != %v")
	ErrFuncReturnedError       = falta.Newf("<%T> no error expected")
	ErrKeyNotInMap             = falta.Newf(`keys not in map: %v not in <%T>`)
)

func NotNil[T any](val *T) {
	if val == nil {
		panic(ErrAssertionFailed.New("NotNil").Wrap(ErrValueIsNil.New(val)))
	}
}

func NotEmpty[T any](val []T) {
	if len(val) == 0 {
		panic(ErrAssertionFailed.New("NotEmpty").Wrap(ErrContainerIsEmpty.New(val)))
	}
}

func NotNegative[T constraints.Integer | constraints.Float](val T) {
	if val < 0 {
		panic(ErrAssertionFailed.New("NotNegative").Wrap(ErrValueIsNegative.New(val)))
	}
}

func NotZero[T constraints.Integer | constraints.Float](val T) {
	if val == 0 {
		panic(ErrAssertionFailed.New("NotZero").Wrap(ErrValueIsZero.New(val)))
	}
}

func Positive[T constraints.Integer | constraints.Float](val T) {
	if val <= 0 {
		panic(ErrAssertionFailed.New("Positive").Wrap(ErrValueIsNotPositive.New(val)))
	}
}

func NoError[T any](val T, err error) T {
	if err != nil {
		panic(ErrAssertionFailed.New("NoError").Wrap(ErrFuncReturnedError.New(err).Wrap(err)))
	}

	return val
}

func HasKeys[K comparable, V any, M ~map[K]V](m M, ks ...K) {
	var missing []K

	for _, k := range ks {
		if _, ok := m[k]; !ok {
			missing = append(missing, k)
		}
	}

	if len(missing) > 0 {
		panic(ErrAssertionFailed.New("HasKey").Wrap(ErrKeyNotInMap.New(missing, m)))
	}
}

func SameSize[T any](a, b []T) {
	if len(a) != len(b) {
		panic(ErrAssertionFailed.New("SameSize").Wrap(ErrSlicesAreDifferentSizes.New(len(a), len(b))))
	}
}

func Handle(err *error) {
	if r := recover(); r != nil {
		errRecovered := r.(error)
		*err = errRecovered
	}
}
