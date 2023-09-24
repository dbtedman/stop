package main

import (
	"net/http"
	"time"
)

type Proxy interface {
	ListenAndServe(addr string, handler http.HandlerFunc) error
}

type ServerProxy struct {
}

var _ Proxy = &ServerProxy{}

func (s *ServerProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
		Handler:      handler,
	}

	return server.ListenAndServe()
}

type TestProxy struct {
	Addr    string
	Handler http.HandlerFunc
}

var _ Proxy = &TestProxy{}

func (t *TestProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	t.Addr = addr
	t.Handler = handler
	return nil
}
