// We will use the http package from Go's standard library
package server

import (
	"net/http"
	"net/http/httptest"
)

func TestGetNews(addr string, handler Handler) error {
	t.Run("Get News from Source", func(t *T.testing) {
		request := http.NewRequest(http.MethodGet, "/news", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
