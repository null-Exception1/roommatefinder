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
var CacheBlocks map[string]*structs.Block = make(map[string]*structs.Block, 0)
var CacheExpiry time.Time
var CacheMutex sync.RWMutex
