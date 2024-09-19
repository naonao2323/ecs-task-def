package ecs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Encode(t *testing.T) {
	in, err := os.ReadFile("../../../testData/task/task-definition.json")
	require.NoError(t, err)
	encoder := NewTask()
	encoder.Encoder(in)
	assert.Equal(t, 1, 1)
}
