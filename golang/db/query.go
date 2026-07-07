package db

import (
	"database/sql"
	"golang/structs"

	"github.com/sirupsen/logrus"
)

func Query(query string, db *sql.DB) []structs.Person {

	rows, _ := db.Query(query)

	if rows != nil {
		defer rows.Close()
	}

	logrus.WithFields(logrus.Fields{
		"package":  "db",
		"function": "query",
	}).Debug("fetch query recieved, making results..")

	results := make([]structs.Person, 0)

	for rows.Next() {
		var admnno string
		var name string
		var social string
		var socialtype string
		var roomno string
		var blockno string
		var created_at string

		err := rows.Scan(&admnno, &name, &social, &socialtype, &roomno, &blockno, &created_at)

		p := structs.Person{Admnno: admnno, Name: name, Social: social, Socialtype: socialtype, Roomno: roomno, Blockno: blockno, Created_at: created_at}

		results = append(results, p)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"package":  "db",
				"function": "query",
				"error":    err,
			}).Error("Error occured in query")

		}

	}
	defer rows.Close()
	logrus.WithFields(logrus.Fields{
		"package":  "db",
		"function": "query",
		"results":  len(results),
	}).Debug("returning results from query...")

	return results
}
