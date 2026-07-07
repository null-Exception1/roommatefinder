package globals

import (
	"database/sql"
	"golang/structs"
	"sync"
	"sync/atomic"
	"time"
)

// Editables
const CacheBlocksSeconds = 60
const CacheRoomsSeconds = 60
const CacheBlocksRoutine = 30
const CacheRoomsRoutine = 30
const NumWorkers = 50
const NumCacheRoomsJobs = 1000
const NumCacheRoomsJobsResults = 1000

var Globaldb *sql.DB
var Ticker *time.Ticker = time.NewTicker(200 * time.Millisecond)
var RatelimitChannel chan time.Time

// For caching blocks
var CacheBlocks map[string]*structs.Block = make(map[string]*structs.Block, 0)
var CachedBlocksJSON string
var CacheExpiry time.Time
var CacheMutex sync.RWMutex

// For caching rooms : block -> room number -> structs.Room
var CacheRooms map[string]map[string]*structs.Room = make(map[string]map[string]*structs.Room, 0) // deprecated
var CachedRoomsJSON sync.Map

// Work group for caching rooms
var CacheRoomsJobsWaitGroup sync.WaitGroup
var CacheRoomsJobs chan structs.RoomsJob = make(chan structs.RoomsJob, NumCacheRoomsJobs)
var CacheRoomsJobsResults chan structs.RoomsJobResult = make(chan structs.RoomsJobResult, NumCacheRoomsJobsResults)

// Per block expiry time
var CacheBlocksExpiry sync.Map
var CacheRoomsMutex sync.RWMutex

// Analytics
var CacheHits atomic.Int32
var CacheMisses atomic.Int32
