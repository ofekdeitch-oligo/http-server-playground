package main

import "time"

func handleBatch(stacks []Stack) BatchResult {
	summaries := NewHashMap[string, LibrarySummary](100)

	for _, stack := range stacks {
		uniqueIdentifier := ""

		ok, existingEntry := summaries.Get(uniqueIdentifier)

		if ok {
			existingEntry.maxLastExecutedAt = maxDate(existingEntry.maxLastExecutedAt, stack.timestamp)
			summaries.Set(uniqueIdentifier, existingEntry)

		} else {
			summary := LibrarySummary{
				maxLastExecutedAt: stack.timestamp,
			}

			summaries.Set(uniqueIdentifier, summary)
		}
	}

	return BatchResult{
		librarySummaries: summaries.Values(),
	}
}

func maxDate(date1 time.Time, date2 time.Time) time.Time {
	if date1.After(date2) {
		return date1
	}

	return date2
}
