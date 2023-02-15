package snowflake

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func InitSnowflake() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func NextID() (uint64, error) {
	id, err := sf.NextID()
	return id, err
}
