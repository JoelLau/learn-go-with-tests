package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SubPlayerStore struct {
	scores map[string]int
}

func (s *SubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{
		store: &SubPlayerStore{
			scores: map[string]int{
				"Pepper": 20,
				"Floyd":  10,
			},
		},
	}

	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

}

func newGetScoreRequest(name string) *http.Request {
	route := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodGet, route, nil)

	return req
}

func assertStatus(t testing.TB, have, want int) {
	t.Helper()

	if have != want {
		t.Errorf("did not get correct status, have %d, want %d", have, want)
	}
}

func assertResponseBody(t testing.TB, have, want string) {
	t.Helper()

	if have != want {
		t.Errorf("response body does not match, have %q want %q", have, want)
	}
}
