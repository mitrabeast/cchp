package interfaced

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
) Interface {
	return &B{
		UniversalField:  UniversalField,
		BSpecificField1: BSpecificField1,
		BSpecificField2: BSpecificField2,
	}
}

func (b *B) FirstOp(c int, d int) int {
	return 2*c*b.BSpecificField2 + d*b.BSpecificField1
}

func (b *B) SecondOp(d string) string {
	return fmt.Sprintf("from b (%s): %s", b.UniversalField, d)
}

func (b *B) ThirdOp() string {
	return fmt.Sprintf("b specific: %d, %d", b.BSpecificField1, b.BSpecificField2)
}
