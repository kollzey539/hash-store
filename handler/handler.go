package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kollzey539/hash-store/storage"
	"github.com/kollzey539/hash-store/util"

	"github.com/gorilla/mux"
)

type Request struct {
	Data string `json:"data"`
}

type Response struct {
	Hash string `json:"hash,omitempty"`
	Data string `json:"data,omitempty"`
}

func CreateHashHandler(s *storage.S3Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// reuse StoreHandler logic, injecting storage
		body, err := io.ReadAll(r.Body)
		if err != nil || strings.TrimSpace(string(body)) == "" {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var req Request
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		hash := util.GenerateSHA256(req.Data)
		err = s.PutItem(hash, req.Data)
		if err != nil {
			fmt.Println("PutItem error:", err) // <-- add this line
			http.Error(w, "Failed to store: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(Response{Hash: hash})
	}
}

func GetStringHandler(s *storage.S3Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := mux.Vars(r)["hash"]
		data, err := s.GetItem(hash)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(Response{Data: data})
	}
}
