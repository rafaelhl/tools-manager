package toolscreate_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaehl/tools-manager/tools"
	"github.com/rafaehl/tools-manager/tools/handler/toolscreate"
	"github.com/rafaehl/tools-manager/tools/handler/toolscreate/mocks"
)

var expected = tools.Tool{
	Title:       "hotel",
	Link:        "https://github.com/typicode/hotel",
	Description: "Local app manager. Start apps within your browser, developer tool with local .localhost domain and https out of the box.",
	Tags:        tools.NewTags("node", "organizing", "webapps", "domain", "developer", "https", "proxy"),
}

func TestHandler_ServeHTTP_BadRequest(t *testing.T) {
	cases := []struct {
		title string
		body  string
	}{
		{title: "invalid json", body: "{"},
		{title: "invalid schema", body: `{"tool": "go"}`},
		{title: "empty title", body: `{"description": "Teste", "tags": "go"}`},
		{title: "empty description", body: `{"title": "Teste", "tags": "go"}`},
		{title: "empty tags", body: `{"title": "Teste", "description": "Teste"}`},
	}

	creator := new(mocks.ToolsCreator)
	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			response := doPost(creator, "/tools", strings.NewReader(test.body))
			creator.AssertExpectations(t)
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	}
}

func TestHandler_ServeHTTP_CreateTool(t *testing.T) {
	body, _ := ioutil.ReadFile("mocks/request_body.json")
	creator := new(mocks.ToolsCreator)
	creator.On("CreateTool", mock.Anything, expected).Return(nil)

	response := doPost(creator, "/tools", strings.NewReader(string(body)))

	creator.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func doPost(creator *mocks.ToolsCreator, target string, body io.Reader) *httptest.ResponseRecorder {
	handler := toolscreate.NewHandler(creator)
	mux := chi.NewMux()
	mux.Method(handler.Method(), handler.Pattern(), handler)

	request := httptest.NewRequest(http.MethodPost, target, body)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	return recorder
}
