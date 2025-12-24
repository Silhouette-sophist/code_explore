package indexer

import (
	"code_explore/internal/base/embed"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/indexer/es8" // 导入 es8 索引器
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
)

const (
	indexName          = "eino_example"
	fieldContent       = "content"
	fieldContentVector = "content_vector"
	fieldExtraLocation = "location"
	docExtraLocation   = "location"
)

func NewEsIndexer(ctx context.Context) (*es8.Indexer, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		log.Panicf("connect es8 failed, err=%v", err)
	}

	// 创建 embedding 组件
	arcEmbed, err := embed.NewArcEmbed(ctx)
	if err != nil {
		log.Panicf("create arcEmbed failed, err=%v", err)
	}

	// 创建 es 索引器组件
	return es8.NewIndexer(ctx, &es8.IndexerConfig{
		Client:    client,
		Index:     indexName,
		BatchSize: 10,
		DocumentToFields: func(ctx context.Context, doc *schema.Document) (field2Value map[string]es8.FieldValue, err error) {
			return map[string]es8.FieldValue{
				fieldContent: {
					Value:    doc.Content,
					EmbedKey: fieldContentVector, // 对文档内容进行向量化并保存向量到 "content_vector" 字段
				},
				fieldExtraLocation: {
					Value: doc.MetaData[docExtraLocation],
				},
			}, nil
		},
		Embedding: arcEmbed,
	})
}
