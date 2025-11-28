package model_tool

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type ViewFile struct {
}

type ViewFileParam struct {
	FilePath  string `json:"file_path"`
	StartLine int    `json:"start_line"`
	ReadLines int    `json:"read_lines"`
}

type ViewFileResult struct {
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Content   string `json:"content"`
}

func (lt *ViewFile) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "view_file",
		Desc: "read file content",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"file_path": {
				Desc:     "file path",
				Type:     schema.String,
				Required: true,
			},
			"start_line": {
				Desc:     "start line",
				Type:     schema.Integer,
				Required: false,
			},
			"read_lines": {
				Desc:     "read lines",
				Type:     schema.Integer,
				Required: false,
			},
		}),
	}, nil
}

func (lt *ViewFile) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var viewFileParam ViewFileParam
	if err := json.Unmarshal([]byte(argumentsInJSON), &viewFileParam); err != nil {
		logger.CtxInfof(ctx, "unmarshal argumentsInJSON fail: %v", err)
		return "", err
	}
	logger.CtxInfof(ctx, "viewFileParam: %v", viewFileParam)
	stat, err := os.Stat(viewFileParam.FilePath)
	if err != nil {
		logger.CtxInfof(ctx, "stat fail: %v", err)
		return "", err
	}
	if stat.IsDir() {
		logger.CtxInfof(ctx, "file path %s is dir", viewFileParam.FilePath)
		return "", errors.New("not a directory")
	}
	fileBytes, err := os.ReadFile(viewFileParam.FilePath)
	if err != nil {
		logger.CtxInfof(ctx, "read file fail: %v", err)
		return "", err
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	startLine := viewFileParam.StartLine
	endLine := startLine + viewFileParam.ReadLines
	if startLine > len(fileLines) {
		logger.CtxInfof(ctx, "startLine %d > fileLines %d", startLine, len(fileLines))
		return "", errors.New("start line is greater than file line")
	}
	if endLine == startLine {
		endLine = len(fileLines)
	}
	result := &ViewFileResult{
		StartLine: startLine,
		EndLine:   endLine,
	}
	content := make([]string, 0)
	for i, line := range fileLines {
		if i < startLine || i >= endLine {
			continue
		}
		content = append(content, line)
	}
	result.Content = strings.Join(content, "\n")
	marshal, err := json.Marshal(result)
	if err != nil {
		logger.CtxInfof(ctx, "marshal fail: %v", err)
		return "", err
	}
	return string(marshal), nil
}
