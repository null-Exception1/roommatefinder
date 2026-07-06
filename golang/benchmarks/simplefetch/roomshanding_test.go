package simplefetch

import (
	"golang/handlers"
	initfuncs "golang/init"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkRoomsBlocksHandler(t *testing.B) {

	godotenv.Load("../../.testenv")
	initfuncs.Database()

	for i := 0; i <= 20; i++ {

		req := httptest.NewRequest("GET", "/rooms?block="+strconv.Itoa(i), nil)

		w := httptest.NewRecorder()

		handlers.Rooms(w, req)

		var resp string
		resp = w.Body.String()

		if resp == "" {
			t.Fatal("Failed")
		}
	}
}
