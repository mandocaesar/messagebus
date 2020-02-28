package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBus(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
}
