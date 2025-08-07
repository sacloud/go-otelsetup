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

package main

import (
	"context"
	"fmt"

	otelsetup "github.com/sacloud/go-otelsetup"
	"github.com/sacloud/go-otelsetup/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// 1) あらかじめOpenTelemetry Collectorを起動し、OTLPでトレース/メトリクスを受け取れるようにしておく
	//
	// $ docker compose up -d
	//
	// 2) OTLPエンドポイントを環境変数で指定した上で起動
	//
	// $ OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317 go run .
	//
	// 3) Jaegerでトレース出力を確認
	//
	//  $ open http://localhost:16686

	shutdown, err := otelsetup.InitWithOptions(context.Background(),
		otelsetup.Options{
			ServiceName:      "otelsetup",
			ServiceVersion:   version.Version,
			ServiceNamespace: "sacloud",
		})
	if err != nil {
		panic(err)
	}
	defer shutdown(context.Background()) //nolint:errcheck

	tracer := otel.Tracer("go-otelsetup")
	ctx, span := tracer.Start(context.Background(), "example")
	defer span.End()

	fmt.Println("SpanID:", trace.SpanContextFromContext(ctx).SpanID())
}
