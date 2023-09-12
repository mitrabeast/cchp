package tabled

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
) *Base {
	return &Base{
		C: C{
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
	return fmt.Sprintf("from c (%s): %s", c.UniversalField, d)
}

func C_thirdOp(c_ *Base) string {
	c := c_.C
	return fmt.Sprintf("b specific: %d, %s", c.CSpecificField1, c.CSpecificField2)
}
