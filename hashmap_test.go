package main

import (
	"testing"
)

func Test_HashMap_Set(t *testing.T) {

	// ARRANGE

	m := NewHashMap[LibraryIdentifier, int](10)

	library1 := LibraryIdentifier{
		name:    "Spring",
		version: "1.0.0",
	}

	// ACT

	m.Set(library1, 5)

	// ASSERT
	ok, value := m.Get(library1)
	Test{t}.Expect(ok).ToBeTrue()
	Test{t}.Expect(value).ToEqual(5)
}

func Test_HashMap_SetTwice(t *testing.T) {

	// ARRANGE

	m := NewHashMap[LibraryIdentifier, int](10)

	library1 := LibraryIdentifier{
		name:    "Spring",
		version: "1.0.0",
	}

	library2 := LibraryIdentifier{
		name:    "Spring",
		version: "1.0.0",
	}

	// ACT

	m.Set(library1, 5)
	m.Set(library2, 6)

	// ASSERT
	ok, value := m.Get(library1)
	Test{t}.Expect(ok).ToBeTrue()
	Test{t}.Expect(value).ToEqual(6)

	ok, value = m.Get(library2)
	Test{t}.Expect(ok).ToBeTrue()
	Test{t}.Expect(value).ToEqual(6)
}

func Test_HashMap_NoValue(t *testing.T) {

	// ARRANGE

	m := NewHashMap[LibraryIdentifier, int](10)

	library1 := LibraryIdentifier{
		name:    "Spring",
		version: "1.0.0",
	}

	// ACT

	// do nothing

	// ASSERT
	ok, _ := m.Get(library1)
	Test{t}.Expect(ok).ToBeFalse()
}

func Test_HashMap_Delete(t *testing.T) {
	// ARRANGE

	m := NewHashMap[LibraryIdentifier, int](10)

	library1 := LibraryIdentifier{
		name:    "Spring",
		version: "1.0.0",
	}

	// ACT
	m.Set(library1, 5)
	m.Delete(library1)

	// do nothing

	// ASSERT
	ok, _ := m.Get(library1)
	Test{t}.Expect(ok).ToBeFalse()
}

func Test_HashMap_Values(t *testing.T) {
	// ARRANGE

	m := NewHashMap[int, int](10)

	// ACT
	m.Set(1, 5)
	m.Set(2, 4)
	m.Set(3, 9)

	// ASSERT
	values := m.Values()
	Test{t}.Expect(len(values)).ToEqual(3)

	ok, _ := find(values, 5)
	Test{t}.Expect(ok).ToBeTrue()

	ok, _ = find(values, 4)
	Test{t}.Expect(ok).ToBeTrue()

	ok, _ = find(values, 9)
	Test{t}.Expect(ok).ToBeTrue()
}

type LibraryIdentifier struct {
	name    string
	version string
}

func find(arr []int, value int) (bool, int) {
	for i, v := range arr {
		if v == value {
			return true, i
		}
	}

	return false, 0
}
