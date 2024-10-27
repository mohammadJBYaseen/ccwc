package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionsWithNoValueShouldFailed(t *testing.T) {
	_, err := NewCommandParser().Parse([]string{"-"})
	if err == nil {
		t.Error("Expected error")
	}
}

func TestEmptyOptionShouldFailed(t *testing.T) {
	_, err := NewCommandParser().Parse([]string{"-"})
	assert.Error(t, err)
}

func TestValidOptionShouldSuccess(t *testing.T) {
	_, err := NewCommandParser().Parse([]string{"-lwcm"})
	assert.NoError(t, err)
}

func TestWrongOptionShouldFailed(t *testing.T) {
	_, err := NewCommandParser().Parse([]string{"-d"})
	assert.Error(t, err)
}
