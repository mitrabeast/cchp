package main

import (
	"fmt"
	"log"

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

func main() {
	runtime, ticks := util.MeasureRuntime(func() {
		DoInterface(interfaced.NewA("A", 1, 2, 3))
		DoInterface(interfaced.NewB("B", 10, 20))
		DoInterface(interfaced.NewC("C", 100, "200"))
	})
	log.Printf("%10s, took: %8v, %8v cycles", "interfaced", runtime, ticks)

	runtime, ticks = util.MeasureRuntime(func() {
		DoTable(tabled.NewA("A", 1, 2, 3))
		DoTable(tabled.NewB("B", 10, 20))
		DoTable(tabled.NewC("C", 100, "200"))
	})
	log.Printf("%10s, took: %8v, %8v cycles", "tabled", runtime, ticks)
}
