package main

import (
	"fmt"
	"time"

	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"cchp/util"
)

func main() {
	var runtime time.Duration
	var ticks uint64
	var a, b, c string

	// make it work faster somehow ???
	aItem := interfaced.NewC("C", util.RandInt(), util.RandString(10))
	bItem := tabled.NewC("C", util.RandInt(), util.RandString(10))
	cItem := fptred.NewC("C", util.RandInt(), util.RandString(10))
	for i := int64(0); i < 999_999; i++ {
		logic.DoInterfaced(aItem)
		logic.DoTabled(bItem)
		logic.DoDirectTabled(bItem)
		logic.DoFunctionPointed(cItem)
	}

	fptredA := fptred.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	fptredB := fptred.NewB("B", util.RandInt(), util.RandInt())
	fptredC := fptred.NewC("C", util.RandInt(), util.RandString(10))

	runtime, ticks = util.MeasureRuntime(func() {
		a = logic.DoFunctionPointed(fptredA)
		b = logic.DoFunctionPointed(fptredB)
		c = logic.DoFunctionPointed(fptredC)
	})
	fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%s, took: %8v, %8v cycles\n", "function pointed", runtime, ticks)

	tabledA := tabled.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	tabledB := tabled.NewB("B", util.RandInt(), util.RandInt())
	tabledC := tabled.NewC("C", util.RandInt(), util.RandString(10))

	runtime, ticks = util.MeasureRuntime(func() {
		a = logic.DoTabled(tabledA)
		b = logic.DoTabled(tabledB)
		c = logic.DoTabled(tabledC)
	})
	fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%s, took: %8v, %8v cycles\n", "tabled", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		a = logic.DoDirectTabled(tabledA)
		b = logic.DoDirectTabled(tabledB)
		c = logic.DoDirectTabled(tabledC)
	})
	fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%s, took: %8v, %8v cycles\n", "direct tabled", runtime, ticks)

	interfacedA := interfaced.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	interfacedB := interfaced.NewB("B", util.RandInt(), util.RandInt())
	interfacedC := interfaced.NewC("C", util.RandInt(), util.RandString(10))

	runtime, ticks = util.MeasureRuntime(func() {
		a = logic.DoInterfaced(interfacedA)
		b = logic.DoInterfaced(interfacedB)
		c = logic.DoInterfaced(interfacedC)
	})
	fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%s, took: %8v, %8v cycles\n", "interfaced", runtime, ticks)
}
