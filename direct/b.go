package direct

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
) B {
	return B{
		UniversalField:  UniversalField,
		BSpecificField1: BSpecificField1,
		BSpecificField2: BSpecificField2,
	}
}

func B_FirstOp(b B, c int, d int) int {
	return 2*c*b.BSpecificField2 + d*b.BSpecificField1
}

func B_SecondOp(b B, d string) string {
	var builder strings.Builder
	builder.WriteString("from b(")
	builder.WriteString(b.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func B_ThirdOp(b B) string {
	var builder strings.Builder
	builder.WriteString("b specific: ")
	builder.WriteString(strconv.FormatInt(int64(b.BSpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(strconv.FormatInt(int64(b.BSpecificField2), 10))
	return builder.String()
}
