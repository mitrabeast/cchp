package main

import (
	"fmt"

	"cchp/fptred"
	"cchp/interfaced"
	"cchp/tabled"
	"cchp/util"
)

func DoInterface(instance interfaced.Interface) {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	_, _, _ = first, second, third
	// return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoTable(instance *tabled.Base) {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	_, _, _ = first, second, third
	// return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoDirectTable(instance *tabled.Base) {
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

func main() {
	interfacedA := interfaced.NewA("A", 1, 2, 3)
	interfacedB := interfaced.NewB("B", 10, 20)
	interfacedC := interfaced.NewC("C", 100, "200")
	runtime, ticks := util.MeasureRuntime(func() {
		DoInterface(interfacedA)
		DoInterface(interfacedB)
		DoInterface(interfacedC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "interfaced", runtime, ticks)

	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	runtime, ticks = util.MeasureRuntime(func() {
		DoTable(tabledA)
		DoTable(tabledB)
		DoTable(tabledC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "tabled", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		DoDirectTable(tabledA)
		DoDirectTable(tabledB)
		DoDirectTable(tabledC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "direct tabled", runtime, ticks)

	fptredA := fptred.NewA("A", 1, 2, 3)
	fptredB := fptred.NewB("B", 10, 20)
	fptredC := fptred.NewC("C", 100, "200")
	runtime, ticks = util.MeasureRuntime(func() {
		DoFunctionPointed(fptredA)
		DoFunctionPointed(fptredB)
		DoFunctionPointed(fptredC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "function pointed", runtime, ticks)

}
