package main

import (
	"fmt"

	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"cchp/util"
)

func main() {
	interfacedA := interfaced.NewA("A", 1, 2, 3)
	interfacedB := interfaced.NewB("B", 10, 20)
	interfacedC := interfaced.NewC("C", 100, "200")
	runtime, ticks := util.MeasureRuntime(func() {
		logic.DoInterfaced(interfacedA)
		logic.DoInterfaced(interfacedB)
		logic.DoInterfaced(interfacedC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "interfaced", runtime, ticks)

	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	runtime, ticks = util.MeasureRuntime(func() {
		logic.DoTabled(tabledA)
		logic.DoTabled(tabledB)
		logic.DoTabled(tabledC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "tabled", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		logic.DoDirectTabled(tabledA)
		logic.DoDirectTabled(tabledB)
		logic.DoDirectTabled(tabledC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "direct tabled", runtime, ticks)

	fptredA := fptred.NewA("A", 1, 2, 3)
	fptredB := fptred.NewB("B", 10, 20)
	fptredC := fptred.NewC("C", 100, "200")
	runtime, ticks = util.MeasureRuntime(func() {
		logic.DoFunctionPointed(fptredA)
		logic.DoFunctionPointed(fptredB)
		logic.DoFunctionPointed(fptredC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "function pointed", runtime, ticks)

}
