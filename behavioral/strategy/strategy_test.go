package strategy

import (
	"context"
	"errors"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	t.Run("creates processor", func(t *testing.T) {
		connector := PowerBIConnector{}

		processor, err := NewProcessor(connector)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}

		if processor == nil {
			t.Fatal("expected processor")
		}
	})

	t.Run("nil connector", func(t *testing.T) {
		processor, err := NewProcessor(nil)

		if processor != nil {
			t.Fatal("expected nil processor")
		}

		if !errors.Is(err, ErrNilConnector) {
			t.Fatalf("expected ErrNilConnector, got %v", err)
		}
	})
}

func TestProcessorProcess(t *testing.T) {
	tests := []struct {
		name          string
		connector     Connector
		expectedCount int
		expectedName  string
	}{
		{
			name:          "power bi",
			connector:     PowerBIConnector{},
			expectedCount: 2,
			expectedName:  "Sales Dashboard",
		},
		{
			name:          "tableau",
			connector:     TableauConnector{},
			expectedCount: 2,
			expectedName:  "Finance Workbook",
		},
		{
			name:          "mlflow",
			connector:     MLflowConnector{},
			expectedCount: 2,
			expectedName:  "Customer Churn",
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor, err := NewProcessor(tt.connector)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			metadata, err := processor.Process(ctx)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(metadata) != tt.expectedCount {
				t.Fatalf(
					"expected %d metadata items, got %d",
					tt.expectedCount,
					len(metadata),
				)
			}

			if metadata[0].Name != tt.expectedName {
				t.Fatalf(
					"expected first metadata %q, got %q",
					tt.expectedName,
					metadata[0].Name,
				)
			}
		})
	}
}

func TestProcessorProcessConnectorFailure(t *testing.T) {
	expected := errors.New("connector unavailable")

	connector := failingConnector{
		err: expected,
	}

	processor, err := NewProcessor(connector)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = processor.Process(context.Background())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, expected) {
		t.Fatalf("expected wrapped error")
	}
}

func TestConnectorNames(t *testing.T) {
	tests := []struct {
		name      string
		connector Connector
		expected  string
	}{
		{
			name:      "power bi",
			connector: PowerBIConnector{},
			expected:  "Power BI",
		},
		{
			name:      "tableau",
			connector: TableauConnector{},
			expected:  "Tableau",
		},
		{
			name:      "mlflow",
			connector: MLflowConnector{},
			expected:  "MLflow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.connector.Name(); got != tt.expected {
				t.Fatalf(
					"expected %q, got %q",
					tt.expected,
					got,
				)
			}
		})
	}
}

type failingConnector struct {
	err error
}

func (f failingConnector) Name() string {
	return "Failing Connector"
}

func (f failingConnector) Fetch(context.Context) ([]Metadata, error) {
	return nil, f.err
}
