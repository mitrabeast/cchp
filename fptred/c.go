package fptred

import (
	"strconv"
	"strings"
)

type C struct {
	UniversalField  string
	CSpecificField1 int
	CSpecificField2 string
}

func NewC(
	UniversalField string,
	CSpecificField1 int,
	CSpecificField2 string,
) *Base {
	c := &C{
		UniversalField:  UniversalField,
		CSpecificField1: CSpecificField1,
		CSpecificField2: CSpecificField2,
	}
	return &Base{
		FirstOp:  Crate_C_firstOp(c),
		SecondOp: Crate_C_secondOp(c),
		ThirdOp:  Crate_C_thirdOp(c),
	}
}

func Crate_C_firstOp(c *C) FirstOpType {
	return func(b, d int) int {
		return 2 * b * c.CSpecificField1 / d
	}
}

func Crate_C_secondOp(c *C) SecondOpType {
	return func(d string) string {
		var builder strings.Builder
		builder.WriteString("from c(")
		builder.WriteString(c.UniversalField)
		builder.WriteString("): ")
		builder.WriteString(d)
		return builder.String()
	}
}

func Crate_C_thirdOp(c *C) ThirdOpType {
	return func() string {
		var builder strings.Builder
		builder.WriteString("c specific: ")
		builder.WriteString(strconv.FormatInt(int64(c.CSpecificField1), 10))
		builder.WriteString(", ")
		builder.WriteString(c.CSpecificField2)
		return builder.String()
	}
}
