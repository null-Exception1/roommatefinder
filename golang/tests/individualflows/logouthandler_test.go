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

func TestLogoutHandler(t *testing.T) {
	godotenv.Load("../../.testenv")

	initfuncs.Database()
	_, err := globals.Globaldb.Exec(`
        INSERT INTO sessions (id, admnno) VALUES ($1, $2)
    `, "a18fbd57f9bbfd0450659cb69333415f", "69")
	if err != nil {
		t.Fatalf("failed to insert session: %v", err)
	}

	req := httptest.NewRequest("GET", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "sess_id", Value: "a18fbd57f9bbfd0450659cb69333415f"})
	w := httptest.NewRecorder()

	handlers.Logout(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if w.Body.String() != "logged out" {
		t.Errorf("expected 'logged out', got %q", w.Body.String())
	}

	// cleanup
	globals.Globaldb.Exec("DELETE FROM sessions WHERE admnno=$1", "69")
}
