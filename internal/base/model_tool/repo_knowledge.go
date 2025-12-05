package model_tool

import (
	"context"
	"encoding/json"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type RepoKnowledge struct {
}

type RepoKnowledgeParam struct {
	GitRepo string `json:"git_repo"`
}

func (lt *RepoKnowledge) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "view_file",
		Desc: "read file content",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"git_repo": {
				Desc:     "repository url, you execute git remote -v to obtain",
				Type:     schema.String,
				Required: true,
			},
		}),
	}, nil
}

func (lt *RepoKnowledge) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var repoKnowledgeParam RepoKnowledgeParam
	if err := json.Unmarshal([]byte(argumentsInJSON), &repoKnowledgeParam); err != nil {
		logger.CtxInfof(ctx, "unmarshal argumentsInJSON fail: %v", err)
		return "", err
	}
	logger.CtxInfof(ctx, "repoKnowledgeParam: %v", repoKnowledgeParam)
	const exploreMockMsg = "code_explore是一个llm探索工程，从基础到应用层搭建起来"
	return exploreMockMsg, nil
}
