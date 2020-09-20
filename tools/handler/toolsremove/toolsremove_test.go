package toolsremove_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaehl/tools-manager/tools/handler/toolsremove"
	"github.com/rafaehl/tools-manager/tools/handler/toolsremove/mocks"
)

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	remover := new(mocks.ToolsRemover)
	response := doDelete(remover, "/tools/a")

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestHandler_ServeHTTP_InternalError(t *testing.T) {
	remover := new(mocks.ToolsRemover)
	remover.On("RemoveTool", mock.Anything, uint(1)).Return(errors.New("internal error"))
	response := doDelete(remover, "/tools/1")

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestHandler_ServeHTTP_RemoveTool(t *testing.T) {
	remover := new(mocks.ToolsRemover)
	remover.On("RemoveTool", mock.Anything, uint(1)).Return(nil)
	response := doDelete(remover, "/tools/1")

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func doDelete(remover *mocks.ToolsRemover, target string) *httptest.ResponseRecorder {
	handler := toolsremove.NewHandler(remover)
	mux := chi.NewMux()
	mux.Method(handler.Method(), handler.Pattern(), handler)

	request := httptest.NewRequest(http.MethodDelete, target, nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	return recorder
}
