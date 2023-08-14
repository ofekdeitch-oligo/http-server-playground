package main

import "testing"

func Test_HashSet_Add(t *testing.T) {

	// ARRANGE

	m := NewHashSet[LibraryIdentifier2](10)

	library1 := LibraryIdentifier2{
		name:    "Spring",
		version: "1.0.0",
	}

	// ACT

	m.Add(library1)

	// ASSERT
	ok := m.Contains(library1)
	Test{t}.Expect(ok).ToBeTrue()
}
