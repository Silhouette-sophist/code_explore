package model_tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// BashTool 执行bash命令的工具结构体
type BashTool struct {
	// 可配置的命令执行超时时间，默认10秒
	Timeout time.Duration
}

// BashToolParam 工具入参结构体（已移除git_repo）
type BashToolParam struct {
	// 要执行的bash命令（必填）
	Cmd string `json:"cmd"`
}

// Info 返回工具的元信息，描述工具名称、用途和参数规范
func (lt *BashTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "bash_tool",
		Desc: "execute general bash commands",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"cmd": {
				Desc:     "bash command to execute",
				Type:     schema.String,
				Required: true, // 改为必填参数
			},
		}),
	}, nil
}

// InvokableRun 工具核心执行方法，解析参数并执行bash命令
func (lt *BashTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	// 1. 解析入参
	var bashToolParam BashToolParam
	if err := json.Unmarshal([]byte(argumentsInJSON), &bashToolParam); err != nil {
		logger.CtxInfof(ctx, "unmarshal argumentsInJSON fail: %v, input: %s", err, argumentsInJSON)
		return "", fmt.Errorf("参数解析失败: %w", err)
	}
	logger.CtxInfof(ctx, "bashToolParam parsed: %+v", bashToolParam)

	// 2. 参数校验（仅校验cmd参数）
	cmdStr := strings.TrimSpace(bashToolParam.Cmd)
	if cmdStr == "" {
		return "", fmt.Errorf("cmd 参数不能为空")
	}

	// 3. 安全校验：防止命令注入（核心安全措施）
	if err := validateCommand(cmdStr); err != nil {
		logger.CtxErrorf(ctx, "command validation failed: %v, cmd: %s", err, cmdStr)
		return "", fmt.Errorf("命令校验失败（安全限制）: %w", err)
	}

	// 4. 执行bash命令
	logger.CtxInfof(ctx, "executing bash command: %s", cmdStr)
	result, err := lt.executeBashCommand(ctx, cmdStr)
	if err != nil {
		logger.CtxErrorf(ctx, "execute bash command fail: %v, cmd: %s", err, cmdStr)
		return "", fmt.Errorf("命令执行失败: %w", err)
	}

	// 5. 处理执行结果
	if strings.TrimSpace(result) == "" {
		return fmt.Sprintf("命令执行成功，但无输出。执行的命令：%s", cmdStr), nil
	}
	return result, nil
}

// executeBashCommand 执行bash命令，包含超时控制
func (lt *BashTool) executeBashCommand(ctx context.Context, cmdStr string) (string, error) {
	// 设置默认超时时间
	timeout := lt.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	// 创建带超时的上下文
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 执行bash命令（-c 参数用于执行字符串形式的命令）
	cmd := exec.CommandContext(ctxWithTimeout, "/bin/bash", "-c", cmdStr)

	// 捕获标准输出和标准错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 区分超时错误和执行错误
		if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
			return "", fmt.Errorf("命令执行超时（%v）: %w", timeout, err)
		}
		return string(output), fmt.Errorf("命令返回非0状态码: %w, 输出: %s", err, string(output))
	}

	return string(output), nil
}

// validateCommand 命令安全校验，防止命令注入
func validateCommand(cmdStr string) error {
	// 禁止的危险命令（可根据业务扩展）
	forbiddenCmds := []string{"rm -rf", "sudo", "mv", "cp", "chmod", "chown", ">&", "|", ";", "&"}
	for _, forbidden := range forbiddenCmds {
		if strings.Contains(strings.ToLower(cmdStr), forbidden) {
			return fmt.Errorf("命令包含危险操作: %s", forbidden)
		}
	}

	return nil
}
