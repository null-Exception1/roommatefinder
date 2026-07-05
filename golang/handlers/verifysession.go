package handlers

import (
	"database/sql"
	"fmt"
	"golang/globals"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Verify(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := req.Cookie("sess_id")

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/verify",
		"cookie":   cookie != nil,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /verify by ", req.RemoteAddr)

	fmt.Println(cookie)
	if err != nil {
		fmt.Fprint(w, "invalid")
		return
	}

	var expires time.Time
	err = globals.Globaldb.QueryRow("SELECT expires_at FROM sessions WHERE id=$1", cookie.Value).
		Scan(&expires)

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/verify",
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("checking session in db")

	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"error":    err,
			"token":    cookie,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Info("no valid session found")

		fmt.Fprint(w, "invalid")
		return
	}
	if time.Now().After(expires) {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"token":    cookie,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Info("expired token")

		fmt.Fprint(w, "expired")
	} else {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/verify",
			"token":    cookie,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Info("valid token")

		fmt.Fprint(w, "valid")

	}

}
