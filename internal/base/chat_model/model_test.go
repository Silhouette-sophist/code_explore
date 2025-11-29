package chat_model

import (
	"code_explore/internal/base/model_callback"
	"context"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestModel(t *testing.T) {
	ctx := context.Background()
	model_callback.InitAppendGlobalHandlers(ctx, model_callback.NewModelFinishTraceCallback(ctx))
	model, err := NewChatModel(ctx, DoubaoThinking)
	if err != nil {
		t.Fatal(err)
	}
	msg, err := model.Generate(ctx, []*schema.Message{
		schema.UserMessage("eino框架是什么？"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msg)
}

func TestModelWithCallback(t *testing.T) {
	ctx := context.Background()
	model_callback.InitAppendGlobalHandlers(ctx, model_callback.NewChatModelCallback(ctx))
	model, err := NewChatModel(ctx, DoubaoThinking)
	if err != nil {
		t.Fatal(err)
	}
	msg, err := model.Generate(ctx, []*schema.Message{
		schema.UserMessage("eino框架是什么？"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msg)
}

func TestModelWithInitCallback(t *testing.T) {
	ctx := context.Background()
	model, err := NewChatModel(ctx, DefaultModel)
	if err != nil {
		t.Fatal(err)
	}
	ctx = model_callback.ContextWrapHandlers(ctx, model_callback.NewChatModelCallback(ctx))
	msg, err := model.Generate(ctx, []*schema.Message{
		schema.UserMessage("eino框架是什么？"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msg)
}
