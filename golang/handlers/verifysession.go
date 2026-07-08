package handlers

import (
	"database/sql"
	"fmt"
	"golang/globals"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Verify(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("SITE_URL"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := req.Cookie("sess_id")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"reason":   "cookie_missing",
		}).Info("missing session cookie")

		fmt.Fprint(w, "invalid")
		return
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/verify",
		"method":   req.Method,
	}).Debug("checking session in db")

	var expires time.Time
	err = globals.Globaldb.QueryRow("SELECT expires_at FROM sessions WHERE id=$1", cookie.Value).Scan(&expires)

	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"token":    cookie.Value,
		}).Info("no matching session found in db")

		fmt.Fprint(w, "invalid")
		return
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"error":    err,
		}).Error("database query execution failed")

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if time.Now().After(expires) {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"token":    cookie.Value,
		}).Info("expired token presented")

		fmt.Fprint(w, "expired")
		return
	}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/verify",
		"token":    cookie.Value,
	}).Info("valid token authorized")

	fmt.Fprint(w, "valid")
}
