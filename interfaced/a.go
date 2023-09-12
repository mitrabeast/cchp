package interfaced

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
	return fmt.Sprintf("from a (%s): %s", a.UniversalField, d)
}

func (a *A) ThirdOp() string {
	return fmt.Sprintf("a specific: %d, %d, %d", a.ASpecificField1, a.ASpecificField2, a.ASpecificField3)
}
