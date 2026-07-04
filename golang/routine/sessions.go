package routine

import (
	"database/sql"
	"fmt"
	"golang/globals"
	"time"
)

func StartSessionCleanup(db *sql.DB) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			//fmt.Println("deletion running")
			_, err := globals.Globaldb.Exec(`DELETE FROM sessions WHERE expires_at < NOW()`)
			if err != nil {
				fmt.Println("cleanup error:", err)
			}
		}
	}()
}
