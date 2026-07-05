package individualflows

import (
	"golang/globals"
	"golang/handlers"
	initfuncs "golang/init"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoginHandler(t *testing.T) {
	// Seed DB with a fake user first
	godotenv.Load("../../.env")

	initfuncs.Database()

	_, err := globals.Globaldb.Exec(`
    INSERT INTO people (admn_hash, name, social, socialtype, roomno, blockno)
    VALUES ($1, $2, $3, $4, $5, $6)
`, "69", "Shaurya", "discordusername", "Discord", 123, 16)

	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}

	req := httptest.NewRequest("GET", "/login?admn_hash=69&name=Shaurya", nil)
	w := httptest.NewRecorder()

	handlers.Login(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if w.Body.String() == "not found" {
		t.Errorf("expected token, got %q", w.Body.String())
	}

	// cleanup
	globals.Globaldb.Exec("DELETE FROM people WHERE admn_hash='69'")
	globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno='69'")
}
