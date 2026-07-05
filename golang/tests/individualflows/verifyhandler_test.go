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

func TestVerifyHandler(t *testing.T) {
	// Load env + init DB
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("failed to load .env: %v", err)
	}
	initfuncs.Database()

	// Seed DB with a fake session
	fakeToken := "a18fbd57f9bbfd0450659cb69333415f"
	_, err := globals.Globaldb.Exec(`
        INSERT INTO sessions (id, admnno, expires_at)
        VALUES ($1, $2, NOW() + interval '1 hour')
    `, fakeToken, "69")
	if err != nil {
		t.Fatalf("failed to insert session: %v", err)
	}

	// Make request with cookie
	req := httptest.NewRequest("GET", "/verify", nil)
	req.AddCookie(&http.Cookie{Name: "sess_id", Value: fakeToken})
	w := httptest.NewRecorder()

	handlers.Verify(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if w.Body.String() != "valid" {
		t.Errorf("expected 'valid', got %q", w.Body.String())
	}

	// Cleanup
	if _, err := globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno=$1", "69"); err != nil {
		t.Logf("cleanup failed: %v", err)
	}
}
