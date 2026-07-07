package simplefetch

import (
	"golang/caching"
	"golang/globals"
	"golang/handlers"
	initfuncs "golang/init"
	"golang/routine"
	"golang/structs"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkRoomsBlocksHandler(t *testing.B) {
	godotenv.Load("../../.testenv")
	initfuncs.Logging()
	initfuncs.Database()

	// one-time setup
	globals.CacheRoomsJobs = make(chan structs.RoomsJob, globals.NumCacheRoomsJobs)
	for w := 0; w < globals.NumWorkers; w++ {
		globals.CacheRoomsJobsWaitGroup.Go(func() {
			for job := range globals.CacheRoomsJobs {
				caching.CacheRoomsUpdate(job.Blockno)
			}
		})
	}

	var total int64
	go routine.WorkersResults()

	t.ResetTimer() // ignore setup cost

	t.RunParallel(func(pb *testing.PB) {
		localIteration := 0
		for pb.Next() {
			req := httptest.NewRequest("GET", "/rooms?block="+strconv.Itoa(localIteration), nil)
			w := httptest.NewRecorder()
			handlers.Rooms(w, req)
			atomic.AddInt64(&total, 1)
			localIteration++
		}
	})

	t.StopTimer()

	close(globals.CacheRoomsJobs)
	globals.CacheRoomsJobsWaitGroup.Wait()

	t.ReportMetric(float64(total)/t.Elapsed().Seconds(), "req/s")
	// optional: cleanup after all iterations
}
