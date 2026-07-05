package handlers

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Rooms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	q := req.URL.Query()

	blockno := q.Get("block")

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"block":    blockno,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /rooms")

	sqlquery := fmt.Sprintf("SELECT * FROM people WHERE blockno='%s';", blockno)

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"query":    sqlquery,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("sql query created")

	rows := db.Query(string(sqlquery), globals.Globaldb)
	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"rows":     len(rows),
		"method":   req.Method,
		"remote":   req.RemoteAddr,
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
		"package":  "handlers",
		"endpoint": "/rooms",
		"rooms":    len(rooms),
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("rooms is mapped")

	str, _ := json.Marshal(rooms)

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/rooms",
		"status":   http.StatusOK,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("response sent")

	fmt.Fprintf(w, "%s", string(str))

}
