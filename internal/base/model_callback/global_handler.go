package model_callback

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
)

func InitAppendGlobalHandlers(ctx context.Context, handler callbacks.Handler) {
	callbacks.AppendGlobalHandlers(handler)
}
