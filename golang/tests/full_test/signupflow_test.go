package full_test

import (
	"golang/globals"
	handlers "golang/handlers"
	initiation "golang/init"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestSignupflow(t *testing.T) {
	_ = godotenv.Load("../../.env")
	initiation.Database()

	frontendHash := "69" // what the client sends
	peppered := globals.SecureHash(frontendHash, os.Getenv("PEPPER"))

	// registration flow
	logrus.Info("/registration testing in progress..")

	req := httptest.NewRequest("GET",
		"/registration?admn_hash="+frontendHash+"&name=Shaurya&social=discordusername&socialtype=Discord&blockno=16&roomno=123&created_at=now",
		nil)
	w := httptest.NewRecorder()

	handlers.RegistrationHandler(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if body != "done" {
		t.Errorf("expected 'done', got %q", body)
	}
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		t.Fatal("expected a cookie, got none")
	}
	token := cookies[0].Value

	logrus.Info("/registration testing complete..")

	defer func() {
		globals.Globaldb.Exec("DELETE FROM people WHERE admn_hash=$1", peppered)
		globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno=$1", peppered)
	}()

	logrus.Info("got token as (assume its not malformed) ", token)

	// verify flow
	logrus.Info("/verify testing in progress..")

	req = httptest.NewRequest("GET", "/verify", nil)
	req.AddCookie(&http.Cookie{Name: "sess_id", Value: token})
	w = httptest.NewRecorder()

	handlers.Verify(w, req)

	resp = w.Result()
	body = w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if body != "valid" {
		t.Errorf("expected 'valid', got %q", body)
	}

	logrus.Info("/verify ended..")

	// logout flow
	logrus.Info("/logout testing in progress..")

	req = httptest.NewRequest("GET", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "sess_id", Value: token})
	w = httptest.NewRecorder()

	handlers.Logout(w, req)

	resp = w.Result()
	body = w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if body != "logged out" {
		t.Errorf("expected 'logged out', got %q", body)
	}

	logrus.Info("/logout testing complete..")

	// login flow
	logrus.Info("/login testing in progress..")

	req = httptest.NewRequest("GET", "/login?admn_hash="+frontendHash+"&name=Shaurya", nil)
	w = httptest.NewRecorder()

	handlers.Login(w, req)

	resp = w.Result()
	body = w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if body == "not found" {
		t.Errorf("expected token, got %q", body)
	}

	logrus.Info("/login testing complete..")
}
