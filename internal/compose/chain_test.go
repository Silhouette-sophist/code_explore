package compose

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestNewChatModelWithTool(t *testing.T) {
	ctx := context.Background()
	tool, err := NewChatModelWithTool(ctx)
	if err != nil {
		t.Fatal(err)
	}
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	output, err := tool.Invoke(ctx, []*schema.Message{
		schema.UserMessage(fmt.Sprintf("帮我分析一下%s目录结构", abs)),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
}
