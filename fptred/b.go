package fptred

import (
	"strconv"
	"strings"
)

type B struct {
	UniversalField  string
	BSpecificField1 int
	BSpecificField2 int
}

func NewB(
	UniversalField string,
	BSpecificField1 int,
	BSpecificField2 int,
) *Base {
	b := &B{
		UniversalField:  UniversalField,
		BSpecificField1: BSpecificField1,
		BSpecificField2: BSpecificField2,
	}
	return &Base{
		FirstOp:  Create_B_firstOp(b),
		SecondOp: Create_B_secondOp(b),
		ThirdOp:  Create_B_thirdOp(b),
	}
}

func Create_B_firstOp(b *B) FirstOpType {
	return func(c, d int) int {
		return 2*c*b.BSpecificField2 + d*b.BSpecificField1
	}
}

func Create_B_secondOp(b *B) SecondOpType {
	return func(d string) string {
		var builder strings.Builder
		builder.WriteString("from b(")
		builder.WriteString(b.UniversalField)
		builder.WriteString("): ")
		builder.WriteString(d)
		return builder.String()
	}
}

func Create_B_thirdOp(b *B) ThirdOpType {
	return func() string {
		var builder strings.Builder
		builder.WriteString("b specific: ")
		builder.WriteString(strconv.FormatInt(int64(b.BSpecificField1), 10))
		builder.WriteString(", ")
		builder.WriteString(strconv.FormatInt(int64(b.BSpecificField2), 10))
		return builder.String()
	}
}
