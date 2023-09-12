package direct

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
	return fmt.Sprintf("from a (%s): %s", a.UniversalField, d)
}

func A_ThirdOp(a A) string {
	return fmt.Sprintf("a specific: %d, %d, %d", a.ASpecificField1, a.ASpecificField2, a.ASpecificField3)
}
