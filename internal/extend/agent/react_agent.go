package agent

import (
	"code_explore/internal/base/chat_model"
	"code_explore/internal/base/model_tool"
	"context"
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func NewReactAgent(ctx context.Context) (*react.Agent, error) {
	// 1.chatModel创建
	chatModel, err := chat_model.NewChatModel(ctx, chat_model.DoubaoThinking)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 2.初始化所需的 tools
	tools := compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{
			&model_tool.ListFile{},
			&model_tool.ViewFile{},
		},
	}
	// 3初始化agent
	agent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: chatModel,
		ToolsConfig:      tools,
		MaxStep:          200,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return agent, nil
}
