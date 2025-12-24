package retriever

import (
	"code_explore/internal/base/embed"
	"context"
	"encoding/json"
	"log"

	"github.com/cloudwego/eino-ext/components/retriever/es8"
	"github.com/cloudwego/eino-ext/components/retriever/es8/search_mode"
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexName          = "eino_example"
	fieldContent       = "content"
	fieldContentVector = "content_vector"
	fieldExtraLocation = "location"
	docExtraLocation   = "location"
)

func NewEs8Retriever(ctx context.Context) (*es8.Retriever, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		log.Panicf("connect es8 failed, err=%v", err)
	}

	arcEmbed, err := embed.NewArcEmbed(ctx)
	if err != nil {
		log.Panicf("create arcEmbed failed, err=%v", err)
	}

	// 创建检索器组件
	return es8.NewRetriever(ctx, &es8.RetrieverConfig{
		Client: client,
		Index:  indexName,
		TopK:   5,
		SearchMode: search_mode.SearchModeApproximate(&search_mode.ApproximateConfig{
			QueryFieldName:  fieldContent,
			VectorFieldName: fieldContentVector,
			Hybrid:          true,
			// RRF 仅在特定许可证下可用
			// 参见：https://www.elastic.co/subscriptions
			RRF:             false,
			RRFRankConstant: nil,
			RRFWindowSize:   nil,
		}),
		ResultParser: func(ctx context.Context, hit types.Hit) (doc *schema.Document, err error) {
			doc = &schema.Document{
				ID:       *hit.Id_,
				Content:  "",
				MetaData: map[string]any{},
			}

			var src map[string]any
			if err = json.Unmarshal(hit.Source_, &src); err != nil {
				return nil, err
			}

			for field, val := range src {
				switch field {
				case fieldContent:
					doc.Content = val.(string)
				case fieldContentVector:
					var v []float64
					for _, item := range val.([]interface{}) {
						v = append(v, item.(float64))
					}
					doc.WithDenseVector(v)
				case fieldExtraLocation:
					doc.MetaData[docExtraLocation] = val.(string)
				}
			}

			if hit.Score_ != nil {
				doc.WithScore(float64(*hit.Score_))
			}

			return doc, nil
		},
		Embedding: arcEmbed, // 你的 embedding 组件
	})
}
