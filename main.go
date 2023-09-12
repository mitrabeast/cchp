package main

import (
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"cchp/util"
	"fmt"
	"time"
)

func main() {
	var runtime time.Duration
	var ticks uint64

	// make it work faster somehow ???
	a := interfaced.NewC("C", 100, "200")
	b := tabled.NewC("C", 100, "200")
	c := fptred.NewC("C", 100, "200")
	for i := int64(0); i < 999_999; i++ {
		logic.DoInterfaced(a)
		logic.DoTabled(b)
		logic.DoDirectTabled(b)
		logic.DoFunctionPointed(c)
	}

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

	interfacedA := interfaced.NewA("A", 1, 2, 3)
	interfacedB := interfaced.NewB("B", 10, 20)
	interfacedC := interfaced.NewC("C", 100, "200")
	runtime, ticks = util.MeasureRuntime(func() {
		logic.DoInterfaced(interfacedA)
		logic.DoInterfaced(interfacedB)
		logic.DoInterfaced(interfacedC)
	})
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "interfaced", runtime, ticks)
}
