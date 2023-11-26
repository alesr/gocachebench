package gocachebench

import (
	"testing"
)

func BenchmarkCache_key1(b *testing.B) {
	b.StopTimer()

	cb := newCacheBench()

	b.Run("key1", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey1()
		}
		b.StopTimer()
	})

	b.Run("key1Async", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey1Async()
		}
		b.StopTimer()
	})

	b.Run("key2", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey2()

		}
		b.StopTimer()
	})

	b.Run("key2Async", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey2Async()

		}
		b.StopTimer()
	})

	b.Run("key3", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey3()
		}
		b.StopTimer()
	})

	b.Run("key3Async", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey3Async()
		}
		b.StopTimer()
	})

	b.Run("key4", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey4()
		}
		b.StopTimer()
	})

	b.Run("key4Sync", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey4Async()
		}
		b.StopTimer()
	})

	b.Run("key5", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey5()
		}
		b.StopTimer()
	})

	b.Run("key5Async", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setKey5Async()
		}
		b.StopTimer()
	})

	b.Run("all", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setAll()
		}
		b.StopTimer()
	})

	b.Run("allAsync", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_ = cb.setAllAsync()
		}
		b.StopTimer()
	})
}
