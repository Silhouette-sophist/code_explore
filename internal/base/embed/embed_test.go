package embed

import (
	"context"
	"testing"
)

func TestEmbed(t *testing.T) {
	ctx := context.Background()
	embed, err := NewArcEmbed(ctx)
	if err != nil {
		t.Fatal(err)
	}
	embeddings, err := embed.EmbedStrings(ctx, []string{
		"趋之若鹜",
		"交头接耳",
	})
	if err != nil {
		t.Fatal(err)
	}
	for i, embedding := range embeddings {
		t.Logf("embedding %d %d %v", i, len(embedding), embedding)
	}
}
