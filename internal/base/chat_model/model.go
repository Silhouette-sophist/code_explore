package chat_model

import (
	"code_explore/internal/base/config"
	"context"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

// NewChatModel 创建chatModel实例，当前均有ark提供模型服务
func NewChatModel(ctx context.Context) (*ark.ChatModel, error) {
	modelConfig := &ark.ChatModelConfig{
		APIKey: config.GetEnv(config.ApiKey, ""),
		Model:  config.GetEnv(config.ModelId, ""),
	}
	return ark.NewChatModel(ctx, modelConfig)
}
