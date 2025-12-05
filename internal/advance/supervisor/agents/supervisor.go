package agents

import (
	"code_explore/internal/base/chat_model"
	"context"
	_ "embed"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/prebuilt/supervisor"
)

//go:embed supervisor.md
var supervisorInfo []byte

func CreateSupervisor(ctx context.Context) (adk.Agent, error) {
	m, err := chat_model.NewChatModel(ctx)
	if err != nil {
		return nil, err
	}

	sv, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "supervisor",
		Description: "the agent responsible to supervise tasks",
		Instruction: string(supervisorInfo),
		Model:       m,
		Exit:        &adk.ExitTool{},
	})
	if err != nil {
		return nil, err
	}

	codeExploreAgent, err := NewCodeExploreAgent(ctx)
	if err != nil {
		return nil, err
	}
	repoKnowledgeAgent, err := NewRepoKnowledgeAgent(ctx)
	if err != nil {
		return nil, err
	}

	return supervisor.New(ctx, &supervisor.Config{
		Supervisor: sv,
		SubAgents:  []adk.Agent{codeExploreAgent, repoKnowledgeAgent},
	})
}
