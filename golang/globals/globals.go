package globals

import (
	"database/sql"
	"golang/structs"
	"sync"
	"time"
)

var Globaldb *sql.DB
var Ticker *time.Ticker
var RatelimitChannel chan time.Time

// For caching blocks
var CacheBlocks map[string]*structs.Block = make(map[string]*structs.Block, 0)
var CachedBlocksJSON string
var CacheExpiry time.Time
var CacheMutex sync.RWMutex

// For caching rooms : block -> room number -> structs.Room
var CacheRooms map[string]map[string]*structs.Room = make(map[string]map[string]*structs.Room, 0)
var CachedRoomsJSON map[string]string = make(map[string]string, 0)

// Per block expiry time
var CacheBlocksExpiry map[string]time.Time = make(map[string]time.Time, 0)
var CacheRoomsMutex sync.RWMutex

// Analytics
var CacheHits int = 0
var CacheMisses int = 0
