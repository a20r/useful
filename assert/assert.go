package assert

import (
	"errors"

	"github.com/a20r/falta"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/constraints"
)

var (
	ErrAssertionFailed         = falta.Newf("assert: [%s] assertion failed")
	ErrValueIsNil              = falta.Newf("<%T> is nil")
	ErrContainerIsEmpty        = falta.Newf("<%T> container is empty")
	ErrValueIsNegative         = falta.Newf("%v < 0")
	ErrValueIsNotPositive      = falta.Newf("%v <= 0")
	ErrValueIsZero             = falta.Newf("value == 0")
	ErrErrorIsNotNil           = falta.Newf("err is not nil: %v")
	ErrSlicesAreDifferentSizes = falta.Newf("slices are different sizes: %v != %v")
	ErrFuncReturnedError       = falta.Newf("<%T> no error expected: %v")
	ErrKeyNotInMap             = falta.Newf(`keys not in map: %v not in <%T>`)
	ErrPanic                   = falta.Newf("panic: %s")
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
		panic(ErrAssertionFailed.New("NotZero").Wrap(ErrValueIsZero.New()))
	}
}

func Positive[T constraints.Integer | constraints.Float](val T) {
	if val <= 0 {
		panic(ErrAssertionFailed.New("Positive").Wrap(ErrValueIsNotPositive.New(val)))
	}
}

func Success[T any](val T, err error) T {
	if err != nil {
		panic(ErrAssertionFailed.New("Success").Wrap(ErrFuncReturnedError.New(err, err).Wrap(err)))
	}

	return val
}

func NoError(err error) {
	if err != nil {
		panic(ErrAssertionFailed.New("NoError").Wrap(ErrErrorIsNotNil.New(err)))
	}
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

func SameSize[T, U any](a []T, b []U) {
	if len(a) != len(b) {
		panic(ErrAssertionFailed.New("SameSize").Wrap(ErrSlicesAreDifferentSizes.New(len(a), len(b))))
	}
}

func Handle(err *error, l ...*logrus.Entry) {
	if r := recover(); r != nil {
		errRecovered, ok := r.(error)

		// It's not an panic made by us
		if !ok || !errors.Is(errRecovered, ErrAssertionFailed) {
			panic(r)
		}

		if len(l) > 0 {
			l[0].WithError(errRecovered).Errorf("Assertion failed: %v", errRecovered)
		}

		*err = errRecovered
	}
}

func FatalOnError(err *error, l ...*logrus.Entry) {
	if r := recover(); r != nil {
		errRecovered, ok := r.(error)

		// It's not an panic made by us
		if !ok || !errors.Is(errRecovered, ErrAssertionFailed) {
			panic(r)
		}

		*err = errRecovered
		log := logrus.WithField("func", "a20r/useful/assert.FatalOnError")

		if len(l) > 0 {
			log = l[0]
		}

		log.WithError(*err).Fatalf("[assert] Assertion failed: %v", *err)
	}
}

func CatchPanic(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok && errors.Is(e, ErrAssertionFailed) {
			panic(e)
		}

		*err = ErrAssertionFailed.New("NoPanic").Wrap(ErrPanic.New(r))
	}
}
