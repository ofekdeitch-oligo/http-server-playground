package main

import (
	"reflect"

	"github.com/mitchellh/hashstructure/v2"
)

type HashMap[K comparable, V any] struct {
	buckets [][]Tuple[K, V]
}

type Tuple[K comparable, V any] struct {
	key   K
	value V
}

func (hashmap HashMap[K, V]) GetSlot(key K) int {
	hash, err := hashstructure.Hash(key, hashstructure.FormatV2, nil)

	if err != nil {
		panic(err)
	}

	return abs(int(hash)) % len(hashmap.buckets)
}

func (hashmap HashMap[K, V]) Set(key K, value V) {
	slot := hashmap.GetSlot(key)
	bucket := hashmap.buckets[slot]

	ok, existingTuple, _ := findKeyInBucket(bucket, key)

	if ok {
		existingTuple.value = value
		return
	}

	tuple := Tuple[K, V]{key, value}
	bucket = append(bucket, tuple)

	hashmap.buckets[slot] = bucket
}

func (hashmap HashMap[K, V]) Get(key K) (bool, V) {
	slot := hashmap.GetSlot(key)
	bucket := hashmap.buckets[slot]

	ok, tuple, _ := findKeyInBucket(bucket, key)

	if ok {
		return true, tuple.value
	}

	return false, noValue[V]()
}

func (hashmap HashMap[K, V]) Delete(key K) {
	slot := hashmap.GetSlot(key)
	bucket := hashmap.buckets[slot]

	ok, _, index := findKeyInBucket(bucket, key)

	if ok {
		bucketWithoutTuple := removeFromSlice(bucket, index)
		hashmap.buckets[slot] = bucketWithoutTuple
	}
}

func NewHashMap[K comparable, V any](size int) HashMap[K, V] {
	buckets := make([][]Tuple[K, V], size)
	return HashMap[K, V]{buckets}
}

func findKeyInBucket[K comparable, V any](bucket []Tuple[K, V], key K) (bool, *Tuple[K, V], int) {
	for index, tuple := range bucket {

		if reflect.DeepEqual(tuple.key, key) {
			return true, &bucket[index], index
		}
	}

	return false, noValue[*Tuple[K, V]](), noValue[int]()
}

func removeFromSlice[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (hashmap HashMap[K, V]) Values() []V {
	result := make([]V, 0)

	for _, bucket := range hashmap.buckets {
		for _, tuple := range bucket {
			result = append(result, tuple.value)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
