package main

import "time"

func handleBatch(stacks []Stack) BatchResult {
	summaries := NewHashMap[Identifier, LibrarySummary](100)

	for _, stack := range stacks {
		for _, entry := range stack.entries {
			uniqueIdentifier := Identifier{
				library: LibraryIdentifier{
					name:    entry.library.name,
					version: entry.library.version,
				},
			}

			ok, existingEntry := summaries.Get(uniqueIdentifier)

			if ok {
				existingEntry.maxLastExecutedAt = maxDate(existingEntry.maxLastExecutedAt, stack.timestamp)
				existingEntry.containerIds.Add(stack.containerId)
				summaries.Set(uniqueIdentifier, existingEntry)
			} else {
				containerIds := NewHashSet[string](10)
				containerIds.Add(stack.containerId)

				summary := LibrarySummary{
					identifier:        uniqueIdentifier,
					maxLastExecutedAt: stack.timestamp,
					containerIds:      containerIds,
				}

				summaries.Set(uniqueIdentifier, summary)
			}
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
