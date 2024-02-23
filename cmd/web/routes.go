package main

import (
	"net/http"

	"sqlc-tutorial/assets"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handle("GET /static/*filepath", fileServer)

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /author/create", app.addAuthor)
	mux.HandleFunc("POST /authors/create", app.api.Create)

	return app.recoverPanic(app.securityHeaders(mux))
}
