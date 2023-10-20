package assert_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/a20r/useful/assert"
)

func TestAssertions(t *testing.T) {
	sum := func(in []int) (total int, err error) {
		defer assert.Handle(&err)
		assert.NotEmpty(in)

		for _, v := range in {
			assert.Positive(v)
			total += v
		}

		return total, nil
	}

	fmt.Println(sum(nil))
	fmt.Println(sum([]int{42, 1, 2, -1}))
	fmt.Println(sum([]int{0, 1, 2}))
	fmt.Println(sum([]int{3, 1, 2}))

	readAll := func(fname string) (contents []byte, err error) {
		defer assert.Handle(&err)
		f := assert.NoError(os.Open(fname))
		contents = assert.NoError(io.ReadAll(f))
		return
	}

	fmt.Println(readAll("does-not-exist.txt"))

	dot := func(a, b []float64) (product float64, err error) {
		defer assert.Handle(&err)
		assert.SameSize(a, b)

		for i := range a {
			product += a[i] * b[i]
		}

		return
	}

	fmt.Println(dot([]float64{2, 4, 8}, []float64{8, 4, 2}))
	fmt.Println(dot([]float64{2, 4, 8}, []float64{8, 4}))
	fmt.Println(dot([]float64{8}, []float64{8, 4, 2}))

	printName := func(m map[string]string) (err error) {
		defer assert.Handle(&err)
		assert.HasKeys(m, "firstName", "lastName")
		fmt.Printf("%s, %s", m["lastName"], m["firstName"])
		return
	}

	fmt.Println(printName(map[string]string{"firstName": "Alex", "lastName": "Wallar"}))
	fmt.Println(printName(map[string]string{"firstName": "Alan"}))
	fmt.Println(printName(map[string]string{"lastName": "Wallar"}))
	fmt.Println(printName(map[string]string{"whatever": "yo"}))
	fmt.Println(printName(map[string]string{}))
}
