package chat_model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/joho/godotenv"
)

const (
	ApiKey  = "arkApiKey"
	ModelId = "onlineModelId"
)

// NewChatModel 创建chatModel实例，当前均有ark提供模型服务
func NewChatModel(ctx context.Context) (*ark.ChatModel, error) {
	modelConfig := &ark.ChatModelConfig{
		APIKey: GetEnv(ApiKey, ""),
		Model:  GetEnv(ModelId, ""),
	}
	return ark.NewChatModel(ctx, modelConfig)
}

func init() {
	// 获取项目根目录（根据实际情况调整）
	projectRoot, err := getProjectRoot()
	if err != nil {
		panic(err)
	}
	envPath := filepath.Join(projectRoot, ".env")
	// 检查 .env 文件是否存在
	if _, err := os.Stat(envPath); err == nil {
		// 加载 .env 文件
		if err := godotenv.Load(envPath); err != nil {
			panic(err)
		}
		fmt.Println("成功加载 .env 文件")
	} else if os.IsNotExist(err) {
		fmt.Println(".env 文件不存在，使用系统环境变量")
	}
}

// GetEnv 获取环境变量值，支持默认值
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 获取项目根目录（根据实际项目结构调整）
func getProjectRoot() (string, error) {
	// 方式1: 从当前文件向上查找
	// 假设此文件在 projectRoot/config/ 目录下
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// 向上查找直到找到 go.mod 文件
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break // 到达根目录
		}
		dir = parentDir
	}
	return dir, nil
}
