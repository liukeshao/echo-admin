package uid

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenUint64() (uint64, error) {
	return sf.NextID()
}
