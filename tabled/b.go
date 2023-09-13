package tabled

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
	return &Base{
		B: &B{
			UniversalField:  UniversalField,
			BSpecificField1: BSpecificField1,
			BSpecificField2: BSpecificField2,
		},
		Type: BType,
	}
}

func B_firstOp(b_ *Base, c int, d int) int {
	b := b_.B
	return 2*c*b.BSpecificField2 + d*b.BSpecificField1
}

func B_secondOp(b_ *Base, d string) string {
	b := b_.B
	var builder strings.Builder
	builder.WriteString("from b(")
	builder.WriteString(b.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func B_thirdOp(b_ *Base) string {
	b := b_.B
	var builder strings.Builder
	builder.WriteString("b specific: ")
	builder.WriteString(strconv.FormatInt(int64(b.BSpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(b.BSpecificField2), 10))
	return builder.String()
}
