package main

import "testing"

func Test_Optional1_Present(t *testing.T) {
	optional := NewOptional(nil)

	Test{t}.Expect(optional.Present()).ToBeFalse()
}

func Test_Optional2_Present(t *testing.T) {
	val := 5
	optional := NewOptional(&val)

	Test{t}.Expect(optional.Present()).ToBeTrue()
}

func Test_Optional1_GetOrReturn(t *testing.T) {
	val := 4
	optional := NewOptional(&val)

	Test{t}.Expect(optional.GetOrReturn(2)).ToEqual(4)
}

func Test_Optional2_GetOrReturn(t *testing.T) {
	optional := NewOptional(nil)

	Test{t}.Expect(optional.GetOrReturn(2)).ToEqual(2)
}
