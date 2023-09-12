package fptred

type Base struct {
	FirstOp  FirstOpType
	SecondOp SecondOpType
	ThirdOp  ThirdOpType
}

type FirstOpType func(b int, c int) int
type SecondOpType func(d string) string
type ThirdOpType func() string
