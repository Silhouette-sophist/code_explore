package invoke_option

import (
	"code_explore/internal/base/model_callback"
	"context"

	"github.com/cloudwego/eino/compose"
)

// NewInvokeCallbackOption todo 2.作为调用option补充！！！
func NewInvokeCallbackOption(ctx context.Context) compose.Option {
	return compose.WithCallbacks(model_callback.NewModelFinishTraceCallback(ctx))
}
