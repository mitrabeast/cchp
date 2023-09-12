package direct

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
) B {
	return B{
		UniversalField:  UniversalField,
		BSpecificField1: BSpecificField1,
		BSpecificField2: BSpecificField2,
	}
}

func B_FirstOp(b B, c int, d int) int {
	return 2*c*b.BSpecificField2 + d*b.BSpecificField1
}

func B_SecondOp(b B, d string) string {
	return fmt.Sprintf("from b (%s): %s", b.UniversalField, d)
}

func B_ThirdOp(b B) string {
	return fmt.Sprintf("b specific: %d, %d", b.BSpecificField1, b.BSpecificField2)
}
