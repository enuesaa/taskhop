package gql

import (
	"net/http"
	"time"

	"github.com/enuesaa/taskhop/internal/usecase"
	"github.com/gorilla/websocket"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func New(ucase usecase.Usecase) http.HandlerFunc {
	// see https://github.com/99designs/gqlgen/issues/1664
	gqhandle := handler.New(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Usecase: ucase,
		},
	}))
	gqhandle.AddTransport(transport.Options{})
	gqhandle.AddTransport(transport.GET{})
	gqhandle.AddTransport(transport.POST{})
	gqhandle.AddTransport(transport.MultipartForm{})
	gqhandle.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	gqhandle.Use(extension.Introspection{})

	return gqhandle.ServeHTTP
}
