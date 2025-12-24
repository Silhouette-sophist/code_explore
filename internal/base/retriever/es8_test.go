package retriever

import (
	"context"
	"log"
	"testing"

	"github.com/cloudwego/eino-ext/components/retriever/es8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestEs8Retriever(t *testing.T) {
	ctx := context.Background()
	es8Retriever, err := NewEs8Retriever(ctx)
	if err != nil {
		t.Fatal(err)
	}
	docs, err := es8Retriever.Retrieve(ctx, "useful data")
	if err != nil {
		log.Panicf("retrieve docs failed, err=%v", err)
	}
	for i, doc := range docs {
		t.Logf("doc[%d]: %v", i, doc)
	}

	caseSense := true
	docs, err = es8Retriever.Retrieve(ctx, "useful data",
		es8.WithFilters([]types.Query{{
			Term: map[string]types.TermQuery{
				fieldExtraLocation: {
					CaseInsensitive: &caseSense,
					Value:           "234",
				},
			},
		}}),
	)
	if err != nil {
		log.Panicf("retrieve docs failed, err=%v", err)
	}
	for i, doc := range docs {
		t.Logf("=doc[%d]: %v", i, doc)
	}
}
