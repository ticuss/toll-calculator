package main

import (
	"fmt"
	"os"

	"github.com/tolling/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(d types.Distance) error {
	fmt.Fprintf(os.Stdout, `storing Data %s`, []any{d}...)
	m.data[d.OBUID] += d.Value
	return nil
}

func (m *MemoryStore) Get(id int) (float64, error) {
	dist, ok := m.data[id]
	for i, v := range m.data {
		fmt.Print(i, v)
	}
	if !ok {
		return 0, fmt.Errorf("distance not found for obu %d", id)
	}
	return dist, nil
}
