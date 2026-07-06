package simplefetch

import (
	"golang/handlers"
	initfuncs "golang/init"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkBlocksHandler(t *testing.B) {

	godotenv.Load("../../.testenv")

	initfuncs.Database()
	var w *httptest.ResponseRecorder
	var req *http.Request
	for i := 0; i <= t.N; i++ {
		req = httptest.NewRequest("GET", "/rooms?block=16", nil)

		w = httptest.NewRecorder()

		handlers.Blocks(w, req)

		result := w.Result()
		if result.Status != "200 OK" {
			t.Fatal("Failed")
		}

	}
}
