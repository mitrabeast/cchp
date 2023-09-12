package main

import (
	"fmt"

	"cchp/interfaced"
	"cchp/tabled"
	"cchp/util"
)

func DoInterface(instance interfaced.Interface) string {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoTable(instance *tabled.Base) string {
	first := instance.FirstOp(2, 3)
	second := instance.SecondOp("hello")
	third := instance.ThirdOp()

	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func DoDirectTable(instance *tabled.Base) string {
	first := tabled.FirstOpsTable[instance.Type](instance, 2, 3)
	second := tabled.SecondOpsTable[instance.Type](instance, "hello")
	third := tabled.ThirdOpsTable[instance.Type](instance)

	return fmt.Sprintf("%d, %s, %s", first, second, third)
}

func main() {
	runtime, ticks := util.MeasureRuntime(func() {
		DoInterface(interfaced.NewA("A", 1, 2, 3))
		DoInterface(interfaced.NewB("B", 10, 20))
		DoInterface(interfaced.NewC("C", 100, "200"))
	})
	fmt.Printf("%15s, took: %8v, %8v cycles\n", "interfaced", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		DoTable(tabled.NewA("A", 1, 2, 3))
		DoTable(tabled.NewB("B", 10, 20))
		DoTable(tabled.NewC("C", 100, "200"))
	})
	fmt.Printf("%15s, took: %8v, %8v cycles\n", "tabled", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		DoDirectTable(tabled.NewA("A", 1, 2, 3))
		DoDirectTable(tabled.NewB("B", 10, 20))
		DoDirectTable(tabled.NewC("C", 100, "200"))
	})
	fmt.Printf("%15s, took: %8v, %8v cycles\n", "direct tabled", runtime, ticks)
}
