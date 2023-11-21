package main

import "errors"

type Optional struct {
	value *int
}

func (optional Optional) Get() (*int, error) {
	if optional.value != nil {
		return optional.value, nil
	}

	return nil, errors.New("optional does not contain a value")
}

func (optional Optional) Present() bool {
	return optional.value != nil
}

func (optional Optional) GetOrReturn(fallback int) int {
	if optional.Present() {
		return *optional.value
	}

	return fallback
}

func NewOptional(value *int) *Optional {
	optional := new(Optional)
	optional.value = value

	return optional
}
