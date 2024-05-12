package utils

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	suffix := WithShardingSuffix("t1", 10)
	log.Println(suffix)
}
