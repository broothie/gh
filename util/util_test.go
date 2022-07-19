package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeepCopy(t *testing.T) {
	thing := map[string]any{"a": 1, "b": []int{2, 3}}

	assert.Equal(t,
		map[string]any{"a": 1, "b": []int{2, 3}},
		DeepCopy(thing),
	)
}
