package simplefetch

import (
	"golang/handlers"
	initfuncs "golang/init"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkBlocksHandler(t *testing.B) {

	godotenv.Load("../../.testenv")
	//fmt.Println("CACHING: ", os.Getenv("CACHING"))
	initfuncs.Logging()
	initfuncs.Database()

	t.ResetTimer()
	var w *httptest.ResponseRecorder
	var req *http.Request
	for j := 0; j < t.N; j++ {
		for i := 0; i <= 20; i++ {
			req = httptest.NewRequest("GET", "/rooms?block="+strconv.Itoa(i), nil)

			w = httptest.NewRecorder()

			handlers.Blocks(w, req)

			result := w.Result()
			if result.Status != "200 OK" {
				t.Fatal("Failed")
			}
		}
	}
	/*
		fmt.Println("CACHE HITS: ", globals.CacheHits)
		fmt.Println("CACHE MISSES: ", globals.CacheMisses)
	*/
}
