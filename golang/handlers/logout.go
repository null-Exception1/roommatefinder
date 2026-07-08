package handlers

import (
	"fmt"
	"golang/globals"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func Logout(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("SITE_URL"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Try to read the cookie
	cookie, err := req.Cookie("sess_id")

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/logout",
		"cookie":   cookie,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /logout")

	if err == nil {
		result, err := globals.Globaldb.Exec("DELETE FROM sessions WHERE id=$1", cookie.Value)

		http.SetCookie(w, &http.Cookie{
			Name:     "sess_id",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteNoneMode,
			MaxAge:   -1, // expire immediately
		})

		if err != nil {

			logrus.WithFields(logrus.Fields{
				"package":  "handlers",
				"endpoint": "/logout",
				"cookie":   cookie,
				"result":   result,
				"error":    err,
				"method":   req.Method,
				"remote":   req.RemoteAddr,
			}).Error("logout incomplete, error occurred")
		} else {
			logrus.WithFields(logrus.Fields{
				"package":  "handlers",
				"endpoint": "/logout",
				"cookie":   cookie,
				"result":   result,
				"error":    err,
				"method":   req.Method,
				"remote":   req.RemoteAddr,
			}).Info("logout complete, no errors")
		}

	} else {
		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/logout",
			"cookie":   cookie,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Error("logout incomplete, error occured", req.RemoteAddr)
	}

	fmt.Fprint(w, "logged out")
}
