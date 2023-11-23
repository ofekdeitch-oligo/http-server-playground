package goat

import (
	"testing"
)

type Suite struct {
	T *testing.T
}

func NewSuite(t *testing.T) Suite {
	return Suite{T: t}
}

func (suite Suite) Test(name string, test func(t *testing.T)) {
	println("")
	println(name)
	println("")

	suite.T.Run(name, test)
}

func (test Suite) Expect(value any) AssertionBuilder {
	return AssertionBuilder{t: test.T, value: value}
}

type AssertionBuilder struct {
	t     *testing.T
	value any
}

func (builder AssertionBuilder) ToEqual(expectedValue any) {
	actualValue := builder.value

	if actualValue != expectedValue {
		builder.t.Errorf("\n\nexpected value to be %s, but was %s\n\n", expectedValue, actualValue)
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
