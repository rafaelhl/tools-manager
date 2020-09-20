//go:generate mockery -name=ToolsRemover

package toolsremove

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type (
	ToolsRemover interface {
		RemoveTool(ctx context.Context, id uint) error
	}

	handler struct {
		remover ToolsRemover
	}
)

func NewHandler(remover ToolsRemover) handler {
	return handler{
		remover: remover,
	}
}

func (handler) Method() string {
	return http.MethodDelete
}

func (handler) Pattern() string {
	return "/tools/{id}"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.Data(w, r, []byte(fmt.Sprintf("%+v", err)))
		return
	}

	err = h.remover.RemoveTool(r.Context(), uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.Data(w, r, []byte(fmt.Sprintf("%+v", err)))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
