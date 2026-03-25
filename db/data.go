package db

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const name = "resttest/db"

var tracer = otel.Tracer(name)

type mockdata struct {
	SerialNumber uint   `json:"serial_number"`
	Make         string `json:"make"`
	Model        string `json:"model"`
}

var data = []mockdata{
	{SerialNumber: 1, Make: "Honda", Model: "CB100"},
	{SerialNumber: 2, Make: "Honda", Model: "CB100"},
	{SerialNumber: 3, Make: "Honda", Model: "CB100"},
	{SerialNumber: 4, Make: "Honda", Model: "CB100"},
	{SerialNumber: 5, Make: "Honda", Model: "CB100"},
	{SerialNumber: 6, Make: "Honda", Model: "CB100"},
	{SerialNumber: 7, Make: "Honda", Model: "CB100"},
	{SerialNumber: 8, Make: "Honda", Model: "CB100"},
	{SerialNumber: 9, Make: "Honda", Model: "CB100"},
	{SerialNumber: 10, Make: "Honda", Model: "CB100"},
}

// GetAllItems returns all mock data items with tracing
func GetAllItems(ctx context.Context) []mockdata {
	ctx, span := tracer.Start(ctx, "db.get_all_items",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	recordCount := len(data)
	span.SetAttributes(
		attribute.Int("db.record_count", recordCount),
		attribute.String("db.operation", "GET_ALL"),
		attribute.String("db.table", "mockdata"),
	)

	// Create a copy to avoid external modifications
	result := make([]mockdata, len(data))
	copy(result, data)

	return result
}

// GetItemsCount returns the count of items with tracing
func GetItemsCount(ctx context.Context) int {
	ctx, span := tracer.Start(ctx, "db.get_count",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	count := len(data)
	span.SetAttributes(
		attribute.Int("db.count", count),
		attribute.String("db.operation", "COUNT"),
		attribute.String("db.table", "mockdata"),
	)

	return count
}

// GetItemBySerialNumber returns a single item by serial number with tracing
func GetItemBySerialNumber(ctx context.Context, serialNumber uint) (*mockdata, bool) {
	ctx, span := tracer.Start(ctx, "db.get_by_id",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	span.SetAttributes(
		attribute.Int("db.requested_serial_number", int(serialNumber)),
		attribute.String("db.operation", "GET_BY_ID"),
		attribute.String("db.table", "mockdata"),
	)

	for _, item := range data {
		if item.SerialNumber == serialNumber {
			span.SetAttributes(
				attribute.Bool("db.found", true),
				attribute.String("db.result_make", item.Make),
				attribute.String("db.result_model", item.Model),
			)
			// Return a copy to avoid external modifications
			result := item
			return &result, true
		}
	}

	span.SetAttributes(attribute.Bool("db.found", false))
	return nil, false
}
