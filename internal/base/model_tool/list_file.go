package model_tool

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type ListFile struct {
}

type ListFileParam struct {
	Directory string `json:"directory"`
}

type ListFileResult []ListFileItem

type ListFileItem struct {
	Path string `json:"path"`
	Dir  bool   `json:"dir"`
	Name string `json:"name"`
}

func (lt *ListFile) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "list_file",
		Desc: "list file in directory",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"directory": {
				Desc:     "directory path",
				Type:     schema.String,
				Required: true,
			},
		}),
	}, nil
}

func (lt *ListFile) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var listFileParam ListFileParam
	if err := json.Unmarshal([]byte(argumentsInJSON), &listFileParam); err != nil {
		logger.CtxInfof(ctx, "unmarshal argumentsInJSON fail: %v", err)
		return "", err
	}
	logger.CtxInfof(ctx, "listFileParam: %v", listFileParam)
	stat, err := os.Stat(listFileParam.Directory)
	if err != nil {
		logger.CtxInfof(ctx, "stat fail: %v", err)
		return "", err
	}
	if !stat.IsDir() {
		logger.CtxInfof(ctx, "not dir: %v", listFileParam.Directory)
		return "", errors.New("not a directory")
	}
	if strings.HasPrefix(stat.Name(), ".") {
		logger.CtxInfof(ctx, "hidden dir: %v", listFileParam.Directory)
		return "", errors.New("hidden dir")
	}
	dirs, err := os.ReadDir(listFileParam.Directory)
	if err != nil {
		logger.CtxInfof(ctx, "read dir fail: %v", err)
		return "", err
	}
	logger.CtxInfof(ctx, "read dir: %v", listFileParam.Directory)
	items := make([]ListFileItem, 0, len(dirs))
	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), ".") {
			continue
		}
		items = append(items, ListFileItem{
			Path: path.Join(listFileParam.Directory, dir.Name()),
			Dir:  dir.IsDir(),
			Name: dir.Name(),
		})
	}
	marshal, err := json.Marshal(items)
	if err != nil {
		logger.CtxInfof(ctx, "marshal fail: %v", err)
		return "", err
	}
	return string(marshal), nil
}
