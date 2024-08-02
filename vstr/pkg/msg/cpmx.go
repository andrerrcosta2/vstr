// github.com/andrerrcosta2 *** andrerrcosta@gmail.com

package msg

import (
	"google.golang.org/protobuf/proto"
)

type Cpmx[T proto.Message, K proto.Message] interface {
	Spm(msg T, addr string) (K, error)
	//Close()
}
