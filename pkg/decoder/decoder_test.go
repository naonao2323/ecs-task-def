package decoder

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_Decode_Task_Difinition(t *testing.T) {
	tests := []struct {
		name       string
		definition ecs.TaskDefinition
		format     Format
		expected   []byte
	}{
		{
			name: "succeeded in decoding to ecs task json",
			definition: ecs.TaskDefinition{
				TaskDefinitionArn: "test",
				Family:            "test",
				ExecutionRoleArn:  "test",
			},
			format: Json,
			expected: []byte(`{
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
		},
		{
			name: "succeeded in decoding to ecs task json",
			definition: ecs.TaskDefinition{
				TaskDefinitionArn: "test",
				Family:            "test",
				ExecutionRoleArn:  "test",
			},
			format: Yaml,
			expected: []byte(
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
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			logger, err := zap.NewDevelopment()
			require.NoError(t, err)
			decoder := NewDecoder[ecs.TaskDefinition](logger)
			result, err := decoder.Decode(test.definition, test.format)
			require.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_Decode_Container_Difinition(t *testing.T) {
	tests := []struct {
		name       string
		definition []ecs.ContainerDefinition
		format     Format
		expected   []byte
	}{
		{
			name: "succeeded in decoding to ecs container json",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test",
					Image: "test",
				},
			},
			format: Json,
			expected: []byte(`[
  {
    "name": "test",
    "image": "test",
    "essential": false,
    "restartPolicy": {
      "enabled": false
    }
  }
]`),
		},
		{
			name: "succeeded in decoding to ecs container json",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test",
					Image: "test",
				},
			},
			format: Yaml,
			expected: []byte(`- name: test
  image: test
  essential: false
`),
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			logger, err := zap.NewDevelopment()
			require.NoError(t, err)
			decoder := NewDecoder[[]ecs.ContainerDefinition](logger)
			result, err := decoder.Decode(test.definition, test.format)
			require.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
