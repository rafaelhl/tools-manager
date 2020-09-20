package toolslist_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaehl/tools-manager/tools"
	"github.com/rafaehl/tools-manager/tools/handler/toolslist"
	"github.com/rafaehl/tools-manager/tools/handler/toolslist/mocks"
)

var allTools = []tools.Tool{
	{
		ID:          1,
		Title:       "Notion",
		Link:        "https://notion.so",
		Description: "All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized. ",
		Tags: tools.NewTags(
			"organization",
			"planning",
			"collaboration",
			"writing",
			"calendar",
		),
	},
	{
		ID:          2,
		Title:       "json-server",
		Link:        "https://github.com/typicode/json-server",
		Description: "Fake REST API based on a json schema. Useful for mocking and creating APIs for front-end devs to consume in coding challenges.",
		Tags: tools.NewTags(
			"api",
			"json",
			"schema",
			"node",
			"github",
			"rest",
		),
	},
	{
		ID:          3,
		Title:       "fastify",
		Link:        "https://www.fastify.io/",
		Description: "Extremely fast and simple, low-overhead web framework for NodeJS. Supports HTTP2.",
		Tags: tools.NewTags(
			"web",
			"framework",
			"node",
			"http2",
			"https",
			"localhost",
		),
	},
}

func TestHandler_ServeHTTP_InternalError(t *testing.T) {
	expected := "internal error"
	err := errors.New(expected)
	fetcher := new(mocks.ToolsFetcher)
	fetcher.On("FetchTools", mock.Anything, "").Return(nil, err)

	response := doGet(fetcher, "/tools")

	fetcher.AssertExpectations(t)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, expected, response.Body.String())
}

func TestHandler_ServeHTTP_AllTools(t *testing.T) {
	expected, _ := ioutil.ReadFile("mocks/all_tools.json")
	fetcher := new(mocks.ToolsFetcher)
	fetcher.On("FetchTools", mock.Anything, "").Return(allTools, nil)

	response := doGet(fetcher, "/tools")

	fetcher.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.JSONEq(t, string(expected), response.Body.String())
}

func TestHandler_ServeHTTP_FilteredTools(t *testing.T) {
	expected, _ := ioutil.ReadFile("mocks/filtered_tools.json")
	fetcher := new(mocks.ToolsFetcher)
	fetcher.
		On("FetchTools", mock.Anything, "node").
		Return(allTools[1:], nil)

	response := doGet(fetcher, "/tools?tag=node")

	fetcher.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.JSONEq(t, string(expected), response.Body.String())
}

func doGet(fetcher *mocks.ToolsFetcher, target string) *httptest.ResponseRecorder {
	handler := toolslist.NewHandler(fetcher)
	mux := chi.NewMux()
	mux.Method(handler.Method(), handler.Pattern(), handler)

	request := httptest.NewRequest(http.MethodGet, target, nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	return recorder
}
