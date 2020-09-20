//go:generate mockery -name=ToolsFetcher

package toolslist

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/rafaehl/tools-manager/tools"
)

type (
	ToolsFetcher interface {
		FetchTools(ctx context.Context, tag string) ([]tools.Tool, error)
	}

	handler struct {
		fetcher ToolsFetcher
	}
)

func NewHandler(fetcher ToolsFetcher) handler {
	return handler{
		fetcher: fetcher,
	}
}

func (handler) Method() string {
	return http.MethodGet
}

func (handler) Pattern() string {
	return "/tools"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")

	tools, err := h.fetcher.FetchTools(r.Context(), tag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.Data(w, r, []byte(fmt.Sprintf("%+v", err)))
		return
	}

	render.JSON(w, r, tools)
}
