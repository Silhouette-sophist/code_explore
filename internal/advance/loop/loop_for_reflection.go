package loop

import (
	"context"
	"log"

	"github.com/cloudwego/eino/adk"
)

func LoopForReflection(ctx context.Context, query string) ([]*adk.AgentEvent, error) {
	a, err := adk.NewLoopAgent(ctx, &adk.LoopAgentConfig{
		Name:          "reflection_agent",
		Description:   "Reflection agent with main and critique agent for iterative task solving.",
		SubAgents:     []adk.Agent{NewMainAgent(ctx), NewCritiqueAgent(ctx)},
		MaxIterations: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		EnableStreaming: true, // you can disable streaming here
		Agent:           a,
	})
	iter := runner.Query(ctx, query)
	agentEvents := make([]*adk.AgentEvent, 0)
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		agentEvents = append(agentEvents, event)
	}
	return agentEvents, nil
}
