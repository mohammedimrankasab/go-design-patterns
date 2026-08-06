// Package strategy demonstrates the Strategy design pattern using a
// production-inspired metadata ingestion platform.
package strategy

import (
	"context"
	"errors"
	"fmt"
)

// ErrNilConnector indicates that no connector was provided.
var ErrNilConnector = errors.New("connector cannot be nil")

// Metadata represents metadata retrieved from an external system.
type Metadata struct {
	ID   string
	Name string
	Type string
}

// Connector defines the strategy for fetching metadata.
type Connector interface {
	// Name returns the connector name.
	Name() string

	// Fetch retrieves metadata from an external system.
	Fetch(ctx context.Context) ([]Metadata, error)
}

// Processor orchestrates metadata ingestion.
type Processor struct {
	connector Connector
}

// NewProcessor creates a Processor using the supplied connector.
func NewProcessor(connector Connector) (*Processor, error) {
	if connector == nil {
		return nil, ErrNilConnector
	}

	return &Processor{
		connector: connector,
	}, nil
}

// Process retrieves metadata using the configured strategy.
func (p *Processor) Process(ctx context.Context) ([]Metadata, error) {
	metadata, err := p.connector.Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"fetch metadata using %s connector: %w",
			p.connector.Name(),
			err,
		)
	}

	return metadata, nil
}

//
// Power BI
//

// PowerBIConnector implements metadata retrieval from Power BI.
type PowerBIConnector struct{}

func (PowerBIConnector) Name() string {
	return "Power BI"
}

func (PowerBIConnector) Fetch(context.Context) ([]Metadata, error) {
	return []Metadata{
		{
			ID:   "dataset-001",
			Name: "Sales Dashboard",
			Type: "Dataset",
		},
		{
			ID:   "report-001",
			Name: "Executive Report",
			Type: "Report",
		},
	}, nil
}

//
// Tableau
//

// TableauConnector implements metadata retrieval from Tableau.
type TableauConnector struct{}

func (TableauConnector) Name() string {
	return "Tableau"
}

func (TableauConnector) Fetch(context.Context) ([]Metadata, error) {
	return []Metadata{
		{
			ID:   "workbook-001",
			Name: "Finance Workbook",
			Type: "Workbook",
		},
		{
			ID:   "datasource-001",
			Name: "Snowflake",
			Type: "Datasource",
		},
	}, nil
}

//
// MLflow
//

// MLflowConnector implements metadata retrieval from MLflow.
type MLflowConnector struct{}

func (MLflowConnector) Name() string {
	return "MLflow"
}

func (MLflowConnector) Fetch(context.Context) ([]Metadata, error) {
	return []Metadata{
		{
			ID:   "model-001",
			Name: "Customer Churn",
			Type: "Model",
		},
		{
			ID:   "experiment-001",
			Name: "Hyperparameter Tuning",
			Type: "Experiment",
		},
	}, nil
}
