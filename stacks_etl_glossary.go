package main

import "time"

type BatchResult struct {
	librarySummaries []LibrarySummary
}

type LibrarySummary struct {
	maxLastExecutedAt time.Time
}
