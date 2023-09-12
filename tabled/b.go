package tabled

import "fmt"

type B struct {
	UniversalField  string
	BSpecificField1 int
	BSpecificField2 int
}

func NewB(
	UniversalField string,
	BSpecificField1 int,
	BSpecificField2 int,
) *Base {
	return &Base{
		B: B{
			UniversalField:  UniversalField,
			BSpecificField1: BSpecificField1,
			BSpecificField2: BSpecificField2,
		},
		Type: BType,
	}
}

func B_firstOp(b_ *Base, c int, d int) int {
	b := b_.B
	return 2*c*b.BSpecificField2 + d*b.BSpecificField1
}

func B_secondOp(b_ *Base, d string) string {
	b := b_.B
	return fmt.Sprintf("from b (%s): %s", b.UniversalField, d)
}

func B_thirdOp(b_ *Base) string {
	b := b_.B
	return fmt.Sprintf("b specific: %d, %d", b.BSpecificField1, b.BSpecificField2)
}
