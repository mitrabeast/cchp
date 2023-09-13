package tabled

import (
	"strconv"
	"strings"
)

type A struct {
	UniversalField  string
	ASpecificField1 int
	ASpecificField2 int
	ASpecificField3 int
}

func NewA(
	UniversalField string,
	ASpecificField1 int,
	ASpecificField2 int,
	ASpecificField3 int,
) *Base {
	return &Base{
		A: &A{
			UniversalField:  UniversalField,
			ASpecificField1: ASpecificField1,
			ASpecificField2: ASpecificField2,
			ASpecificField3: ASpecificField3,
		},
		Type: AType,
	}
}

func A_firstOp(a_ *Base, b int, c int) int {
	a := a_.A
	return 1*a.ASpecificField1 + a.ASpecificField2/a.ASpecificField3
}

func A_secondOp(a_ *Base, d string) string {
	a := a_.A
	var builder strings.Builder
	builder.WriteString("from a(")
	builder.WriteString(a.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func A_thirdOp(a_ *Base) string {
	a := a_.A
	var builder strings.Builder
	builder.WriteString("a specific: ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField2), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField3), 10))
	return builder.String()
}
