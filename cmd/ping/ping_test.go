package ping

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	// 172.16.0.31
	entity := pingEntity{time.Duration(60 * 1000 * 1000), 2, []string{"www.baidu.com", "www.google.com"}, 64}
	entity.ping()
}
