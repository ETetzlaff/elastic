package elastic

import (
	"encoding/json"
	"testing"
)

func TestMaxAggregation(t *testing.T) {
	agg := NewMaxAggregation().Field("price")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"max":{"field":"price"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestMaxAggregationWithFormat(t *testing.T) {
	agg := NewMaxAggregation().Field("price").Format("00000.00")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"max":{"field":"price","format":"00000.00"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
