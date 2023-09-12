package logic_test

import (
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"testing"
)

func BenchmarkDoInterfaced(b *testing.B) {
	interfacedA := interfaced.NewA("A", 1, 2, 3)
	interfacedB := interfaced.NewB("B", 10, 20)
	interfacedC := interfaced.NewC("C", 100, "200")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logic.DoInterfaced(interfacedA)
		logic.DoInterfaced(interfacedB)
		logic.DoInterfaced(interfacedC)
	}
}

func BenchmarkDoTabled(b *testing.B) {
	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logic.DoTabled(tabledA)
		logic.DoTabled(tabledB)
		logic.DoTabled(tabledC)
	}
}

func BenchmarkDoDirectTabled(b *testing.B) {
	tabledA := tabled.NewA("A", 1, 2, 3)
	tabledB := tabled.NewB("B", 10, 20)
	tabledC := tabled.NewC("C", 100, "200")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logic.DoDirectTabled(tabledA)
		logic.DoDirectTabled(tabledB)
		logic.DoDirectTabled(tabledC)
	}
}

func BenchmarkDoFunctionPointed(b *testing.B) {
	fptredA := fptred.NewA("A", 1, 2, 3)
	fptredB := fptred.NewB("B", 10, 20)
	fptredC := fptred.NewC("C", 100, "200")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logic.DoFunctionPointed(fptredA)
		logic.DoFunctionPointed(fptredB)
		logic.DoFunctionPointed(fptredC)
	}
}
