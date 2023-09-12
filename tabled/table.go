package tabled

type Base struct {
	Type ItemType
	*A
	*B
	*C
}

type ItemType int

const (
	AType ItemType = iota
	BType
	CType
)

type FirstOpType func(a *Base, b int, c int) int
type SecondOpType func(a *Base, d string) string
type ThirdOpType func(a *Base) string

var FirstOpsTable = []FirstOpType{
	AType: A_firstOp,
	BType: B_firstOp,
	CType: C_firstOp,
}

var SecondOpsTable = []SecondOpType{
	AType: A_secondOp,
	BType: B_secondOp,
	CType: C_secondOp,
}

var ThirdOpsTable = []ThirdOpType{
	AType: A_thirdOp,
	BType: B_thirdOp,
	CType: C_thirdOp,
}

func (item *Base) FirstOp(b int, c int) int {
	return FirstOpsTable[item.Type](item, b, c)
}

func (item *Base) SecondOp(d string) string {
	return SecondOpsTable[item.Type](item, d)
}

func (item *Base) ThirdOp() string {
	return ThirdOpsTable[item.Type](item)
}
