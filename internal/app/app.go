package app

import (
	"log"
	"net/http"
)

type (
	App struct {
		server     *http.Server
		router     *Router
		middleware []MiddlewareFunc
	}

	MiddlewareFunc func(next http.Handler) http.Handler
)

func ProvideApp(addr string) *App {
	app := &App{
		server: new(http.Server),
		router: NewRouter(),
	}

	app.server.Handler = app
	app.server.Addr = addr

	return app
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := app.router.Find(req.Method, getPath(req))
	handler = applyMiddleware(handler, app.middleware...)

	handler.ServeHTTP(w, req)
	return
}

func (app *App) Use(m ...MiddlewareFunc) {
	app.middleware = append(app.middleware, m...)
}

func (app *App) Get(path string, handler http.Handler) {
	app.router.Add(http.MethodGet, path, handler)
}

func (app *App) Post(path string, handler http.Handler) {
	app.router.Add(http.MethodPost, path, handler)
}

func (app *App) Start() {
	log.Fatal(app.server.ListenAndServe())
}

func (app *App) Handler(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(handler)
}

func applyMiddleware(handler http.Handler, middleware ...MiddlewareFunc) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}

func getPath(req *http.Request) string {
	path := req.URL.RawPath
	if path == "" {
		path = req.URL.Path
	}

	return path
}
