package agent

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestNewReactAgent(t *testing.T) {
	ctx := context.Background()
	rootDir, _ := filepath.Abs("./../..")
	agent, err := NewReactAgent(ctx)
	if err != nil {
		t.Fatal(err)
	}
	generate, err := agent.Generate(ctx, []*schema.Message{
		schema.UserMessage(fmt.Sprintf("帮我看看目录%s下是否有go.mod文件", rootDir)),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("generate: %+v", *generate)
}
