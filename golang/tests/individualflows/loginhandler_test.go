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

func TestLoginHandler(t *testing.T) {
	_ = godotenv.Load("../../.env")
	initfuncs.Database()

	frontendHash := "69" // what the client would send
	peppered := globals.SecureHash(frontendHash, os.Getenv("PEPPER"))

	_, err := globals.Globaldb.Exec(`
        INSERT INTO people (admn_hash, name, social, socialtype, roomno, blockno)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, peppered, "Shaurya", "discordusername", "Discord", 123, 16)

	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}

	req := httptest.NewRequest("GET", "/login?admn_hash="+frontendHash+"&name=Shaurya", nil)
	w := httptest.NewRecorder()

	handlers.Login(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if w.Body.String() == "not found" {
		t.Errorf("expected token, got %q", w.Body.String())
	}

	_, _ = globals.Globaldb.Exec("DELETE FROM people WHERE admn_hash=$1", peppered)
	_, _ = globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno=$1", peppered)
}
