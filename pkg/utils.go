package pkg

import "context"

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value("request_id").(string); ok {
		return v
	}
	return ""
}

func GetCorrelationID(ctx context.Context) string {
	if v, ok := ctx.Value("correlation_id").(string); ok {
		return v
	}
	return ""
}

func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.Value("trace_id").(string); ok {
		return v
	}
	return ""
}
