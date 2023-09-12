package fptred

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
	c := &C{
		UniversalField:  UniversalField,
		CSpecificField1: CSpecificField1,
		CSpecificField2: CSpecificField2,
	}
	return &Base{
		FirstOp:  Crate_C_firstOp(c),
		SecondOp: Crate_C_secondOp(c),
		ThirdOp:  Crate_C_thirdOp(c),
	}
}

func Crate_C_firstOp(c *C) FirstOpType {
	return func(b, d int) int {
		return 2 * b * c.CSpecificField1 / d
	}
}

func Crate_C_secondOp(c *C) SecondOpType {
	return func(d string) string {
		return fmt.Sprintf("from c (%s): %s", c.UniversalField, d)
	}
}

func Crate_C_thirdOp(c *C) ThirdOpType {
	return func() string {
		return fmt.Sprintf("b specific: %d, %s", c.CSpecificField1, c.CSpecificField2)
	}
}
