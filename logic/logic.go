package logic

import (
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/tabled"
)

func DoInterfaced(instance interfaced.Interface) {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	_, _, _ = first, second, third
	// return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoTabled(instance *tabled.Base) {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	_, _, _ = first, second, third
	// return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoDirectTabled(instance *tabled.Base) {
	first := tabled.FirstOpsTable[instance.Type](instance, 2, 3)
	second := tabled.SecondOpsTable[instance.Type](instance, "hello")
	third := tabled.ThirdOpsTable[instance.Type](instance)

	_, _, _ = first, second, third
	// return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoFunctionPointed(instance *fptred.Base) {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	_, _, _ = first, second, third
}
