package agents

import (
	"code_explore/internal/base/chat_model"
	"code_explore/internal/base/model_tool"
	"context"
	_ "embed"
	"fmt"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

//go:embed repo_knowledge.md
var repoKnowledge []byte

func NewRepoKnowledgeAgent(ctx context.Context) (adk.Agent, error) {
	m, err := chat_model.NewChatModel(ctx)
	if err != nil {
		return nil, err
	}
	return adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "repo_knowledge_agent",
		Description: "the agent responsible to repo knowledge query for the repository",
		Instruction: string(repoKnowledge),
		Model:       m,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{
					&model_tool.BashTool{},
				},
				UnknownToolsHandler: func(ctx context.Context, name, input string) (string, error) {
					return fmt.Sprintf("unknown tool: %s", name), nil
				},
			},
		},
	})
}
