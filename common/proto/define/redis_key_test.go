package define

import (
	"log"
	"testing"
)

func TestWithParam(t *testing.T) {
	result := AccountToken.WithParams("1111", "222")
	log.Println(result)
}
