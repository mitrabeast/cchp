package tabled

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
	return fmt.Sprintf("from a (%s): %s", a.UniversalField, d)
}

func A_thirdOp(a_ *Base) string {
	a := a_.A
	return fmt.Sprintf("a specific: %d, %d, %d", a.ASpecificField1, a.ASpecificField2, a.ASpecificField3)
}
