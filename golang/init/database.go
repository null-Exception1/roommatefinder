package init

import (
	"database/sql"
	"golang/globals"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Database() {
	connStr := os.Getenv("DATABASE_URL")
	logrus.Debug(connStr)
	//connStr = "postgres://devuser:devpass@localhost:5432/roommatefinder?sslmode=disable"
	var err any
	globals.Globaldb, err = sql.Open("postgres", connStr)

	if err != nil {

		logrus.WithFields(logrus.Fields{
			"package": "init",
			"error":   err,
		}).Error("error in initiating db conn")

		panic(err)
	} else {
		logrus.WithFields(logrus.Fields{
			"package": "init",
		}).Info("initated database connection")
	}

}
