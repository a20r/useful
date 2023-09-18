package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

//go:generate go run main.go

const maxIn = 10
const maxOut = 10

func main() {
	f := jen.NewFile("funcs")

	for i := 0; i <= maxIn; i++ {
		for j := 0; j <= maxOut; j++ {
			t := types(f, i, j)
			params(t, "X", i)
			params(t, "Y", j)
		}
	}

	if err := f.Save("../types.go"); err != nil {
		panic(err)
	}
}

func types(f *jen.File, in, out int) *jen.Statement {
	name := fmt.Sprintf("Fn%dx%d", in, out)
	types := make([]jen.Code, in+out)

	for i := 0; i < in; i++ {
		types[i] = jen.Id(fmt.Sprintf("X%d", i))
	}

	for i := 0; i < out; i++ {
		types[in+i] = jen.Id(fmt.Sprintf("Y%d", i))
	}

	if len(types) > 0 {
		types[len(types)-1].(*jen.Statement).Id("any")
	}

	return f.Type().Id(name).Types(types...).Func()
}

func params(f *jen.Statement, pre string, num int) {
	ps := make([]jen.Code, num)

	for i := 0; i < num; i++ {
		ps[i] = jen.Id(fmt.Sprintf("%s%d", strings.ToLower(pre), i)).Id(fmt.Sprintf("%s%d", strings.ToUpper(pre), i))
	}

	f.Params(ps...)
}
