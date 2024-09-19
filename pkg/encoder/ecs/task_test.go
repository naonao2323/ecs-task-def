package ecs

import (
	"ecs-task-def-action/pkg/encoder"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Encode(t *testing.T) {
	in, err := os.ReadFile("../../../testData/task/task-definition.json")
	require.NoError(t, err)
	task := NewTask()
	task.Encode(in, encoder.Json)
}
