package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSomething(t *testing.T) {
	assertError(t)
	assertNil(t)

	a := "Hello"
	b := "Hello"
	assertEqual(t, a, b)
}

func assertError(t *testing.T) {
	var err error
	assert.NoError(t, err, "no err")
}

func assertNil(t *testing.T) {
	assert.Nil(t, nil, "should be nil")
}

func assertEqual(t *testing.T, a, b string) {
	assert.Equal(t, a, b, "The two words should be the same.")
}
