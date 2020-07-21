package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type store struct {
	data map[string]string

	m sync.RWMutex
}

var (
	addr = flag.String("addr", ":9000", "http service address")

	s = store{
		data: map[string]string{},
		m:    sync.RWMutex{},
	}
)

func main() {
	flag.Parse()

	r := httprouter.New()
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	err := http.ListenAndServe(*addr, r)

	if err != nil {
		log.Fatal("server failed", err)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")

	if k == "" {
		s.m.RLock()
		fmt.Fprintf(w, "read list: %v", s.data)
		s.m.RUnlock()
		return
	}
	s.m.RLock()
	fmt.Fprintf(w, "read entry: s.data[%s] = %s", k, s.data[k])
	s.m.RUnlock()
}
