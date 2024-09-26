package encoder

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_Container_Definition(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		format   Format
		expected *[]ecs.ContainerDefinition
	}{
		{
			name: "succceeded in encoding to ecs containers definition json",
			in: []byte(`[
  {
    "name": "test",
    "image": "test",
    "essential": false,
    "restartPolicy": {
      "enabled": false
    }
  }
]`),
			format: Json,
			expected: &[]ecs.ContainerDefinition{
				{
					Name:  "test",
					Image: "test",
				},
			},
		},
		{
			name: "succeeded in encoding to ecs containers definition yaml",
			in: []byte(`- name: test
  image: test
  essential: false
`),
			format: Yaml,
			expected: &[]ecs.ContainerDefinition{
				{
					Name:  "test",
					Image: "test",
				},
			},
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			logger, err := zap.NewDevelopment()
			require.NoError(t, err)
			encoder := NewEncoder[[]ecs.ContainerDefinition](logger)
			result, err := encoder.Encode(test.in, test.format)
			require.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_Task_Definition(t *testing.T) {
	tests := []struct {
		name     string
		in       []byte
		format   Format
		expected *ecs.TaskDefinition
	}{
		{
			name: "succceeded in encoding to ecs task definition json",
			in: []byte(`{
  "taskDefinitionArn": "test",
  "containerDefinitions": null,
  "family": "test",
  "executionRoleArn": "test",
  "revision": 0,
  "status": "",
  "requiresAttributes": null,
  "compatibilities": null,
  "registeredAt": "",
  "registeredBy": ""
}`),
			format: Json,
			expected: &ecs.TaskDefinition{
				TaskDefinitionArn: "test",
				Family:            "test",
				ExecutionRoleArn:  "test",
			},
		},
		{
			name: "succceeded in encoding to ecs task definition json",
			in: []byte(
				`taskDefinitionArn: test
containerDefinitions: []
family: test
executionRoleArn: test
revision: 0
status: ""
requiresAttributes: []
compatibilities: []
registeredAt: ""
registeredBy: ""
`),
			format: Yaml,
			expected: &ecs.TaskDefinition{
				TaskDefinitionArn: "test",
				Family:            "test",
				ExecutionRoleArn:  "test",
			},
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			logger, err := zap.NewDevelopment()
			require.NoError(t, err)
			encoder := NewEncoder[ecs.TaskDefinition](logger)
			result, err := encoder.Encode(test.in, test.format)
			require.NoError(t, err)
			assert.Equal(t, test.expected.TaskDefinitionArn, result.TaskDefinitionArn)
			assert.Equal(t, test.expected.Family, result.Family)
			assert.Equal(t, test.expected.ExecutionRoleArn, result.ExecutionRoleArn)
		})
	}
}
