package http

import (
	"net/http"

	"github.com/christianvozar/hex-example/internal/domain/watcher"
)

type server struct {
	watcher watcher.Watcher
}

func NewServer(watcher watcher.Watcher) *server {
	return &server{watcher: watcher}
}

func (s *server) TriggerUpdate(w http.ResponseWriter, r *http.Request) {
	if err := s.watcher.ManualUpdate(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func StartHTTPServer(watcher watcher.Watcher, address string) error {
	s := NewServer(watcher)

	http.HandleFunc("/update", s.TriggerUpdate)

	return http.ListenAndServe(address, nil)
}
