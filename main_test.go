package main

import (
	"reflect"
	"testing"
)

var staticTransactions = GenerateStaticTransactions(10, 1_000_00, 100)

const goroutines = 8

func BenchmarkSumCategoryTransactionsMutex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsMutex(staticTransactions, goroutines)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsMutexStandalone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsMutexStandalone(staticTransactions, goroutines)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsChanel(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsChanel(staticTransactions, goroutines)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsChanelStandalone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsChanelStandalone(staticTransactions, goroutines)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsStandalone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsStandalone(staticTransactions)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsMutexAgain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsMutex(staticTransactions, goroutines)
		b.StopTimer()
		if reflect.DeepEqual(result, make(map[int]*Report)) {
			b.Fatal("invalid type result")
		}
		b.StartTimer()
	}
}
