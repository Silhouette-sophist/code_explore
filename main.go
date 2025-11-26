package main

import (
	"code_explore/internal"
	"context"
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	// 从环境变量中获取您的API KEY，配置方法见：https://www.volcengine.com/docs/82379/1399008
	client := arkruntime.NewClientWithApiKey(internal.ArkApiKey)
	ctx := context.Background()
	req := model.CreateChatCompletionRequest{
		Model: "deepseek-v3-1-terminus",
		Messages: []*model.ChatCompletionMessage{
			&model.ChatCompletionMessage{
				Role: "user",
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("You are a helpful assistant."),
				},
			},
		},
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("standard chat error: %v\n", err)
		return
	}
	fmt.Println(*resp.Choices[0].Message.Content.StringValue)
}
