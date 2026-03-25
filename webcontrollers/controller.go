package webcontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resttest/db"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func HandleHealthCheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Serivce is healthy")
}

func HandleGetItems(res http.ResponseWriter, req *http.Request) {
	ctx, span := tracer.Start(req.Context(), "controller.get_items",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	items := db.GetAllItems(ctx)
	span.SetAttributes(attribute.Int("controller.items_retrieved", len(items)))

	buffer, err := json.Marshal(items)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(attribute.Bool("controller.error", true))
		fmt.Printf("error marshaling items: %v\n", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Write(buffer)
}
