package uuid

import (
	"github.com/segmentio/ksuid"
)

func NewUId() string {
	return ksuid.New().String()
}
