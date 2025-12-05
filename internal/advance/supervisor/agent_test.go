package supervisor

import (
	"code_explore/internal/advance/supervisor/agents"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudwego/eino/adk"
)

func TestSupervisor(t *testing.T) {
	ctx := context.Background()
	supervisor, err := agents.CreateSupervisor(ctx)
	if err != nil {
		t.Fatal(err)
	}
	query := "帮我查找一下supervisor adk的逻辑"
	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent:           supervisor,
		EnableStreaming: true,
	})
	iter := runner.Query(ctx, query)
	events := make([]*adk.AgentEvent, 0)
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}
		events = append(events, next)
	}
	marshal, err := json.Marshal(events)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshal))
}
