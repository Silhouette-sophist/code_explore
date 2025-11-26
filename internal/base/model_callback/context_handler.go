package model_callback

import (
	"context"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components"
)

// ContextWrapHandlers todo 3.在context中嵌入manager，包含全局handler和当前handler !!! 注意，非compose要自己传递runInfo才能做callback
func ContextWrapHandlers(ctx context.Context, handler callbacks.Handler) context.Context {
	return callbacks.InitCallbacks(ctx, &callbacks.RunInfo{
		Name:      "model",
		Type:      "model",
		Component: components.ComponentOfChatModel,
	}, handler)
}
