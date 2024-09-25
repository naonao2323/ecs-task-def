package ecs

import (
	"ecs-task-def-action/pkg/plovider/ecs"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Task(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		tag        string
		appName    string
		definition ecs.TaskDefinition
		expected   ecs.TaskDefinition
	}{
		{
			name:    "appName does not match any appNames in the definition",
			tag:     "target-tag",
			appName: "target-appName",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:2",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:2",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
		},
		{
			name:    "appName matches a appName in the definition",
			tag:     "target-tag",
			appName: "test2",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:2",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:target-tag",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
		},
		{
			name:    "appName matches an appName in the definition, but the tags are identical",
			tag:     "2",
			appName: "test2",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:2",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:1",
					},
					{
						Name:  "test2",
						Image: "test:2",
					},
					{
						Name:  "test3",
						Image: "test:3",
					},
				},
			},
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			transformer := NewTaskTransformer()
			result := transformer.Transform(test.tag, test.appName, test.definition)
			require.Equal(t, test.expected, result)
		})
	}
}
