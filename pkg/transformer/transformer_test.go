package transformer

import (
	"testing"

	"github.com/naonao2323/ecs-task-def/pkg/plovider/ecs"

	"github.com/stretchr/testify/assert"
)

func Test_Transformer_Task_Definition(t *testing.T) {
	tests := []struct {
		name       string
		tag        string
		appName    string
		definition ecs.TaskDefinition
		expected   ecs.TaskDefinition
	}{
		{
			name:    "appName matches, Tag does not match",
			tag:     "target1",
			appName: "test1",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:test1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:target1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
		},
		{
			name:    "appName and Tag both match",
			tag:     "test1",
			appName: "test1",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:test1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:test1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
		},
		{
			name:    "appName does not match, Tag matches",
			tag:     "test1",
			appName: "test10",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:test1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:test1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
		},
		{
			name:    "Neither appName nor Tag match",
			tag:     "target1",
			appName: "test1",
			definition: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:target1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
			expected: ecs.TaskDefinition{
				ContainerDefinitions: []ecs.ContainerDefinition{
					{
						Name:  "test1",
						Image: "test:target1",
					},
					{
						Name:  "test2",
						Image: "test:test2",
					},
					{
						Name:  "test3",
						Image: "test:test3",
					},
				},
			},
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			transformer := NewTransformer[ecs.TaskDefinition]()
			result := transformer.Transform(test.tag, test.appName, test.definition)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_Transformer_Container_Definition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		tag        string
		appName    string
		definition []ecs.ContainerDefinition
		expected   []ecs.ContainerDefinition
	}{
		{
			name:    "appName matches, Tag does not match",
			tag:     "target1",
			appName: "test1",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:test1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
			expected: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:target1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
		},
		{
			name:    "appName and Tag both match",
			tag:     "test1",
			appName: "test1",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:test1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
			expected: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:test1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
		},
		{
			name:    "appName does not match, Tag matches",
			tag:     "test1",
			appName: "test10",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:test1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
			expected: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:test1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
		},
		{
			name:    "Neither appName nor Tag match",
			tag:     "target1",
			appName: "test1",
			definition: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:target1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
			expected: []ecs.ContainerDefinition{
				{
					Name:  "test1",
					Image: "test:target1",
				},
				{
					Name:  "test2",
					Image: "test:test2",
				},
				{
					Name:  "test3",
					Image: "test:test3",
				},
			},
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			transformer := NewTransformer[[]ecs.ContainerDefinition]()
			result := transformer.Transform(test.tag, test.appName, test.definition)
			assert.Equal(t, test.expected, result)
		})
	}
}
