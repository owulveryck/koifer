package koifer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"unicode/utf8"

	"github.com/owulveryck/koifer"

	"github.com/owulveryck/koifer/db/memory"
)

func TestMemoryDB(t *testing.T) {
	db := memory.NewDB()
	db.UpsertUser("dexter", "killer")
	testDB(db, t)
}

func testDB(db koifer.UserRepository, t *testing.T) {
	authService := koifer.NewAuthService(db)
	ts := httptest.NewServer(authService)
	t.Run("test ok", func(t *testing.T) {
		user := &koifer.User{"dexter", "killer"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		var token struct {
			Token string
		}
		err = dec.Decode(&token)
		if err != nil {
			t.Fatal(err)
		}
		if utf8.RuneCountInString(token.Token) != 32 {
			t.Errorf("expected 32 chars token, got %v", utf8.RuneCountInString(token.Token))
		}
	})
	t.Run("test bad password", func(t *testing.T) {
		user := &koifer.User{"dexter", "friend"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusUnauthorized, resp.StatusCode)
		}
	})
	t.Run("bad content", func(t *testing.T) {
		user := "bad content"
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusBadRequest, resp.StatusCode)
		}
	})
	t.Run("no name", func(t *testing.T) {
		user := struct {
			Password string
		}{
			Password: "killer",
		}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusUnauthorized, resp.StatusCode)
		}
	})
	t.Run("no password", func(t *testing.T) {
		user := struct {
			Name string
		}{
			Name: "dexter",
		}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusUnauthorized, resp.StatusCode)
		}
	})
	t.Run("test user not found", func(t *testing.T) {
		user := &koifer.User{"olivier", "killer"}
		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.Post(ts.URL, "application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("bad return code, expeted %v, got %v", http.StatusUnauthorized, resp.StatusCode)
		}
	})

}
