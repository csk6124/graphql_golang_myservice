package main

import (
	"context"
	"flag"
	"fmt"
	"myservice/graphql/todo/gen"
	"myservice/graphql/todo/gen/resolver"
	"net/http"
	"time"

	graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"google.golang.org/grpc/metadata"
)

const servicePort = "9999"

func main() {
	errc := make(chan error)
	ctx := context.Background()

	startHTTPServer(ctx, servicePort, errc)
	fmt.Print("exit", <-errc)
}

// graphql server start
func startHTTPServer(
	ctx context.Context,
	port string,
	errc chan error,
) {
	var (
		httpAddr          = flag.String("http", ":"+port, "HTTP listen address")
		readHeaderTimeout = 60 * time.Second
		writeTimeout      = 60 * time.Second
		idleTimeout       = 60 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
		queryPath         = "/query"
		graphqlPath       = "/graphql"
	)

	// http 세팅
	newResolver := resolver.NewResolver(
		ctx,
	)

	middleware := authMiddleware(ctx, graphqlHandler.NewDefaultServer(
		gen.NewExecutableSchema(gen.Config{Resolvers: &newResolver}),
	))
	mux := http.NewServeMux()
	mux.Handle(queryPath, middleware)
	mux.Handle(graphqlPath, playground.Handler("GraphQL playground", queryPath))
	handler := cors.New(corsOption()).Handler(mux)
	http := &http.Server{
		Addr:              *httpAddr,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	// 서버 스타트
	go func() {
		fmt.Printf(
			"connect to http://localhost%s/graphql for GraphQL playground",
			*httpAddr,
		)
		errc <- http.ListenAndServe()
	}()
}

// auth token
func authMiddleware(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := metadata.Pairs("Authorization", r.Header.Get("Authorization"))
		ctx = metadata.NewIncomingContext(ctx, md)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// corsOption cors option
func corsOption() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{
			"Origin",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"DeviceInfo",
			"Authorization",
			"X-Requested-With",
		},
		AllowCredentials: false,
	}
}
