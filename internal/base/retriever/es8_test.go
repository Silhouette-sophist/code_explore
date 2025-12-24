package retriever

import (
	"context"
	"log"
	"testing"
)

func TestEs8Retriever(t *testing.T) {
	ctx := context.Background()
	es8Retriever, err := NewEs8Retriever(ctx)
	if err != nil {
		t.Fatal(err)
	}
	docs, err := es8Retriever.Retrieve(ctx, "data")
	if err != nil {
		log.Panicf("retrieve docs failed, err=%v", err)
	}
	log.Println(docs)
}
