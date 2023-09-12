package logic

import (
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/tabled"
	"cchp/util"
	"fmt"
)

var (
	a = util.RandInt()
	b = util.RandInt()
	s = util.RandString(10)
)

func DoInterfaced(instance interfaced.Interface) string {
	first := instance.FirstOp(a, b)
	second := instance.SecondOp(s)
	third := instance.ThirdOp()

	// _, _, _ = first, second, third
	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoTabled(instance *tabled.Base) string {
	first := instance.FirstOp(a, b)
	second := instance.SecondOp(s)
	third := instance.ThirdOp()

	// _, _, _ = first, second, third
	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoDirectTabled(instance *tabled.Base) string {
	first := tabled.FirstOpsTable[instance.Type](instance, a, b)
	second := tabled.SecondOpsTable[instance.Type](instance, s)
	third := tabled.ThirdOpsTable[instance.Type](instance)

	// _, _, _ = first, second, third
	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoFunctionPointed(instance *fptred.Base) string {
	first := instance.FirstOp(a, b)
	second := instance.SecondOp(s)
	third := instance.ThirdOp()

	// _, _, _ = first, second, third
	return fmt.Sprintf("%d, %s, %s", first, second, third)
}
