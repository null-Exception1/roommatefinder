package init

import (
	"database/sql"
	"golang/globals"
	"os"

	"github.com/sirupsen/logrus"
)

func Database() {
	connStr := os.Getenv("DATABASE_URL")
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
