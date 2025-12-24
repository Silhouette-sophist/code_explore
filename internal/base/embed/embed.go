package embed

import (
	"code_explore/internal/base/config"
	"context"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

func NewArcEmbed(ctx context.Context) (*ark.Embedder, error) {
	return ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: config.GetEnv(config.ApiKey, ""),
		Model:  config.GetEnv(config.EmbedModelId, ""),
	})
}
