package direct

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
) A {
	return A{
		UniversalField:  UniversalField,
		ASpecificField1: ASpecificField1,
		ASpecificField2: ASpecificField2,
		ASpecificField3: ASpecificField3,
	}
}

func A_FirstOp(a A, b int, c int) int {
	return 1*a.ASpecificField1 + a.ASpecificField2/a.ASpecificField3
}

func A_SecondOp(a A, d string) string {
	var builder strings.Builder
	builder.WriteString("from a(")
	builder.WriteString(a.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func A_ThirdOp(a A) string {
	var builder strings.Builder
	builder.WriteString("a specific: ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField2), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField3), 10))
	return builder.String()
}
