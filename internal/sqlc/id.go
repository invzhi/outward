package sqlc

import "github.com/sony/sonyflake"

var sf *sonyflake.Sonyflake

func init() {
	var err error
	sf, err = sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}
}

type ID int64

func NewID() int64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return int64(id)
}
