// Copyright 2023-2025 The sacloud/go-otelsetup Authors
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

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// ContextForTrace 環境変数からトレースコンテキストを読み取り、propagator経由で値を反映させたコンテキストを返す
func ContextForTrace(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	envCarrier := propagation.MapCarrier{}
	for _, key := range (propagation.TraceContext{}).Fields() {
		envCarrier.Set(key, os.Getenv(key))
	}

	return otel.GetTextMapPropagator().Extract(ctx, envCarrier)
}

// ExtractTextMapCarrier 現在のコンテキストからTextMapCarrierを抽出する
func ExtractTextMapCarrier(ctx context.Context) propagation.TextMapCarrier {
	envCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, envCarrier)
	return envCarrier
}
