package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/This-Is-Prince/loadbalancer/utils"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive() bool {
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}

func NewSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	utils.HandleErr(err)
	return &SimpleServer{Addr: addr, Proxy: httputil.NewSingleHostReverseProxy(serverUrl)}
}

func NewSimpleServers(addrs []string) []Server {
	servers := make([]Server, 0)
	for _, addr := range addrs {
		servers = append(servers, NewSimpleServer(addr))
	}
	return servers
}
