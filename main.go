package main

import (
	"flag"
	"sync"
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
