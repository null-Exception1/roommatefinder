package caching

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"time"

	"github.com/sirupsen/logrus"
)

func CacheRoomsUpdate(blockno string) {

	sqlquery := fmt.Sprintf("SELECT * FROM people WHERE blockno='%s';", blockno)

	logrus.WithFields(logrus.Fields{
		"package": "caching",
		"query":   sqlquery,
	}).Debug("sql query created")

	rows := db.Query(string(sqlquery), globals.Globaldb)

	logrus.WithFields(logrus.Fields{
		"package": "caching",
		"rows":    len(rows),
	}).Debug("query fetch happened")

	rooms := make(map[string]*structs.Room, 0)

	for _, row := range rows {
		if _, ok := rooms[row.Roomno]; ok {
			rooms[row.Roomno].People = append(rooms[row.Roomno].People, &structs.Person{Admnno: row.Admnno, Name: row.Name, Social: row.Social, Socialtype: row.Socialtype, Roomno: row.Roomno, Blockno: row.Blockno, Created_at: row.Created_at})
		} else {
			rooms[row.Roomno] = &structs.Room{People: make([]*structs.Person, 0)}
			rooms[row.Roomno].People = append(rooms[row.Roomno].People, &structs.Person{Admnno: row.Admnno, Name: row.Name, Social: row.Social, Socialtype: row.Socialtype, Roomno: row.Roomno, Blockno: row.Blockno, Created_at: row.Created_at})
		}
	}

	logrus.WithFields(logrus.Fields{
		"package":  "caching",
		"endpoint": "/rooms",
		"rooms":    len(rooms),
	}).Debug("rooms is mapped")

	globals.CacheRoomsMutex.Lock()
	bytes, _ := json.Marshal(rooms)
	globals.CachedRoomsJSON[blockno] = string(bytes)
	globals.CacheBlocksExpiry[blockno] = time.Now().Add(30 * time.Second)
	globals.CacheRoomsMutex.Unlock()

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"block":    blockno,
	}).Debug("updated block cache")

	//fmt.Println(globals.CacheRooms[blockno])
}
