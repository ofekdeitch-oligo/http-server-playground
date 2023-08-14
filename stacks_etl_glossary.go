package main

import (
	"time"
)

type BatchResult struct {
	librarySummaries []LibrarySummary
}

type LibrarySummary struct {
	identifier        Identifier
	maxLastExecutedAt time.Time
	containerIds      HashSet[string]
}

type Identifier struct {
	library LibraryIdentifier
}

type LibraryIdentifier struct {
	name    string
	version string
}
