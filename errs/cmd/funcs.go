package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
)

// func NewFn0x1[Y0 any](fn errs.Fn0x2[Y0, error]) errs.Fn0x1[Y0] {
// 	return func(ctx context.Context) (y0 Y0) {
// 		inCtx, cancel := context.WithCancelCause(ctx)

// 		if err := inCtx.Err(); err != nil {
// 			return
// 		}

// 		y0, err := fn(ctx)

// 		if err != nil {
// 			cancel(err)
// 			return *new(Y0)
// 		}

// 		return y0
// 	}
// }

// funcsCmd represents the funcs command
var funcsCmd = &cobra.Command{
	Use: "funcs",
	Run: func(cmd *cobra.Command, args []string) {
		f := jen.NewFile("errs")

		for i := 0; i <= maxIn; i++ {
			for j := 0; j <= maxOut; j++ {
				t := types(f, i, j)
				params(t, i)
				returns(t, j)
			}
		}

		f.Line().Add(generateNewFn0x1())

		if err := f.Save("../types.go"); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(funcsCmd)
}

func generateNewFn0x1() jen.Code {
	return jen.Func().Id("NewFn0x1").Types(jen.Id("Y0").Id("any")).Params(
		jen.Id("fn").Id("Fn0x2").Types(jen.Id("Y0"), jen.Error()),
	).Params(jen.Id("Fn0x1").Types(jen.Id("Y0"))).Block(
		jen.Return(
			jen.Func().Params(
				jen.Id("ctx").Qual("context", "Context"),
			).Params(
				jen.Id("y0").Id("Y0"),
			).Block(
				jen.Id("inCtx").Op(",").Id("cancel").Op(":=").Qual("context", "WithCancelCause").Call(jen.Id("ctx")),
				jen.If(
					jen.Err().Op(":=").Id("inCtx").Dot("Err").Call(),
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Return(),
				),
				jen.List(jen.Id("y0"), jen.Err()).Op(":=").Id("fn").Call(jen.Id("ctx")),
				jen.If(
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Id("cancel").Call(jen.Err()),
					jen.Return(jen.Op("*").New(jen.Id("Y0"))),
				),
				jen.Return(jen.Id("y0")),
			),
		),
	)
}

const maxIn = 10
const maxOut = 10

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

func params(f *jen.Statement, num int) {
	ps := make([]jen.Code, num+1)

	ps[0] = jen.Id("ctx").Qual("context", "Context")

	for i := 0; i < num; i++ {
		ps[i+1] = jen.Id(fmt.Sprintf("x%d", i)).Id(fmt.Sprintf("X%d", i))
	}

	f.Params(ps...)
}

func returns(f *jen.Statement, num int) {
	ps := make([]jen.Code, num)

	for i := 0; i < num; i++ {
		ps[i] = jen.Id(fmt.Sprintf("y%d", i)).Id(fmt.Sprintf("Y%d", i))
	}

	f.Params(ps...)
}
