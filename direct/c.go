package direct

import "fmt"

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
	return fmt.Sprintf("from c (%s): %s", c.UniversalField, d)
}

func C_ThirdOp(c C) string {
	return fmt.Sprintf("b specific: %d, %s", c.CSpecificField1, c.CSpecificField2)
}
