package main

import (
	"context"
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: service,
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if err != nil {
		panic(err)
	}
	return tracer, closer
}

func TestController(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TestController")

	defer func() {
		span.SetTag("req", "127.0.0.1")
		span.Finish()
	}()
	time.Sleep(time.Second)

	_ctx := opentracing.ContextWithSpan(context.Background(), span)
	TestService(_ctx)
}

func TestService(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TestService")
	defer func() {
		span.SetTag("target", "TestService")
		span.Finish()
	}()
	time.Sleep(time.Second)
}

func main() {
	tracer, closer := initJaeger("admin")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	RootSpan := tracer.StartSpan("root")
	ctx := opentracing.ContextWithSpan(context.Background(), RootSpan)

	TestController(ctx)
}

/*
docker run -d --name jaeger \                
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.33
*/