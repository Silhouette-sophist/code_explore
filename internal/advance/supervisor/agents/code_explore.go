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

//go:embed code_explore.md
var codeExplore []byte

func NewCodeExploreAgent(ctx context.Context) (adk.Agent, error) {
	m, err := chat_model.NewChatModel(ctx)
	if err != nil {
		return nil, err
	}
	return adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "code_explore_agent",
		Description: "the agent responsible to code explore with repository",
		Instruction: string(codeExplore),
		Model:       m,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{
					&model_tool.ListFile{},
					&model_tool.ViewFile{},
				},
				UnknownToolsHandler: func(ctx context.Context, name, input string) (string, error) {
					return fmt.Sprintf("unknown tool: %s", name), nil
				},
			},
		},
	})
}
