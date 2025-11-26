package invoke_option

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/flow/agent/react"
	template "github.com/cloudwego/eino/utils/callbacks"
)

// ReactCallback 仅仅针对react场景
func ReactCallback(ctx context.Context, modelHandler *template.ModelCallbackHandler, toolHandler *template.ToolCallbackHandler) callbacks.Handler {
	return react.BuildAgentCallback(modelHandler, toolHandler)
}

func ReactDefaultCallback(ctx context.Context) callbacks.Handler {
	return ReactCallback(ctx, &template.ModelCallbackHandler{
		OnEnd: func(ctx context.Context, runInfo *callbacks.RunInfo, output *model.CallbackOutput) context.Context {
			logger.CtxInfof(ctx, "[OnFn] info %v %v", runInfo, output)
			return ctx
		},
	}, &template.ToolCallbackHandler{
		OnEnd: func(ctx context.Context, info *callbacks.RunInfo, output *tool.CallbackOutput) context.Context {
			logger.CtxInfof(ctx, "[OnFn] info %v %v", info, output)
			return ctx
		},
	})
}
