# go-otelsetup

opentelemetry-goを用いてOTel SDKのセットアップを行うためのライブラリ

## 利用方法

インストール:

    go get github.com/sacloud/go-otelsetup

環境変数`OTEL_EXPORTER_OTLP_ENDPOINT`でOTLPエンドポイントを指定することでトレース/メトリクスが有効になります。

例: `OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317`

また、`OTEL_SDK_DISABLED`に空文字以外の値を指定することでトレース/メトリクスを無効化できます。

## 利用例

```go
	// SDKの初期化
	shutdown, err := otelsetup.Init(context.Background(), "go-otelsetup", "0.0.1")
	if err != nil {
		panic(err)
	}
	defer shutdown(context.Background())

	// トレースの開始
	tracer := otel.Tracer("github.com/sacloud/go-otelsetup")
	ctx, span := tracer.Start(context.Background(), "example")
	defer span.End()

	fmt.Println("SpanID:", trace.SpanContextFromContext(ctx).SpanID())
```