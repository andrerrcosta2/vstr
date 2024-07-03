package cmx

import "google.golang.org/protobuf/proto"

type Cpmx interface {
	Spm(msg proto.Message, addr string) (proto.Message, error)
}
