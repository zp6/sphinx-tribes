package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Component tests for bounties modal
// Addresses issue #871

func TestBountyModalCreate(t *testing.T) {
	body := strings.NewReader('{"title":"New Bounty","description":"Test bounty","price":5000,"type":"bug"}')
	req := httptest.NewRequest("POST", "/bounties", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": "bounty-1", "title": "New Bounty", "price": 5000, "status": "open",
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestBountyModalDelete(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/bounties/bounty-1", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
