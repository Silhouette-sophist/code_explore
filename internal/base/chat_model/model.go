package chat_model

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

type ModelType string

const (
	DoubaoThinking ModelType = "ark"
	DeepSeek       ModelType = "deep_seek"
	deepSeekModel            = "deepseek-v3-1-terminus"
	doubaoModel              = "doubao-seed-1-6-thinking-250715"
	arkApiKey                = "1b943f16-1653-4cb0-8580-3544855d5c7e"
)

// NewChatModel 创建chatModel实例，当前均有ark提供模型服务
func NewChatModel(ctx context.Context, modelType ModelType) (*ark.ChatModel, error) {
	modelConfig := &ark.ChatModelConfig{
		APIKey: arkApiKey,
	}
	switch modelType {
	case DoubaoThinking:
		modelConfig.Model = doubaoModel
	case DeepSeek:
		modelConfig.Model = deepSeekModel
	}
	return ark.NewChatModel(ctx, modelConfig)
}
