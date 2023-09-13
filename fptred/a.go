package fptred

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
	a := &A{
		UniversalField:  UniversalField,
		ASpecificField1: ASpecificField1,
		ASpecificField2: ASpecificField2,
		ASpecificField3: ASpecificField3,
	}
	return &Base{
		FirstOp:  Create_A_firstOp(a),
		SecondOp: Create_A_secondOp(a),
		ThirdOp:  Create_A_thirdOp(a),
	}
}

func Create_A_firstOp(a *A) FirstOpType {
	return func(b, c int) int {
		return 1*a.ASpecificField1 + a.ASpecificField2/a.ASpecificField3
	}
}

func Create_A_secondOp(a *A) SecondOpType {
	return func(d string) string {
		var builder strings.Builder
		builder.WriteString("from a(")
		builder.WriteString(a.UniversalField)
		builder.WriteString("): ")
		builder.WriteString(d)
		return builder.String()
	}
}

func Create_A_thirdOp(a *A) ThirdOpType {
	return func() string {
		var builder strings.Builder
		builder.WriteString("a specific: ")
		builder.WriteString(strconv.FormatInt(int64(a.ASpecificField1), 10))
		builder.WriteString(", ")
		builder.WriteString(strconv.FormatInt(int64(a.ASpecificField2), 10))
		builder.WriteString(", ")
		builder.WriteString(strconv.FormatInt(int64(a.ASpecificField3), 10))
		return builder.String()
	}
}
