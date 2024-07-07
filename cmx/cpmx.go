// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package cmx

import "google.golang.org/protobuf/proto"

type Cpmx interface {
	Spm(msg proto.Message, addr string) (proto.Message, error)
}
