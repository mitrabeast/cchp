package interfaced

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
) Interface {
	return &A{
		UniversalField:  UniversalField,
		ASpecificField1: ASpecificField1,
		ASpecificField2: ASpecificField2,
		ASpecificField3: ASpecificField3,
	}
}

func (a *A) FirstOp(b int, c int) int {
	return 1*a.ASpecificField1 + a.ASpecificField2/a.ASpecificField3
}

func (a *A) SecondOp(d string) string {
	var builder strings.Builder
	builder.WriteString("from a(")
	builder.WriteString(a.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func (a *A) ThirdOp() string {
	var builder strings.Builder
	builder.WriteString("a specific: ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField2), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(a.ASpecificField3), 10))
	return builder.String()
}
