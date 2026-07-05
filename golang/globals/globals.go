package globals

import (
	"database/sql"
	"time"
)

var Globaldb *sql.DB
var Ticker *time.Ticker
var RatelimitChannel chan time.Time
