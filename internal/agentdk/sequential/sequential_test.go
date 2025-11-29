package sequential

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func TestNewPlanAgent(t *testing.T) {
	ctx := context.Background()
	agent, err := NewPlanAgent(ctx)
	if err != nil {
		t.Fatal(err)
	}
	run := agent.Run(ctx, &adk.AgentInput{
		Messages: []adk.Message{
			schema.UserMessage("帮我规划一下llm学习路线"),
		},
	})
	for {
		next, valid := run.Next()
		if next == nil || !valid {
			break
		}
		message, event, err := adk.GetMessage(next)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(message)
		fmt.Println(event)
	}
}
