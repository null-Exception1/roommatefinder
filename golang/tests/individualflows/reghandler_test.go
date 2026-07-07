package individualflows

import (
	"golang/globals"
	"golang/handlers"
	initfuncs "golang/init"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestRegistrationHandler(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("failed to load .env: %v", err)
	}
	initfuncs.Database()

	req := httptest.NewRequest("GET",
		"/registration?admn_hash=69&name=Shaurya&social=discordusername&socialtype=Discord&blockno=16&roomno=123&created_at=now",
		nil)
	w := httptest.NewRecorder()

	handlers.RegistrationHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if w.Body.String() != "done" {
		t.Errorf("expected 'done', got %q", w.Body.String())
	}

	cookies := resp.Cookies()
	if len(cookies) == 0 {
		t.Fatal("expected a cookie, got none")
	}

	peppered := globals.SecureHash("69", os.Getenv("PEPPER"))

	if _, err := globals.Globaldb.Exec("DELETE FROM people WHERE admn_hash=$1", peppered); err != nil {
		t.Logf("cleanup people failed: %v", err)
	}
	if _, err := globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno=$1", peppered); err != nil {
		t.Logf("cleanup sessions failed: %v", err)
	}
}
