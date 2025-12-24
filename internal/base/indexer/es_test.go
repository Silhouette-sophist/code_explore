package indexer

import (
	"context"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestEs8(t *testing.T) {
	ctx := context.Background()
	indexer, err := NewEsIndexer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	ids, err := indexer.Store(ctx, []*schema.Document{
		{ID: "a", Content: "simple data", MetaData: map[string]any{"location": "123"}},
		{ID: "b", Content: "complex data", MetaData: map[string]any{"location": "234"}},
		{ID: "c", Content: "useful data", MetaData: map[string]any{"location": "345"}},
		{ID: "d", Content: "useless data", MetaData: map[string]any{"location": "456"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for i, id := range ids {
		t.Log(i, id)
	}
}
