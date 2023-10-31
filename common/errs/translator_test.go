package errs

import "testing"

func TestInitTranslatorFromEtcd(t *testing.T) {
	InitTranslatorFromEtcd(`{"Endpoints":["192.168.2.159:2379"],"DialTimeout":5}`)
}
