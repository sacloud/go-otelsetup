// Copyright 2022-2023 The sacloud/go-otelsetup Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otelsetup

import (
	"context"
	"os"
	"testing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func TestContextForTrace(t *testing.T) {
	if !Enabled() {
		t.Skip()
	}
	Init(context.Background(), "test", "0.0.0-dev") //nolint

	// 環境変数を仕込む
	// ref: https://www.w3.org/TR/trace-context/#traceparent-header-field-values
	os.Setenv("traceparent", "00-00000000001111111111222222222233-0000000000111111-00")

	traceState := "foo=bar"
	os.Setenv("tracestate", traceState)

	ctx, span := otel.Tracer("test").Start(context.Background(), "span2")
	defer span.End()

	got := ContextForTrace(ctx)

	// ctxからTextMapCarrierを抽出
	envCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(got, envCarrier)

	if len(envCarrier.Keys()) != 2 {
		t.Fatal("got unexpected elements", envCarrier.Keys())
	}

	if envCarrier.Get("traceparent") == "" {
		t.Fatal("got unexpected traceparent", envCarrier.Get("traceparent"))
	}
	if envCarrier.Get("tracestate") != traceState {
		t.Fatal("got unexpected tracestate", envCarrier.Get("tracestate"))
	}
}

func TestExtractTextMapCarrier(t *testing.T) {
	if !Enabled() {
		t.Skip()
	}
	Init(context.Background(), "test", "0.0.0-dev") //nolint

	ctx, span := otel.Tracer("test").Start(context.Background(), "span2")
	defer span.End()

	spanCtx := trace.SpanContextFromContext(ctx)
	traceState, err := spanCtx.TraceState().Insert("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}
	ctx = trace.ContextWithSpanContext(ctx, spanCtx.WithTraceState(traceState))

	carrier := ExtractTextMapCarrier(ctx)
	if len(carrier.Keys()) != 2 {
		t.Fatal("got unexpected keys", carrier)
	}

	if carrier.Get("traceparent") == "" {
		t.Fatal("got unexpected traceparent", carrier.Get("traceparent"))
	}
	if carrier.Get("tracestate") != "foo=bar" {
		t.Fatal("got unexpected tracestate", carrier.Get("tracestate"))
	}
}
