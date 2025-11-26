package compose

import (
	"code_explore/internal/chat_model"
	"code_explore/internal/model_tool"
	"context"
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func NewChatModelWithTool(ctx context.Context) (compose.Runnable[[]*schema.Message, []*schema.Message], error) {
	// 1.初始化 tools
	todoTools := []tool.BaseTool{
		&model_tool.ListFile{},
	}

	// 2.创建并配置 ChatModel
	chatModel, err := chat_model.NewChatModel(ctx, chat_model.DoubaoThinking)
	if err != nil {
		log.Fatal(err)
	}

	// 3.获取工具信息并绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	for _, tool := range todoTools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolInfos = append(toolInfos, info)
	}
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		log.Fatal(err)
	}

	// 4.编排调用链
	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		log.Fatal(err)
	}
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	// 5.编译并运行 chain
	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return agent, nil
}
