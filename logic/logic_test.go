package logic_test

import (
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"testing"
)

func BenchmarkDoInterface(b *testing.B) {
	interfacedA := interfaced.NewA("A", 1, 2, 3)
	interfacedB := interfaced.NewB("B", 10, 20)
	interfacedC := interfaced.NewC("C", 100, "200")

	b.ResetTimer()
	b.Run("interfaced", func(b *testing.B) {
		logic.DoInterfaced(interfacedA)
		logic.DoInterfaced(interfacedB)
		logic.DoInterfaced(interfacedC)
	})
}

func BenchmarkDoTabled(b *testing.B) {
	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	b.ResetTimer()
	b.Run("tabled", func(b *testing.B) {
		logic.DoTabled(tabledA)
		logic.DoTabled(tabledB)
		logic.DoTabled(tabledC)
	})
}

func BenchmarkDoDirectTabled(b *testing.B) {
	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	b.ResetTimer()
	b.Run("direct tabled", func(b *testing.B) {
		logic.DoDirectTabled(tabledA)
		logic.DoDirectTabled(tabledB)
		logic.DoDirectTabled(tabledC)
	})
}

func BenchmarkDoFunctionPointed(b *testing.B) {
	fptredA := fptred.NewA("A", 1, 2, 3)
	fptredB := fptred.NewB("B", 10, 20)
	fptredC := fptred.NewC("C", 100, "200")

	b.ResetTimer()
	b.Run("func pointed", func(b *testing.B) {
		logic.DoFunctionPointed(fptredA)
		logic.DoFunctionPointed(fptredB)
		logic.DoFunctionPointed(fptredC)
	})
}
