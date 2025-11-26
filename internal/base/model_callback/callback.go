package model_callback

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// NewModelFinishTraceCallback todo 1.对全部组件加上handler，要自己解析类型
func NewModelFinishTraceCallback(ctx context.Context) callbacks.Handler {
	return callbacks.NewHandlerBuilder().OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
		if callbackOutput, ok := output.(*model.CallbackOutput); ok {
			logger.CtxInfof(ctx, "[OnEndFn] %v %+v", *info, callbackOutput.TokenUsage)
		}
		return ctx
	}).OnEndWithStreamOutputFn(func(ctx context.Context, info *callbacks.RunInfo, output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
		return ctx
	}).Build()
}
