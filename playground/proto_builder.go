package main

import "time"

type LibraryInfoBuilder struct {
	value LibraryInfo
}

func NewLibraryInfoBuilder() LibraryInfoBuilder {
	value := LibraryInfo{
		name:    newUuid(),
		version: newUuid(),
	}

	return LibraryInfoBuilder{value: value}
}

func (builder LibraryInfoBuilder) WithName(name string) LibraryInfoBuilder {
	builder.value.name = name
	return builder
}

func (builder LibraryInfoBuilder) WithVersion(version string) LibraryInfoBuilder {
	builder.value.version = version
	return builder
}

func (builder LibraryInfoBuilder) Build() LibraryInfo {
	return builder.value
}

type StackEntryBuilder struct {
	value StackEntry
}

func NewStackEntryBuilder() StackEntryBuilder {
	value := StackEntry{
		library:  NewLibraryInfoBuilder().Build(),
		function: newUuid(),
	}

	return StackEntryBuilder{value: value}
}

func (builder StackEntryBuilder) WithLibrary(library LibraryInfo) StackEntryBuilder {
	builder.value.library = library
	return builder
}

func (builder StackEntryBuilder) WithFunction(function string) StackEntryBuilder {
	builder.value.function = function
	return builder
}

func (builder StackEntryBuilder) Build() StackEntry {
	return builder.value
}

type StackBuilder struct {
	value Stack
}

func NewStackBuilder() StackBuilder {
	value := Stack{
		timestamp:   time.Now(),
		framework:   Pyhthon,
		clientId:    newUuid(),
		entries:     []StackEntry{},
		containerId: newUuid(),
	}

	return StackBuilder{value: value}
}

func (builder StackBuilder) WithTimestamp(timestamp time.Time) StackBuilder {
	builder.value.timestamp = timestamp
	return builder
}

func (builder StackBuilder) WithFramework(framework Framework) StackBuilder {
	builder.value.framework = framework
	return builder
}

func (builder StackBuilder) WithClientId(clientId string) StackBuilder {
	builder.value.clientId = clientId
	return builder
}

func (builder StackBuilder) WithEntries(entries []StackEntry) StackBuilder {
	builder.value.entries = entries
	return builder
}

func (builder StackBuilder) Build() Stack {
	return builder.value
}
