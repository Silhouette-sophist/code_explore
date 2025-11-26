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

/*
1.注入 RunInfo
RunInfo 也需要注入到 Context 中，才会在触发回调时给到 Handler。
a.Graph 托管 RunInfo
Graph 会为内部所有的 Node 自动注入 RunInfo。机制是每个 Node 的运行，都是一个新的子 Context，Graph 向这个新的 Context 中注入对应 Node 的 RunInfo。
b.在 Graph 外注入 RunInfo
不想使用 Graph，但却想使用 Callback，则：
通过 InitCallbacks(ctx context.Context, info *RunInfo, handlers ...Handler) 获取一个新的 Context 并注入 Handlers 以及 RunInfo。
通过 ReuseHandlers(ctx context.Context, info *RunInfo) 来获取一个新的 Context，复用之前 Context 中的 Handler，并设置新的 RunInfo。
*/
