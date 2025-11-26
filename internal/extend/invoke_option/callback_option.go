package invoke_option

import (
	"code_explore/internal/base/model_callback"
	"context"

	"github.com/cloudwego/eino/compose"
)

func NewInvokeCallbackOption(ctx context.Context) compose.Option {
	return compose.WithCallbacks(model_callback.NewModelFinishTraceCallback(ctx))
}
