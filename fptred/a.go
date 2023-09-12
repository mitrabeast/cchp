package fptred

import "fmt"

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
		return fmt.Sprintf("from a (%s): %s", a.UniversalField, d)
	}
}

func Create_A_thirdOp(a *A) ThirdOpType {
	return func() string {
		return fmt.Sprintf("a specific: %d, %d, %d", a.ASpecificField1, a.ASpecificField2, a.ASpecificField3)
	}
}
