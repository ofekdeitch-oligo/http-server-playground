package main

import (
	"testing"
	"time"
)

func Test_StacksEtl_ShouldTakeMaxExecutedOfLibrary(t *testing.T) {

	// ARRANGE
	time1 := time.Now()
	time2 := time1.Add(time.Hour * 24)
	time3 := time1.Add(time.Hour * 6)

	library1 := NewLibraryInfoBuilder().WithName("axios").WithVersion("1.0.0").Build()
	entry := NewStackEntryBuilder().WithLibrary(library1).Build()

	stack1 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry}).WithTimestamp(time1).Build()
	stack2 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry}).WithTimestamp(time2).Build()
	stack3 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry}).WithTimestamp(time3).Build()

	// ACT
	result := handleBatch([]Stack{stack1, stack2, stack3})

	// ASSERT
	Test{t}.Expect(len(result.librarySummaries)).ToEqual(1)

	actualSummary := result.librarySummaries[0]
	Test{t}.Expect(actualSummary.maxLastExecutedAt).ToEqual(time2)
}

func Test_StacksEtl_GivenTwoLibraries_ShouldGetMaxLastExecutedAtForEach(t *testing.T) {
	// ARRANGE
	time1 := time.Now()
	time2 := time1.Add(time.Hour * 6)
	time3 := time1.Add(time.Hour * 24)

	library1 := NewLibraryInfoBuilder().WithName("axios").WithVersion("1.0.0").Build()
	library2 := NewLibraryInfoBuilder().WithName("axios").WithVersion("2.0.0").Build()
	entry1 := NewStackEntryBuilder().WithLibrary(library1).Build()
	entry2 := NewStackEntryBuilder().WithLibrary(library2).Build()

	stack1 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry1, entry2, entry2}).WithTimestamp(time1).Build()
	stack2 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry1, entry1, entry1, entry2}).WithTimestamp(time2).Build()
	stack3 := NewStackBuilder().WithFramework(Nodejs).WithEntries([]StackEntry{entry2}).WithTimestamp(time3).Build()

	// ACT
	result := handleBatch([]Stack{stack1, stack2, stack3})

	// ASSERT
	Test{t}.Expect(len(result.librarySummaries)).ToEqual(2)

	ok1, actualSummary1 := first(
		result.librarySummaries,
		func(summary LibrarySummary) bool { return summary.identifier.library.version == "1.0.0" },
	)

	Test{t}.Expect(ok1).ToEqual(true)
	Test{t}.Expect(actualSummary1.maxLastExecutedAt).ToEqual(time2)

	ok2, actualSummary2 := first(
		result.librarySummaries,
		func(summary LibrarySummary) bool { return summary.identifier.library.version == "2.0.0" },
	)

	Test{t}.Expect(ok2).ToEqual(true)
	Test{t}.Expect(actualSummary2.maxLastExecutedAt).ToEqual(time3)
}

func first[T any](arr []T, predicate func(T) bool) (bool, T) {
	for _, item := range arr {
		if predicate(item) {
			return true, item
		}
	}

	return false, noValue[T]()
}
