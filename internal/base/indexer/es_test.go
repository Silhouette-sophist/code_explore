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
		{ID: "a", Content: "simple data"},
		{ID: "b", Content: "complex data"},
		{ID: "c", Content: "useful data"},
		{ID: "d", Content: "useless data"},
	})
	if err != nil {
		t.Fatal(err)
	}
	for i, id := range ids {
		t.Log(i, id)
	}
}
