package db

import (
	"database/sql"
	"golang/structs"

	"github.com/sirupsen/logrus"
)

func Insert(person structs.Person, db *sql.DB) bool {

	logrus.WithFields(logrus.Fields{
		"package":  "db",
		"function": "insert",
	}).Debug("insert request to db initiated")

	_, err := db.Exec(`
    INSERT INTO people (admn_hash, name, social, socialtype, roomno, blockno) 
    VALUES ($1 , $2, $3, $4, $5, $6)
    ON CONFLICT (admn_hash) DO NOTHING`, person.Admnno, person.Name, person.Social, person.Socialtype, person.Roomno, person.Blockno)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"package":  "db",
			"function": "insert",
		}).Debug("insert successful")
		return false
	} else {
		logrus.WithFields(logrus.Fields{
			"package":  "db",
			"function": "insert",
			"error":    err,
		}).Error("insert unsuccessful, error")
		return true
	}

}
