package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/rafaehl/tools-manager/tools/handler/toolscreate"
	"github.com/rafaehl/tools-manager/tools/handler/toolslist"
	"github.com/rafaehl/tools-manager/tools/handler/toolsremove"
)

type HTTPHandler interface {
	Method() string
	Pattern() string
	http.Handler
}

func main() {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "pong")
	})

	getTools := toolslist.NewHandler(nil)
	router.Method(getTools.Method(), getTools.Pattern(), getTools)

	postTools := toolscreate.NewHandler(nil)
	router.Method(http.MethodPost, "/tools", postTools)

	deleteTool := toolsremove.NewHandler(nil)
	router.Method(http.MethodDelete, "/tools/{id}", deleteTool)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
