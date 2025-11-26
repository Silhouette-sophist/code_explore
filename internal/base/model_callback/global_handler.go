package model_callback

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
)

// InitAppendGlobalHandlers todo 1.作为全局handler，一般程序启动
func InitAppendGlobalHandlers(ctx context.Context, handler callbacks.Handler) {
	callbacks.AppendGlobalHandlers(handler)
}
