package model_callback

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	callbacks2 "github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/eino/utils/callbacks"
)

// NewChatModelCallback todo 2.使用handlerTemplate包含HandlerHelper，handlerTemplate实现了callbacks.Handler，会基于runInfo对不同模块进行callback
/**
type HandlerHelper struct {
	promptHandler      *PromptCallbackHandler
	chatModelHandler   *ModelCallbackHandler
	embeddingHandler   *EmbeddingCallbackHandler
	indexerHandler     *IndexerCallbackHandler
	retrieverHandler   *RetrieverCallbackHandler
	loaderHandler      *LoaderCallbackHandler
	transformerHandler *TransformerCallbackHandler
	toolHandler        *ToolCallbackHandler
	toolsNodeHandler   *ToolsNodeCallbackHandlers
	composeTemplates   map[components.Component]callbacks.Handler
}
*/
func NewChatModelCallback(ctx context.Context) callbacks2.Handler {
	return callbacks.NewHandlerHelper().ChatModel(&callbacks.ModelCallbackHandler{
		OnEnd: func(ctx context.Context, runInfo *callbacks2.RunInfo, output *model.CallbackOutput) context.Context {
			logger.CtxInfof(ctx, "[OnEndFn] %v %+v", *runInfo, output.TokenUsage)
			return ctx
		},
		OnEndWithStreamOutput: func(ctx context.Context, runInfo *callbacks2.RunInfo, output *schema.StreamReader[*model.CallbackOutput]) context.Context {
			return ctx
		},
	}).Handler()
}
