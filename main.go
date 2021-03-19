package main

import (
	"fmt"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/fasthttp/router"
	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func main() {
	config := &statsd.ClientConfig{
		Address: "127.0.0.1:8125",
		Prefix:  "test-client",
	}
	statsd.NewClientWithConfig(config)

	router.New()

	sentry.Init(sentry.ClientOptions{})

	assert.New(nil)

	bytebufferpool.Get()

	fmt.Println(fasthttp.DefaultConcurrency)

	zap.NewExample()

	fmt.Println("hello!")
}
