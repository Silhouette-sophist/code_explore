package model_tool

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path"

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
		return "", err
	}
	stat, err := os.Stat(listFileParam.Directory)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return "", errors.New("not a directory")
	}
	dirs, err := os.ReadDir(listFileParam.Directory)
	if err != nil {
		return "", err
	}
	items := make([]ListFileItem, len(dirs))
	for _, dir := range dirs {
		items = append(items, ListFileItem{
			Path: path.Join(listFileParam.Directory, dir.Name()),
			Dir:  dir.IsDir(),
			Name: dir.Name(),
		})
	}
	marshal, err := json.Marshal(items)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
