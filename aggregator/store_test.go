package main

import (
	"testing"
)

func TestAdminGetBooking(t *testing.T) {
	m := NewMemoryStore()
	dist, ok := m.data[805283365686263469]
	print("dist", dist, ok)
}
