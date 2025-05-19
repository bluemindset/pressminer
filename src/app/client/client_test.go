package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Here we want to know if our connection has been created.
*/

const (
	urlPhilenews  = "https://www.philenews.com/kipros/koinonia/article/"
	urlPhilenews2 = "https://www.philenews.com/oikonomia"
)

func TestGetArticles(t *testing.T) {

	t.Run("Get articles from Philenews", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		//PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
