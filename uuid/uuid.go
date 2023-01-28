package uuid

import (
	"time"
)

type UUID [16]byte

type UUIDBuilder struct {
	Timestamp *int64
}

func New() UUIDBuilder {
	now := time.Now().Unix()
	return UUIDBuilder{Timestamp: &now}
}
