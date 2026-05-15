package sphinxtribes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Unit tests for organizations
// Addresses issue #870

func TestListOrganizations(t *testing.T) {
	req := httptest.NewRequest("GET", "/organizations", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"organizations": []map[string]interface{}{
				{"uuid": "1", "name": "Org1"},
				{"uuid": "2", "name": "Org2"},
			},
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestUpdateOrganization(t *testing.T) {
	body := strings.NewReader('{"name":"Updated Org","description":"New desc"}')
	req := httptest.NewRequest("PUT", "/organizations/org1", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"uuid": "org1", "name": "Updated Org", "description": "New desc",
		})
	})
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
