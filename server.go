package snaparser_server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// RunConfig represents the configuration for running the snaparser server.
type RunConfig struct {
	allowedOrigins, allowedMethods []string
	keyFilePath, certFilePath      string
	endpoint                       string
	limiterType                    limiterType
	port                           int
	useHTTP                        bool
	enableLog                      bool
}

// limiterType represents the type of rate limiter.
type limiterType string

const (
	none    limiterType = "none"
	lenient limiterType = "lenient"
	normal  limiterType = "normal"
	strict  limiterType = "strict"
)

// limitHandler is the handler for exceeding rate limits.
var limitHandler = httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "request limit reached", http.StatusTooManyRequests)
})

// limiters holds the rate limiter functions.
var limiters = map[limiterType]func(next http.Handler) http.Handler{
	none:    nil,
	lenient: httprate.Limit(10, time.Second, httprate.WithKeyFuncs(httprate.KeyByIP), limitHandler),
	normal:  httprate.Limit(10, time.Minute, httprate.WithKeyFuncs(httprate.KeyByIP), limitHandler),
	strict:  httprate.Limit(10, time.Hour, httprate.WithKeyFuncs(httprate.KeyByIP), limitHandler),
}

// run starts the snaparser server.
func run(runConfig *RunConfig) (err error) {
	r := chi.NewRouter()

	lmt, ok := limiters[runConfig.limiterType]
	if !ok {
		return fmt.Errorf("invalid limiter type: %s", runConfig.limiterType)
	}
	if runConfig.limiterType != none {
		r.Use(lmt)
	}

	if runConfig.enableLog {
		r.Use(middleware.Logger)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: runConfig.allowedOrigins,
		AllowedMethods: runConfig.allowedMethods,
	}))

	r.Post(runConfig.endpoint, UploadHandler)

	addr := fmt.Sprintf(":%d", runConfig.port)

	log.Printf("snaparser server listening on port %d\n", runConfig.port)

	if runConfig.useHTTP {
		return http.ListenAndServe(addr, r)
	}

	return http.ListenAndServeTLS(addr, runConfig.certFilePath, runConfig.keyFilePath, r)
}
