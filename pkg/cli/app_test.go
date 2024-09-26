package cli

import (
	"context"
	"ecs-task-def-action/pkg/decoder"
	"ecs-task-def-action/pkg/encoder"
	"ecs-task-def-action/pkg/git"
	"ecs-task-def-action/pkg/github"
	"ecs-task-def-action/pkg/plovider/ecs"
	"ecs-task-def-action/pkg/transformer"
	"errors"
	"testing"

	mock_decoder "ecs-task-def-action/mock/decoder"
	mock_encoder "ecs-task-def-action/mock/encoder"
	mock_git "ecs-task-def-action/mock/git"
	mock_github "ecs-task-def-action/mock/github"
	mock_transformer "ecs-task-def-action/mock/transformer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func Test_execute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	in := []byte("test")
	out := []byte("test")
	app := "test"
	tag := "test1"
	path := "test"
	githubUrl := "https:github.com/test"
	format := encoder.Json
	definition := ecs.TaskDefinition{}
	transformed := ecs.TaskDefinition{}
	tests := []struct {
		name         string
		encoder      func() encoder.Encoder[ecs.TaskDefinition]
		decoder      func() decoder.Decoder[ecs.TaskDefinition]
		transformer  func() transformer.Transformer[ecs.TaskDefinition]
		outputer     func(in []byte, tag, path string) error
		gitClient    func() git.Git
		githubClient func(ctx context.Context) github.Github
		expected     error
	}{
		{
			name: "fail to encode task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, errors.New("encode error"))
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)

				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("encode error"),
		},
		{
			name: "succeeded in encoding empty task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(nil, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)

				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("empty definition"),
		},
		{
			name: "fail to decode task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(nil, errors.New("fail to decode"))
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)

				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to decode"),
		},
		{
			name: "fail to rewrite task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return errors.New("fail to rewrite file")
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)

				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to rewrite file"),
		},
		{
			name: "fail to add task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(errors.New("fail to git add"))
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to git add"),
		},
		{
			name: "fail to commit task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(nil)
				mockGitClient.EXPECT().Commit(tag).Return(errors.New("fail to git commit"))
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to git commit"),
		},
		{
			name: "fail to checkout task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(nil)
				mockGitClient.EXPECT().Commit(tag).Return(nil)
				mockGitClient.EXPECT().CheckOut(tag).Return(errors.New("fail to git checkout"))
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to git checkout"),
		},
		{
			name: "fail to puah task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(nil)
				mockGitClient.EXPECT().Commit(tag).Return(nil)
				mockGitClient.EXPECT().CheckOut(tag).Return(nil)
				mockGitClient.EXPECT().Push(tag).Return(errors.New("fail to git push"))
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				return mockGithubClient
			},
			expected: errors.New("fail to git push"),
		},
		{
			name: "fail to create pull request task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(nil)
				mockGitClient.EXPECT().Commit(tag).Return(nil)
				mockGitClient.EXPECT().CheckOut(tag).Return(nil)
				mockGitClient.EXPECT().Push(tag).Return(nil)
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				mockGithubClient.EXPECT().CreatePullRequest(ctx, tag, tag).Return(errors.New("fail to create pull request"))
				return mockGithubClient
			},
			expected: errors.New("fail to create pull request"),
		},
		{
			name: "succeeded in creating pull request task definition",
			encoder: func() encoder.Encoder[ecs.TaskDefinition] {
				mockEncoder := mock_encoder.NewMockEncoder[ecs.TaskDefinition](mockCtrl)
				mockEncoder.EXPECT().Encode(in, format).Return(&definition, nil)
				return mockEncoder
			},
			decoder: func() decoder.Decoder[ecs.TaskDefinition] {
				mockDecoder := mock_decoder.NewMockDecoder[ecs.TaskDefinition](mockCtrl)
				mockDecoder.EXPECT().Decode(transformed, convertFormat(format)).Return(out, nil)
				return mockDecoder
			},
			transformer: func() transformer.Transformer[ecs.TaskDefinition] {
				mockTransformer := mock_transformer.NewMockTransformer[ecs.TaskDefinition](mockCtrl)
				mockTransformer.EXPECT().Transform(tag, app, definition).Return(transformed)
				return mockTransformer
			},
			outputer: func(in []byte, tag, path string) error {
				return nil
			},
			gitClient: func() git.Git {
				mockGitClient := mock_git.NewMockGit(mockCtrl)
				mockGitClient.EXPECT().Add(path).Return(nil)
				mockGitClient.EXPECT().Commit(tag).Return(nil)
				mockGitClient.EXPECT().CheckOut(tag).Return(nil)
				mockGitClient.EXPECT().Push(tag).Return(nil)
				return mockGitClient
			},
			githubClient: func(ctx context.Context) github.Github {
				mockGithubClient := mock_github.NewMockGithub(mockCtrl)
				mockGithubClient.EXPECT().CreatePullRequest(ctx, tag, tag).Return(nil)
				return mockGithubClient
			},
			expected: nil,
		},
	}

	for _, _test := range tests {
		test := _test
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			result := execute[ecs.TaskDefinition](
				ctx,
				logger,
				in,
				app,
				tag,
				path,
				githubUrl,
				format,
				test.transformer(),
				test.encoder(),
				test.decoder(),
				test.outputer,
				test.gitClient(),
				test.githubClient(ctx),
			)
			assert.Equal(t, test.expected, result)
		})
	}
}
