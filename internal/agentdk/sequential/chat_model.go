package sequential

import (
	"code_explore/internal/base/chat_model"
	"context"
	"log"

	"github.com/cloudwego/eino/adk"
)

func NewAgentWithConfig(ctx context.Context, config *adk.ChatModelAgentConfig) (adk.Agent, error) {
	a, err := adk.NewChatModelAgent(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	return a, nil
}

func NewPlanAgent(ctx context.Context) (adk.Agent, error) {
	chatModel, err := chat_model.NewChatModel(ctx)
	if err != nil {
		return nil, err
	}
	agent, err := NewAgentWithConfig(ctx, &adk.ChatModelAgentConfig{
		Name:        "PlannerAgent",
		Description: "Generates a research plan based on a topic.",
		Instruction: `
You are an expert research planner. 
Your goal is to create a comprehensive, step-by-step research plan for a given topic. 
The plan should be logical, clear, and easy to follow.
The user will provide the research topic. Your output must ONLY be the research plan itself, without any conversational text, introductions, or summaries.`,
		Model:     chatModel,
		OutputKey: "Plan",
	})
	if err != nil {
		log.Fatal(err)
	}
	return agent, nil
}

func NewWriterAgent(ctx context.Context) (adk.Agent, error) {
	chatModel, err := chat_model.NewChatModel(ctx)
	if err != nil {
		return nil, err
	}
	agent, err := NewAgentWithConfig(ctx, &adk.ChatModelAgentConfig{
		Name:        "WriterAgent",
		Description: "Writes a report based on a research plan.",
		Instruction: `
You are an expert academic writer.
You will be provided with a detailed research plan:
{Plan}

Your task is to expand on this plan to write a comprehensive, well-structured, and in-depth report.`,
		Model: chatModel,
	})
	if err != nil {
		log.Fatal(err)
	}
	return agent, nil
}
