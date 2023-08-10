package main

import "time"

type Framework int64

const (
	Pyhthon Framework = iota
	Java
	Go
	Nodejs
)

type Stack struct {
	timestamp time.Time
	framework Framework
	clientId  string
	entries   []StackEntry
}

type StackEntry struct {
	library  LibraryInfo
	function string
}

type LibraryInfo struct {
	name    string
	version string
}
