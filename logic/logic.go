package logic

import (
	"cchp/direct"
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/tabled"
	"cchp/util"
)

var (
	a = util.RandInt()
	b = util.RandInt()
	s = util.RandString(10)
)

func DoInterfaced(instance interfaced.Interface) string {
	first := instance.FirstOp(a, b)
	// second := instance.SecondOp(s)
	// third := instance.ThirdOp()

	_ /*, _, _*/ = first /*, second, third*/
	return ""
	// return fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/
}

func DoTabled(instance *tabled.Base) string {
	first := instance.FirstOp(a, b)
	// second := instance.SecondOp(s)
	// third := instance.ThirdOp()

	_ /*, _, _*/ = first /*, second, third*/
	return ""
	// return fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/
}

func DoDirectTabled(instance *tabled.Base) string {
	first := tabled.FirstOpsTable[instance.Type](instance, a, b)
	// second := tabled.SecondOpsTable[instance.Type](instance, s)
	// third := tabled.ThirdOpsTable[instance.Type](instance)

	_ /*, _, _*/ = first /*, second, third*/
	return ""
	// return fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/
}

func DoFunctionPointed(instance *fptred.Base) string {
	first := instance.FirstOp(a, b)
	// second := instance.SecondOp(s)
	// third := instance.ThirdOp()

	_ /*, _, _*/ = first /*, second, third*/
	return ""
	// return fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/
}

func DoDirectA(instance direct.A) string {
	first := direct.A_FirstOp(instance, a, b)
	// second := direct.A_SecondOp(instance, s)
	// third := direct.A_ThirdOp(instance)

	_ /*, _, _*/ = first /*, second, third*/
	// aResult := fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/

	return ""
}

func DoDirectB(instance direct.B) string {
	first := direct.B_FirstOp(instance, a, b)
	// second := direct.B_SecondOp(instance, s)
	// third := direct.B_ThirdOp(instance)

	_ /*, _, _*/ = first /*, second, third*/
	// bResult := fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/

	return ""
}

func DoDirectC(instance direct.C) string {
	first := direct.C_FirstOp(instance, a, b)
	// second := direct.C_SecondOp(instance, s)
	// third := direct.C_ThirdOp(instance)

	_ /*, _, _*/ = first /*, second, third*/
	// cResult := fmt.Sprintf("%d, %s/*, %s", first*//*, second, third)*/

	return ""
}
