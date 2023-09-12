package logic

import (
	"cchp/direct"
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

func DoDirect(aInstance direct.A, bInstance direct.B, cInstance direct.C) string {
	first := direct.A_FirstOp(aInstance, a, b)
	second := direct.A_SecondOp(aInstance, s)
	third := direct.A_ThirdOp(aInstance)

	_, _, _ = first, second, third
	// aResult := fmt.Sprintf("%d, %s, %s", first, second, third)

	first = direct.B_FirstOp(bInstance, a, b)
	second = direct.B_SecondOp(bInstance, s)
	third = direct.B_ThirdOp(bInstance)

	_, _, _ = first, second, third
	// bResult := fmt.Sprintf("%d, %s, %s", first, second, third)

	first = direct.C_FirstOp(cInstance, a, b)
	second = direct.C_SecondOp(cInstance, s)
	third = direct.C_ThirdOp(cInstance)

	_, _, _ = first, second, third
	// cResult := fmt.Sprintf("%d, %s, %s", first, second, third)

	// return fmt.Sprintf("\t%s\n\t%s\n\t%s\n", aResult, bResult, cResult)

	return ""
}
