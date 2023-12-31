package assert_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/a20r/useful/assert"
	"github.com/sirupsen/logrus"
	tfassert "github.com/stretchr/testify/assert"
)

func TestAssertions(t *testing.T) {
	logrus.StandardLogger().ExitFunc = func(i int) {
		t.Logf("Logrus handled exit code %d", i)
	}
	t.Run("NotEmpty and Positive", func(t *testing.T) {
		as := tfassert.New(t)
		sum := func(in []int) (total int, err error) {
			defer assert.Handle(&err)
			assert.NotEmpty(in)

			for _, v := range in {
				assert.Positive(v)
				total += v
			}

			return total, nil
		}

		_, err := sum(nil)
		as.ErrorIs(err, assert.ErrContainerIsEmpty)

		_, err = sum([]int{42, 1, 2, -1})
		as.ErrorIs(err, assert.ErrValueIsNotPositive)

		_, err = sum([]int{0, 1, 2})
		as.ErrorIs(err, assert.ErrValueIsNotPositive)

		_, err = sum([]int{3, 1, 2})
		as.NoError(err)
	})

	t.Run("Success", func(t *testing.T) {
		as := tfassert.New(t)
		readAll := func(fname string) (contents []byte, err error) {
			defer assert.Handle(&err)
			f := assert.Success(os.Open(fname))
			contents = assert.Success(io.ReadAll(f))
			return
		}

		_, err := readAll("does-not-exist.txt")
		as.ErrorIs(err, assert.ErrFuncReturnedError)
	})

	t.Run("SameSize", func(t *testing.T) {
		as := tfassert.New(t)
		dot := func(a, b []float64) (product float64, err error) {
			defer assert.Handle(&err)
			assert.SameSize(a, b)

			for i := range a {
				product += a[i] * b[i]
			}

			return
		}

		_, err := dot([]float64{2, 4, 8}, []float64{8, 4, 2})
		as.NoError(err)

		_, err = dot([]float64{2, 4, 8}, []float64{8, 4})
		as.ErrorIs(err, assert.ErrSlicesAreDifferentSizes)

		_, err = dot([]float64{8}, []float64{8, 4, 2})
		as.ErrorIs(err, assert.ErrSlicesAreDifferentSizes)
	})

	t.Run("HasKeys", func(t *testing.T) {
		as := tfassert.New(t)
		formatName := func(m map[string]string) (name string, err error) {
			defer assert.Handle(&err)
			assert.HasKeys(m, "firstName", "lastName")
			return fmt.Sprintf("%s, %s\n", m["lastName"], m["firstName"]), nil
		}

		_, err := formatName(map[string]string{"firstName": "Alex", "lastName": "Wallar"})
		as.NoError(err)

		_, err = formatName(map[string]string{"firstName": "Alan"})
		as.ErrorIs(err, assert.ErrKeyNotInMap)

		_, err = formatName(map[string]string{"lastName": "Wallar"})
		as.ErrorIs(err, assert.ErrKeyNotInMap)

		_, err = formatName(map[string]string{"whatever": "yo"})
		as.ErrorIs(err, assert.ErrKeyNotInMap)

		_, err = formatName(map[string]string{})
		as.ErrorIs(err, assert.ErrKeyNotInMap)
	})

	t.Run("CatchPanic and NotZero", func(t *testing.T) {
		as := tfassert.New(t)
		div := func(l ...float64) (v float64, err error) {
			defer assert.Handle(&err)
			defer assert.CatchPanic(&err)
			assert.NotZero(l[1])
			return l[0] / l[1], nil
		}

		_, err := div(1)
		as.ErrorIs(err, assert.ErrPanic)

		_, err = div(1, 0)
		as.ErrorIs(err, assert.ErrValueIsZero)

		_, err = div(1, 2)
		as.NoError(err)
	})

	t.Run("FatalOnError", func(t *testing.T) {
		assertionWillFail := func() (err error) {
			defer assert.FatalOnError(&err)
			assert.NotNil[int](nil)
			return err
		}

		assertionWillFail()
	})
}
