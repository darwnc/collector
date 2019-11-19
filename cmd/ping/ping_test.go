package ping

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	entity := pingEntity{time.Duration(60 * 1000 * 1000), 100, "www.baidu.com", 32}
	entity.ping()
}
