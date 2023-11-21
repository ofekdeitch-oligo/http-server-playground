package main

type HashSet[V comparable] struct {
	values HashMap[V, bool]
}

func NewHashSet[V comparable](size int) HashSet[V] {
	return HashSet[V]{
		values: NewHashMap[V, bool](size),
	}
}

func (hashset HashSet[V]) Add(value V) {
	hashset.values.Set(value, true)
}

func (hashset HashSet[V]) Contains(value V) bool {
	ok, _ := hashset.values.Get(value)
	return ok
}
