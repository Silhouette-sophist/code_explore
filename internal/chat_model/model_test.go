package chat_model

import (
	"context"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestModel(t *testing.T) {
	ctx := context.Background()
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
