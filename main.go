package main

import (
	"fmt"
	"time"

	"cchp/direct"
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"cchp/util"
)

var (
	interfacedA = interfaced.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	interfacedB = interfaced.NewB("B", util.RandInt(), util.RandInt())
	interfacedC = interfaced.NewC("C", util.RandInt(), util.RandString(10))

	tabledA = tabled.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	tabledB = tabled.NewB("B", util.RandInt(), util.RandInt())
	tabledC = tabled.NewC("C", util.RandInt(), util.RandString(10))

	fptredA = fptred.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	fptredB = fptred.NewB("B", util.RandInt(), util.RandInt())
	fptredC = fptred.NewC("C", util.RandInt(), util.RandString(10))

	directA = direct.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	directB = direct.NewB("B", util.RandInt(), util.RandInt())
	directC = direct.NewC("C", util.RandInt(), util.RandString(10))
)

const tests = 999_999

func main() {
	var ranInSum time.Duration
	var ticksSum uint64
	var a, b, c string

	ranInSum, ticksSum = 0, 0
	for i := 0; i < tests; i++ {
		ranIn, ticks := util.MeasureRuntime(func() {
			a = logic.DoDirectTabled(tabledA)
			b = logic.DoDirectTabled(tabledB)
			c = logic.DoDirectTabled(tabledC)
		})
		if i > 5 {
			ranInSum += ranIn
			ticksSum += ticks
		}
	}
	// fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "direct tabled", ranInSum/(tests-6), ticksSum/(tests-6))

	ranInSum, ticksSum = 0, 0
	for i := 0; i < tests; i++ {
		ranIn, ticks := util.MeasureRuntime(func() {
			a = logic.DoFunctionPointed(fptredA)
			b = logic.DoFunctionPointed(fptredB)
			c = logic.DoFunctionPointed(fptredC)
		})
		if i > 5 {
			ranInSum += ranIn
			ticksSum += ticks
		}
	}
	// fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "function pointed", ranInSum/(tests-6), ticksSum/(tests-6))

	ranInSum, ticksSum = 0, 0
	for i := 0; i < tests; i++ {
		ranIn, ticks := util.MeasureRuntime(func() {
			a = logic.DoTabled(tabledA)
			b = logic.DoTabled(tabledB)
			c = logic.DoTabled(tabledC)
		})
		if i > 5 {
			ranInSum += ranIn
			ticksSum += ticks
		}
	}
	// fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "tabled", ranInSum/(tests-6), ticksSum/(tests-6))

	ranInSum, ticksSum = 0, 0
	for i := 0; i < tests; i++ {
		ranIn, ticks := util.MeasureRuntime(func() {
			a = logic.DoInterfaced(interfacedA)
			b = logic.DoInterfaced(interfacedB)
			c = logic.DoInterfaced(interfacedC)
		})
		if i > 5 {
			ranInSum += ranIn
			ticksSum += ticks
		}
	}
	// fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "interfaced", ranInSum/(tests-6), ticksSum/(tests-6))

	ranInSum, ticksSum = 0, 0
	for i := 0; i < tests; i++ {
		ranIn, ticks := util.MeasureRuntime(func() {
			a = logic.DoDirect(directA, directB, directC)
		})
		if i > 5 {
			ranInSum += ranIn
			ticksSum += ticks
		}
	}
	// fmt.Printf("\tresult:\n\t%s\n\t%s\n\t%s\n", a, b, c)
	fmt.Printf("%16s, took: %8v, %8v cycles\n", "direct", ranInSum/(tests-6), ticksSum/(tests-6))

	_, _, _ = a, b, c
}
