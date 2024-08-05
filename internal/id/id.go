package id

import (
	"encoding/json"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var err error
	sf, err = sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}
}

type ID int64

func New() ID {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return ID(id)
}

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id ID) IsZero() bool { return int64(id) == 0 }

func (id ID) MarshalJSON() ([]byte, error) {
	s := id.String()

	b := make([]byte, 0, len(s)+2)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}

	n, err := number.Int64()
	if err != nil {
		return err
	}

	*id = ID(n)
	return nil
}

type NullID struct {
	pgtype.Int8
}

func (id NullID) String() string {
	if !id.Valid {
		return "null"
	}
	return ID(id.Int64).String()
}
