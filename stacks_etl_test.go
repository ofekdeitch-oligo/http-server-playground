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
