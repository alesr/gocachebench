package gocachebench

import (
	"context"
	"log"

	"github.com/dgraph-io/ristretto"
	gc "github.com/eko/gocache/lib/v4/cache"
	gcr "github.com/eko/gocache/store/ristretto/v4"
)

var ctx = context.TODO()

type keyCache[T any] struct {
	*gc.Cache[T]
}

func newKeyCache[T any](store *gcr.RistrettoStore) *keyCache[T] {
	return &keyCache[T]{
		Cache: gc.New[T](store),
	}
}

type cacheBench struct {
	keyCache1 *keyCache[*value1]
	keyCache2 *keyCache[*value2]
	keyCache3 *keyCache[*value3]
	keyCache4 *keyCache[*value4]
	keyCache5 *keyCache[*value5]
}

func newCacheBench() *cacheBench {
	risConfig := ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 20, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	}

	ris, err := ristretto.NewCache(&risConfig)
	if err != nil {
		log.Fatalln("failed to create cache:", err)
	}

	store := gcr.NewRistretto(ris)

	return &cacheBench{
		keyCache1: newKeyCache[*value1](store),
		keyCache2: newKeyCache[*value2](store),
		keyCache3: newKeyCache[*value3](store),
		keyCache4: newKeyCache[*value4](store),
		keyCache5: newKeyCache[*value5](store),
	}
}

func (cb *cacheBench) setKey1() error {
	return cb.keyCache1.Set(ctx, testDataKey1, &testDataValue1)
}

func (cb *cacheBench) setKey1Async() error {
	go func() {
		_ = cb.keyCache1.Set(ctx, testDataKey1, &testDataValue1)
	}()
	return nil
}

func (cb *cacheBench) setKey2() error {
	return cb.keyCache2.Set(ctx, testDataKey2, &testDataValue2)
}

func (cb *cacheBench) setKey2Async() error {
	go func() {
		_ = cb.keyCache2.Set(ctx, testDataKey2, &testDataValue2)
	}()
	return nil
}

func (cb *cacheBench) setKey3() error {
	return cb.keyCache3.Set(ctx, testDataKey3, &testDataValue3)
}

func (cb *cacheBench) setKey3Async() error {
	go func() {
		_ = cb.keyCache3.Set(ctx, testDataKey3, &testDataValue3)
	}()
	return nil
}

func (cb *cacheBench) setKey4() error {
	return cb.keyCache4.Set(ctx, testDataKey4, &testDataValue4)
}

func (cb *cacheBench) setKey4Async() error {
	go func() {
		_ = cb.keyCache4.Set(ctx, testDataKey4, &testDataValue4)
	}()
	return nil
}

func (cb *cacheBench) setKey5() error {
	return cb.keyCache5.Set(ctx, testDataKey5, &testDataValue5)
}

func (cb *cacheBench) setKey5Async() error {
	go func() {
		_ = cb.keyCache5.Set(ctx, testDataKey5, &testDataValue5)
	}()
	return nil
}

func (cb *cacheBench) setAll() error {
	_ = cb.keyCache1.Set(ctx, testDataKey1, &testDataValue1)
	_ = cb.keyCache2.Set(ctx, testDataKey2, &testDataValue2)
	_ = cb.keyCache2.Set(ctx, testDataKey2Alt, &testDataValue2Alt)
	_ = cb.keyCache3.Set(ctx, testDataKey3, &testDataValue3)
	_ = cb.keyCache4.Set(ctx, testDataKey4, &testDataValue4)
	_ = cb.keyCache5.Set(ctx, testDataKey5, &testDataValue5)
	return nil
}

func (cb *cacheBench) setAllAsync() error {
	go func() {
		_ = cb.keyCache1.Set(ctx, testDataKey1, &testDataValue1)
		_ = cb.keyCache2.Set(ctx, testDataKey2, &testDataValue2)
		_ = cb.keyCache2.Set(ctx, testDataKey2Alt, &testDataValue2Alt)
		_ = cb.keyCache3.Set(ctx, testDataKey3, &testDataValue3)
		_ = cb.keyCache3.Set(ctx, testDataKey3, &testDataValue3Alt)
		_ = cb.keyCache4.Set(ctx, testDataKey4, &testDataValue4)
		_ = cb.keyCache5.Set(ctx, testDataKey5, &testDataValue5)
		_ = cb.keyCache5.Set(ctx, testDataKey5, &testDataValue5Alt)
	}()
	return nil
}
