package direct

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
) C {
	return C{
		UniversalField:  UniversalField,
		CSpecificField1: CSpecificField1,
		CSpecificField2: CSpecificField2,
	}
}

func C_FirstOp(c C, b int, d int) int {
	return 2 * b * c.CSpecificField1 / d
}

func C_SecondOp(c C, d string) string {
	var builder strings.Builder
	builder.WriteString("from c(")
	builder.WriteString(c.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func C_ThirdOp(c C) string {
	var builder strings.Builder
	builder.WriteString("c specific: ")
	builder.WriteString(strconv.FormatInt(int64(c.CSpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(c.CSpecificField2)
	return builder.String()
}
