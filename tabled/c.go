package tabled

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
	return &Base{
		C: &C{
			UniversalField:  UniversalField,
			CSpecificField1: CSpecificField1,
			CSpecificField2: CSpecificField2,
		},
		Type: CType,
	}
}

func C_firstOp(c_ *Base, b int, d int) int {
	c := c_.C
	return 2 * b * c.CSpecificField1 / d
}

func C_secondOp(c_ *Base, d string) string {
	c := c_.C
	var builder strings.Builder
	builder.WriteString("from c(")
	builder.WriteString(c.UniversalField)
	builder.WriteString("): ")
	builder.WriteString(d)
	return builder.String()
}

func C_thirdOp(c_ *Base) string {
	c := c_.C
	var builder strings.Builder
	builder.WriteString("c specific: ")
	builder.WriteString(strconv.FormatInt(int64(c.CSpecificField1), 10))
	builder.WriteString(", ")
	builder.WriteString(c.CSpecificField2)
	return builder.String()
}
