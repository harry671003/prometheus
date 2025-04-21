package v1

import (
	"context"
	"net/http"
)

type engineKeyType struct{}

var engineKey = engineKeyType{}

const EngineHeader = "X-PromQL-EngineType"

type EngineType string

const (
	Prometheus EngineType = "prometheus"
	Thanos     EngineType = "thanos"
	None       EngineType = "none"
)

func addEngineTypeToContext(ctx context.Context, r *http.Request) context.Context {
	engine := EngineType(r.Header.Get(EngineHeader))
	switch engine {
	case Prometheus, Thanos:
		return context.WithValue(ctx, engineKey, engine)
	default:
		return context.WithValue(ctx, engineKey, None)
	}
}

func GetEngineType(ctx context.Context) EngineType {
	if engine, ok := ctx.Value(engineKey).(EngineType); ok {
		return engine
	}
	return None
}
