package main

import "testing"

type Test struct {
	t *testing.T
}

func (test Test) Expect(value any) AssertionBuilder {
	return AssertionBuilder{t: test.t, value: value}
}

type AssertionBuilder struct {
	t     *testing.T
	value any
}

func (builder AssertionBuilder) ToEqual(expectedValue any) {
	actualValue := builder.value

	if actualValue != expectedValue {
		builder.t.Errorf("\n\nexpected value to be %d, but was %d\n\n", expectedValue, actualValue)
	}
}

func (builder AssertionBuilder) ToBeTrue() {
	actualValue := builder.value

	if actualValue != true {
		builder.t.Errorf("\n\nexpected value to be true, but was %d\n\n", actualValue)
	}
}

func (builder AssertionBuilder) ToBeFalse() {
	actualValue := builder.value

	if actualValue != false {
		builder.t.Errorf("\n\nexpected value to be false, but was %d\n\n", actualValue)
	}
}

func (builder AssertionBuilder) ToBeNil() {
	actualValue := builder.value

	if actualValue != nil {
		builder.t.Errorf("\n\nexpected value to be nil, but was %d\n\n", actualValue)
	}
}
