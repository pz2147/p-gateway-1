package util

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

var Traceparent = "traceparent"

type SpanUtil struct {}

// GetTraceParentStr 获取trace串
func (u SpanUtil) GetTraceParentStr(ctx context.Context) string {

	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {

		traceId := spanCtx.TraceID().String()
		spanId := spanCtx.SpanID().String()
		flags := spanCtx.TraceFlags().String()

		f := fmt.Sprintf("%s-%s-%s-%s",flags, traceId, spanId, flags)
		return f
	}
	return ""
}
