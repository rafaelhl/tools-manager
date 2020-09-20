//go:generate mockery -name=ToolsCreator

package toolscreate

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/rafaehl/tools-manager/tools"
)

type (
	ToolsCreator interface {
		CreateTool(ctx context.Context, tool tools.Tool) error
	}

	handler struct {
		creator ToolsCreator
	}
)

func NewHandler(creator ToolsCreator) handler {
	return handler{
		creator: creator,
	}
}

func (handler) Method() string {
	return http.MethodPost
}

func (handler) Pattern() string {
	return "/tools"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var tool tools.Tool
	err := json.NewDecoder(r.Body).Decode(&tool)
	if err != nil || Empty(tool.Title, tool.Link, tool.Tags.String()) {
		w.WriteHeader(http.StatusBadRequest)
		render.Data(w, r, []byte(fmt.Sprintf("%+v", err)))
		return
	}

	err = h.creator.CreateTool(r.Context(), tool)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.Data(w, r, []byte(fmt.Sprintf("%+v", err)))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Empty(values ...string) bool {
	for _, v := range values {
		if v == "" {
			return true
		}
	}

	return false
}
