package logic_test

import (
	"testing"

	"cchp/direct"
	"cchp/fptred"
	"cchp/interfaced"
	"cchp/logic"
	"cchp/tabled"
	"cchp/util"
)

var first, second, third string

func BenchmarkDoInterfaced(b *testing.B) {
	interfacedA := interfaced.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	interfacedB := interfaced.NewB("B", util.RandInt(), util.RandInt())
	interfacedC := interfaced.NewC("C", util.RandInt(), util.RandString(10))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		first = logic.DoInterfaced(interfacedA)
		second = logic.DoInterfaced(interfacedB)
		third = logic.DoInterfaced(interfacedC)
	}

	_, _, _ = first, second, third
}

func BenchmarkDoTabled(b *testing.B) {
	tabledA := tabled.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	tabledB := tabled.NewB("B", util.RandInt(), util.RandInt())
	tabledC := tabled.NewC("C", util.RandInt(), util.RandString(10))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		first = logic.DoTabled(tabledA)
		second = logic.DoTabled(tabledB)
		third = logic.DoTabled(tabledC)
	}

	_, _, _ = first, second, third
}

func BenchmarkDoDirectTabled(b *testing.B) {
	tabledA := tabled.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	tabledB := tabled.NewB("B", util.RandInt(), util.RandInt())
	tabledC := tabled.NewC("C", util.RandInt(), util.RandString(10))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		first = logic.DoDirectTabled(tabledA)
		second = logic.DoDirectTabled(tabledB)
		third = logic.DoDirectTabled(tabledC)
	}

	_, _, _ = first, second, third
}

func BenchmarkDoFunctionPointed(b *testing.B) {
	fptredA := fptred.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	fptredB := fptred.NewB("B", util.RandInt(), util.RandInt())
	fptredC := fptred.NewC("C", util.RandInt(), util.RandString(10))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		first = logic.DoFunctionPointed(fptredA)
		second = logic.DoFunctionPointed(fptredB)
		third = logic.DoFunctionPointed(fptredC)
	}

	_, _, _ = first, second, third
}

func BenchmarkDirect(b *testing.B) {
	directA := direct.NewA("A", util.RandInt(), util.RandInt(), util.RandInt())
	directB := direct.NewB("B", util.RandInt(), util.RandInt())
	directC := direct.NewC("C", util.RandInt(), util.RandString(10))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		first = logic.DoDirectA(directA)
		second = logic.DoDirectB(directB)
		third = logic.DoDirectC(directC)
	}

	_ = first
}
