package dds

import "github.com/andrerrcosta2/vstr/dds/ddb"

type Dsn interface {
	Str(block ddb.Blk) error
	Rtv(id string) (ddb.Blk, error)
	Dlt(id string) error
}
