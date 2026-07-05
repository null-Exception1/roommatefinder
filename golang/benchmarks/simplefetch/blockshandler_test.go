package simplefetch

import (
	"golang/handlers"
	initfuncs "golang/init"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkBlocksHandler(t *testing.B) {

	godotenv.Load("../../.env")

	initfuncs.Database()

	req := httptest.NewRequest("GET", "/rooms?block=16", nil)

	w := httptest.NewRecorder()

	handlers.Blocks(w, req)

	var resp string
	resp = w.Body.String()

	if resp == "" {
		t.Fatal("Failed")
	}
}
