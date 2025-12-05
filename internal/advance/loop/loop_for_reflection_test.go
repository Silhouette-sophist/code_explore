package loop

import (
	"context"
	"encoding/json"
	"testing"
)

func TestLoopForReflection(t *testing.T) {
	ctx := context.Background()
	query := "briefly introduce what a multimodal embedding model is."
	reflection, err := LoopForReflection(ctx, query)
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := json.Marshal(reflection)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
