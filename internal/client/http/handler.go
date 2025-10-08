package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	gw "test_task/internal/client/gateway"
)

type JSONResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}

func NewHTTPHandler(gw *gw.Gateway) http.Handler {
	mux := http.NewServeMux()

	// POST /data → Set
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleSet(w, r, gw)
		case http.MethodGet:
			handleGet(w, r, gw)
		default:
			http.Error(w, "only GET and POST are supported", http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func handleSet(w http.ResponseWriter, r *http.Request, gw *gw.Gateway) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		http.Error(w, fmt.Sprintf("invalid JSON: %v", decodeErr), http.StatusBadRequest)
		return
	}
	if req.Key == "" {
		http.Error(w, "missing 'key'", http.StatusBadRequest)
		return
	}

	msg, err := gw.Set(req.Key, req.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(JSONResponse{Success: true, Message: msg})
	if encodeErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, gw *gw.Gateway) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "missing query param 'key'", http.StatusBadRequest)
		return
	}

	val, found, err := gw.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !found {
		encodeErr := json.NewEncoder(w).Encode(JSONResponse{Success: false, Message: "not found"})
		if encodeErr != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	encodeErr := json.NewEncoder(w).Encode(JSONResponse{Success: true, Value: val})
	if encodeErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
