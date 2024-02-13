package cfg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/a20r/mesa"
	"github.com/a20r/useful/cfg"
)

func TestCfg_Validate(t *testing.T) {
	table := mesa.FunctionMesa[any, error]{
		Target: func(ctx *mesa.Ctx, in any) error {
			return cfg.Validate(in)
		},

		Cases: []mesa.FunctionCase[any, error]{
			{
				Name: "All required fields provided",

				Input: struct {
					Foo string `cfg:"required"`
					Bar string `cfg:"required"`
				}{
					Foo: "foo",
					Bar: "bar",
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.NoError(out)
				},
			},
			{
				Name: "Missing single required field",

				Input: struct {
					Foo string `cfg:"required"`
					Bar string `cfg:"required"`
				}{
					Foo: "foo",
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.EqualError(out, "cfg: missing required fields [Bar]")
				},
			},
			{
				Name: "Missing multiple required fields",

				Input: struct {
					Foo string `cfg:"required"`
					Bar string `cfg:"required"`
				}{},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.EqualError(out, "cfg: missing required fields [Foo Bar]")
				},
			},
			{
				Name: "Missing single optional field",

				Input: struct {
					Foo string `cfg:"optional"`
					Bar string `cfg:"required"`
				}{
					Bar: "bar",
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.NoError(out)
				},
			},
			{
				Name: "Missing multiple optional fields",

				Input: struct {
					Foo string `cfg:"optional"`
					Bar string `cfg:"optional"`
				}{},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.NoError(out)
				},
			},
			{
				Name: "Missing required pointer field",

				Input: struct {
					Foo *string `cfg:"required"`
					Bar string  `cfg:"optional"`
				}{
					Bar: "bar",
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.EqualError(out, "cfg: missing required fields [Foo]")
				},
			},
			{
				Name: "Required pointer field provided",

				Input: struct {
					Foo *string `cfg:"required"`
					Bar string  `cfg:"optional"`
				}{
					Foo: new(string),
					Bar: "bar",
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.NoError(out)
				},
			},
			{
				Name: "Self validation succeeds",

				Input: selfValidationStruct{
					Secs: 420,
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					ctx.As.NoError(out)
				},
			},
			{
				Name: "Self validation fails",

				Input: selfValidationStruct{
					Secs: -1,
				},

				Check: func(ctx *mesa.Ctx, in any, out error) {
					expected := fmt.Errorf("Time field cannot be zero")
					ctx.As.Error(out)
					ctx.As.ErrorIs(out, cfg.ErrCfgSelfValidationFailed.New().Wrap(expected))
				},
			},
		},
	}

	table.Run(t)
}

type selfValidationStruct struct {
	Secs time.Duration `cfg:"required"`
}

func (s selfValidationStruct) Validate() error {
	if s.Secs < 0 {
		return fmt.Errorf("Time field cannot be zero")
	}

	return nil
}
